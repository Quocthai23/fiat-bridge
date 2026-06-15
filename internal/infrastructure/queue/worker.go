package queue

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"time"

	"github.com/Quocthai23/fiat-bridge/internal/domain"
	"github.com/Quocthai23/fiat-bridge/internal/infrastructure/blockchain"
	"github.com/Quocthai23/fiat-bridge/internal/infrastructure/db"

	"github.com/ethereum/go-ethereum/common"
	amqp "github.com/rabbitmq/amqp091-go"
)

// StartConsumerWorker starts listening to the RabbitMQ queue for mint transactions
func StartConsumerWorker(ctx context.Context, emitter *blockchain.Emitter) {
	msgs, err := RabbitChannel.Consume(
		MintQueueName, // queue
		"",            // consumer
		false,         // auto-ack (changed to false for reliability)
		false,         // exclusive
		false,         // no-local
		false,         // no-wait
		nil,           // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case d, ok := <-msgs:
				if !ok {
					return
				}

				var tx domain.Transaction
			if err := json.Unmarshal(d.Body, &tx); err != nil {
				log.Printf("Error decoding message: %v\n", err)
				d.Reject(false) // Reject corrupt message (do not requeue)
				continue
			}

			// Check retry count
			retryCount := 0
			if d.Headers != nil {
				if death, ok := d.Headers["x-death"].([]interface{}); ok && len(death) > 0 {
					if deathInfo, ok := death[0].(amqp.Table); ok {
						if count, ok := deathInfo["count"].(int64); ok {
							retryCount = int(count)
						}
					}
				}
			}

			if retryCount >= 3 {
				log.Printf("Tx %s exceeded retry limit. Sending to DLQ.\n", tx.CoreTxId)
				// Reject without requeue -> Goes to DLX -> DLQ
				d.Reject(false)

				// Update DB to Failed status
				var dbTx domain.Transaction
				if err := db.DB.Where("core_tx_id = ?", tx.CoreTxId).First(&dbTx).Error; err == nil {
					dbTx.Status = domain.StatusFailed
					db.UpdateTransaction(&dbTx)
				}
				continue
			}

			// Check Idempotency: Do not process if already completed or pending on chain
			var dbTx domain.Transaction
			if err := db.DB.Where("core_tx_id = ?", tx.CoreTxId).First(&dbTx).Error; err == nil {
				if dbTx.Status == domain.StatusCompleted || dbTx.Status == domain.StatusPendingOnChain {
					log.Printf("Tx %s already in status %s. Skipping RPC to save gas.\n", tx.CoreTxId, dbTx.Status)
					d.Ack(false)
					continue
				}
			}

			// Call Smart Contract (only signs)
			ctx := context.Background()
			userAddr := common.HexToAddress(tx.UserAddress)
			amount, ok := new(big.Int).SetString(tx.Amount, 10)
			if !ok {
				log.Printf("Invalid amount string for TxID %s: %v. Sending to DLQ...\n", tx.CoreTxId, tx.Amount)
				d.Reject(false)
				continue
			}
			signedTx, nonce, err := emitter.MintTokens(ctx, tx.CoreTxId, userAddr, amount)
			if err != nil {
				log.Printf("Failed to sign tokens for TxID %s: %v. Sending to DLQ...\n", tx.CoreTxId, err)
				// Reject without requeue, letting it go to DLX -> DLQ
				d.Reject(false)
				continue
			}

			txHash := signedTx.Hash().Hex()
			log.Printf("Successfully signed tokens for TxID %s. Hash: %s, Nonce: %d\n", tx.CoreTxId, txHash, nonce)

			// Fetch tx from DB to update BEFORE broadcasting
			var updateTx domain.Transaction
			if err := db.DB.Where("core_tx_id = ?", tx.CoreTxId).First(&updateTx).Error; err == nil {
				updateTx.TxHash = txHash
				updateTx.Nonce = nonce
				updateTx.Status = domain.StatusPendingOnChain
				if updateErr := db.UpdateTransaction(&updateTx); updateErr != nil {
					log.Printf("Failed to update DB for TxID %s: %v. Rejecting message to retry.\n", tx.CoreTxId, updateErr)
					emitter.ReleaseNonce(nonce)
					d.Reject(false)
					continue
				}
			} else {
				log.Printf("Failed to find tx %s in DB: %v. Rejecting message.\n", tx.CoreTxId, err)
				d.Reject(false)
				continue
			}

			// 3. Broadcast the transaction
			if err := emitter.BroadcastTransaction(ctx, signedTx); err != nil {
				log.Printf("Failed to broadcast tx %s: %v. GasBumper will handle retries.\n", tx.CoreTxId, err)
			} else {
				log.Printf("Successfully broadcast tx %s\n", tx.CoreTxId)
			}

			// We Ack the message because it's safely in our DB now.
			d.Ack(false)
			}
		}
	}()

	// Start DLQ Worker
	go StartDLQWorker(ctx)

	fmt.Println("Consumer Worker started listening for messages.")
}

// StartBurnConsumerWorker starts listening to the RabbitMQ queue for burn transactions
func StartBurnConsumerWorker(ctx context.Context, emitter *blockchain.Emitter) {
	msgs, err := RabbitChannel.Consume(
		BurnQueueName, // queue
		"",            // consumer
		false,         // auto-ack (changed to false for reliability)
		false,         // exclusive
		false,         // no-local
		false,         // no-wait
		nil,           // args
	)
	if err != nil {
		log.Fatalf("Failed to register a burn consumer: %v", err)
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case d, ok := <-msgs:
				if !ok {
					return
				}

				var tx domain.Transaction
			if err := json.Unmarshal(d.Body, &tx); err != nil {
				log.Printf("Error decoding message: %v\n", err)
				d.Reject(false) // Reject corrupt message (do not requeue)
				continue
			}

			// Check retry count
			retryCount := 0
			if d.Headers != nil {
				if death, ok := d.Headers["x-death"].([]interface{}); ok && len(death) > 0 {
					if deathInfo, ok := death[0].(amqp.Table); ok {
						if count, ok := deathInfo["count"].(int64); ok {
							retryCount = int(count)
						}
					}
				}
			}

			if retryCount >= 3 {
				log.Printf("Tx %s exceeded retry limit. Sending to DLQ.\n", tx.CoreTxId)
				// Reject without requeue -> Goes to DLX -> DLQ
				d.Reject(false)

				// Update DB to Failed status
				var dbTx domain.Transaction
				if err := db.DB.Where("core_tx_id = ?", tx.CoreTxId).First(&dbTx).Error; err == nil {
					dbTx.Status = domain.StatusFailed
					db.UpdateTransaction(&dbTx)
				}
				continue
			}

			// Check Idempotency: Do not process if already completed or pending on chain
			var dbTx domain.Transaction
			if err := db.DB.Where("core_tx_id = ?", tx.CoreTxId).First(&dbTx).Error; err == nil {
				if dbTx.Status == domain.StatusCompleted || dbTx.Status == domain.StatusPendingOnChain {
					log.Printf("Tx %s already in status %s. Skipping RPC to save gas.\n", tx.CoreTxId, dbTx.Status)
					d.Ack(false)
					continue
				}
			}

			// Call Smart Contract (only signs)
			ctx := context.Background()
			userAddr := common.HexToAddress(tx.UserAddress)
			amount, ok := new(big.Int).SetString(tx.Amount, 10)
			if !ok {
				log.Printf("Invalid amount string for TxID %s: %v. Sending to DLQ...\n", tx.CoreTxId, tx.Amount)
				d.Reject(false)
				continue
			}
			signedTx, nonce, err := emitter.BurnTokens(ctx, tx.CoreTxId, userAddr, amount)
			if err != nil {
				log.Printf("Failed to sign tokens for TxID %s: %v. Sending to DLQ...\n", tx.CoreTxId, err)
				// Reject without requeue, letting it go to DLX -> DLQ
				d.Reject(false)
				continue
			}

			txHash := signedTx.Hash().Hex()
			log.Printf("Successfully signed tokens for TxID %s. Hash: %s, Nonce: %d\n", tx.CoreTxId, txHash, nonce)

			// Fetch tx from DB to update BEFORE broadcasting
			var updateTx domain.Transaction
			if err := db.DB.Where("core_tx_id = ?", tx.CoreTxId).First(&updateTx).Error; err == nil {
				updateTx.TxHash = txHash
				updateTx.Nonce = nonce
				updateTx.Status = domain.StatusPendingOnChain
				if updateErr := db.UpdateTransaction(&updateTx); updateErr != nil {
					log.Printf("Failed to update DB for TxID %s: %v. Rejecting message to retry.\n", tx.CoreTxId, updateErr)
					emitter.ReleaseNonce(nonce)
					d.Reject(false)
					continue
				}
			} else {
				log.Printf("Failed to find tx %s in DB: %v. Rejecting message.\n", tx.CoreTxId, err)
				d.Reject(false)
				continue
			}

			// 3. Broadcast the transaction
			if err := emitter.BroadcastTransaction(ctx, signedTx); err != nil {
				log.Printf("Failed to broadcast tx %s: %v. GasBumper will handle retries.\n", tx.CoreTxId, err)
			} else {
				log.Printf("Successfully broadcast tx %s\n", tx.CoreTxId)
			}

			// We Ack the message because it's safely in our DB now.
			d.Ack(false)
			}
		}
	}()

	fmt.Println("Burn Consumer Worker started listening for messages.")
}

func StartDLQWorker(ctx context.Context) {
	msgs, err := RabbitChannel.Consume(
		DLQName,
		"",
		false, // auto-ack changed to false for reliability
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register DLQ consumer: %v", err)
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case d, ok := <-msgs:
				if !ok {
					return
				}

				var tx domain.Transaction
		if err := json.Unmarshal(d.Body, &tx); err != nil {
			d.Reject(false)
			continue
		}

		log.Printf("[DLQ Worker] Processing failed TxID: %s. Initiating Rollback...\n", tx.CoreTxId)
		
		payload, _ := json.Marshal(tx)
		resp, err := http.Post("https://core-banking-internal/api/webhook/rollback", "application/json", bytes.NewBuffer(payload))
		if err == nil && resp.StatusCode == 200 {
			log.Printf("[DLQ Worker] Rollback SUCCESS for TxID: %s\n", tx.CoreTxId)
			d.Ack(false)
		} else {
			log.Printf("[DLQ Worker] Rollback FAIL. Core Banking unreachabe. Requeueing...")
			time.Sleep(10 * time.Second)
			d.Reject(true)
		}
			}
		}
	}()
}
