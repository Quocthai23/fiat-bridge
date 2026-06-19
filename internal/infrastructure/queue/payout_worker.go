package queue

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Quocthai23/fiat-bridge/internal/domain"
	"github.com/Quocthai23/fiat-bridge/internal/infrastructure/db"
)

// StartPayoutWorker starts listening to the RabbitMQ payout queue for off-ramp requests
func StartPayoutWorker(ctx context.Context) {
	msgs, err := RabbitChannel.Consume(
		PayoutQueueName, // queue
		"",              // consumer
		false,           // auto-ack
		false,           // exclusive
		false,           // no-local
		false,           // no-wait
		nil,             // args
	)
	if err != nil {
		log.Fatalf("Failed to register payout consumer: %v", err)
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

				var payloadMap map[string]interface{}
				if err := json.Unmarshal(d.Body, &payloadMap); err != nil {
					log.Printf("Error decoding payout message: %v\n", err)
					d.Reject(false)
					continue
				}

				coreTxId, _ := payloadMap["core_tx_id"].(string)
				amount, _ := payloadMap["amount"].(string)
				dappId, _ := payloadMap["dapp_id"].(string) // passed from listener via OutboxEvent

				log.Printf("Processing Payout for Tx %s, Amount: %s, Dapp: %s\n", coreTxId, amount, dappId)

				var order domain.PayoutOrder
				if err := db.DB.Where("core_tx_id = ?", coreTxId).First(&order).Error; err != nil {
					log.Printf("Payout API Failed. PayoutOrder not found for Tx %s.\n", coreTxId)
					d.Reject(false)
					continue
				}

				// Retrieve sensitive Bank Info from Redis instead of DB
				redisKey := "payout_bank:" + coreTxId
				bankInfoStr, err := db.RedisClient.Get(context.Background(), redisKey).Result()
				if err != nil {
					log.Printf("Payout API Failed. Bank info missing/expired in Redis for Tx %s.\n", coreTxId)
					db.DB.Model(&order).Update("status", "FAILED")
					d.Reject(false)
					continue
				}

				var bankInfo map[string]string
				json.Unmarshal([]byte(bankInfoStr), &bankInfo)

				mockBankPayload, _ := json.Marshal(map[string]string{
					"core_tx_id":   coreTxId,
					"amount":       amount,
					"type":         "PAYOUT",
					"bank_account": bankInfo["bank_account"],
					"bank_bin":     bankInfo["bank_bin"],
				})

				resp, err := http.Post("https://core-banking-internal/api/payout", "application/json", bytes.NewBuffer(mockBankPayload))
				
				if err != nil || resp.StatusCode >= 300 {
					log.Printf("Payout API Failed for Tx %s. Sending to DLQ or alerting admin.\n", coreTxId)
					// In real world: send alert, maybe refund.
					db.DB.Model(&order).Update("status", "FAILED")
					d.Reject(false) // Send to DLQ
				} else {
					log.Printf("Payout SUCCESS for Tx %s\n", coreTxId)
					
					db.DB.Model(&order).Update("status", "COMPLETED")
					db.RedisClient.Del(context.Background(), redisKey) // Clean up Redis
					d.Ack(false)
				}
			}
		}
	}()

	fmt.Println("Payout Worker started listening for messages.")
}
