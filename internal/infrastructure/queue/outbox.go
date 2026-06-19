package queue

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Quocthai23/fiat-bridge/internal/domain"
	"github.com/Quocthai23/fiat-bridge/internal/infrastructure/db"
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// StartOutboxRelay starts a background worker that polls the OutboxEvent table
// for PENDING events and publishes them to RabbitMQ.
func StartOutboxRelay(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}

			err := db.DB.Transaction(func(tx *gorm.DB) error {
				var events []domain.OutboxEvent
				// Find up to 50 pending events with SKIP LOCKED
				if err := tx.Clauses(clause.Locking{Strength: "UPDATE", Options: "SKIP LOCKED"}).
					Where("status = ?", "PENDING").Limit(50).Find(&events).Error; err != nil {
					return err
				}

				for _, ev := range events {
					if ev.EventType == "WEBHOOK" {
						// Dispatch via HTTP POST
						// The Payload is expected to be a JSON object that includes a "webhook_url" field,
						// or the webhook URL could be passed inside the payload and extracted.
						// For simplicity, let's assume the payload contains the webhook url.
						var payloadMap map[string]interface{}
						if err := json.Unmarshal([]byte(ev.Payload), &payloadMap); err != nil {
							log.Printf("Failed to unmarshal webhook payload for outbox event %d: %v\n", ev.ID, err)
							tx.Model(&ev).Update("status", "FAILED")
							continue
						}

						webhookUrl, ok := payloadMap["webhook_url"].(string)
						if !ok || webhookUrl == "" {
							log.Printf("No webhook_url found in payload for outbox event %d\n", ev.ID)
							tx.Model(&ev).Update("status", "FAILED")
							continue
						}

						// Remove webhook_url from the payload sent to DApp
						delete(payloadMap, "webhook_url")
						cleanPayload, _ := json.Marshal(payloadMap)

						resp, err := http.Post(webhookUrl, "application/json", bytes.NewBuffer(cleanPayload))
						if err == nil && resp.StatusCode >= 200 && resp.StatusCode < 300 {
							tx.Model(&ev).Update("status", "SENT")
						} else {
							statusCode := 0
							if resp != nil {
								statusCode = resp.StatusCode
							}
							log.Printf("Failed to send webhook to %s (Status: %d): %v\n", webhookUrl, statusCode, err)
							// For simplicity in MVP, we might mark as failed or leave as pending for retry
							// Since it's SKIP LOCKED, leaving it pending will cause it to be retried
							// But to avoid infinite loops, let's mark it as FAILED if we want simple behavior,
							// or better yet, implement a retry count. For now, leave it PENDING for endless retry.
						}
					} else {
						// Default to RABBITMQ
						var domTx domain.Transaction
						queueName := MintQueueName
						if err := json.Unmarshal([]byte(ev.Payload), &domTx); err == nil {
							if domTx.Type == domain.TxTypeBurn {
								queueName = BurnQueueName
							}
						}

						// Publish directly to RabbitMQ using the raw JSON payload
						err := RabbitChannel.Publish("", queueName, false, false, amqp.Publishing{
							ContentType:  "application/json",
							Body:         []byte(ev.Payload),
							DeliveryMode: amqp.Persistent,
						})

						// If successfully published, mark as SENT
						if err == nil {
							tx.Model(&ev).Update("status", "SENT")
						} else {
							log.Printf("Failed to publish outbox event %d to queue: %v\n", ev.ID, err)
						}
					}
				}
				return nil
			})
			if err != nil {
				log.Printf("Error processing outbox events: %v\n", err)
			}

			select {
			case <-ctx.Done():
				return
			case <-time.After(2 * time.Second):
			}
		}
	}()

	fmt.Println("Outbox Relay started polling for pending events.")
}
