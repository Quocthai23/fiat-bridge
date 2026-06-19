package queue

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

				// Find DApp config to get Bank Account to transfer TO (Wait, this is payout to user? No, payout to the User's bank account. Wait, the user provides their bank account somewhere? Or maybe we just mock it for now since we don't have user bank account).
				// We'll mock the Bank API call
				
				mockBankPayload, _ := json.Marshal(map[string]string{
					"core_tx_id": coreTxId,
					"amount":     amount,
					"type":       "PAYOUT",
				})

				resp, err := http.Post("https://core-banking-internal/api/payout", "application/json", bytes.NewBuffer(mockBankPayload))
				
				if err != nil || resp.StatusCode >= 300 {
					log.Printf("Payout API Failed for Tx %s. Sending to DLQ or alerting admin.\n", coreTxId)
					// In real world: send alert, maybe refund.
					d.Reject(false) // Send to DLQ
				} else {
					log.Printf("Payout SUCCESS for Tx %s\n", coreTxId)
					
					// Update DB if we had a Payout table. For now just Ack.
					d.Ack(false)
				}
			}
		}
	}()

	fmt.Println("Payout Worker started listening for messages.")
}
