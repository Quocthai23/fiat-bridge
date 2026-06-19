package queue

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
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
				now := time.Now()
				// Find up to 50 pending events with SKIP LOCKED
				if err := tx.Clauses(clause.Locking{Strength: "UPDATE", Options: "SKIP LOCKED"}).
					Where("status = ? AND (next_retry_at IS NULL OR next_retry_at <= ?)", "PENDING", now).Limit(50).Find(&events).Error; err != nil {
					return err
				}

				for _, ev := range events {
					if ev.EventType == "WEBHOOK" {
						var payloadMap map[string]interface{}
						if err := json.Unmarshal([]byte(ev.Payload), &payloadMap); err != nil {
							log.Printf("Failed to unmarshal webhook payload for outbox event %d: %v\n", ev.ID, err)
							tx.Model(&ev).Update("status", "FAILED")
							continue
						}

						webhookUrl, _ := payloadMap["webhook_url"].(string)
						webhookSecret, _ := payloadMap["webhook_secret"].(string)

						if webhookUrl == "" {
							log.Printf("No webhook_url found in payload for outbox event %d\n", ev.ID)
							tx.Model(&ev).Update("status", "FAILED")
							continue
						}

						// Remove internal fields
						delete(payloadMap, "webhook_url")
						delete(payloadMap, "webhook_secret")
						cleanPayload, _ := json.Marshal(payloadMap)

						req, err := http.NewRequest("POST", webhookUrl, bytes.NewBuffer(cleanPayload))
						if err != nil {
							handleWebhookFailure(tx, &ev, err)
							continue
						}
						req.Header.Set("Content-Type", "application/json")

						// Generate HMAC-SHA256 signature
						if webhookSecret != "" {
							mac := hmac.New(sha256.New, []byte(webhookSecret))
							mac.Write(cleanPayload)
							signature := hex.EncodeToString(mac.Sum(nil))
							req.Header.Set("X-Bridge-Signature", signature)
						}

						client := &http.Client{Timeout: 10 * time.Second}
						resp, err := client.Do(req)

						if err == nil && resp.StatusCode >= 200 && resp.StatusCode < 300 {
							tx.Model(&ev).Update("status", "SENT")
						} else {
							handleWebhookFailure(tx, &ev, fmt.Errorf("HTTP error or non-2xx status"))
						}
					} else if ev.EventType == "PAYOUT" {
						err := RabbitChannel.Publish("", PayoutQueueName, false, false, amqp.Publishing{
							ContentType:  "application/json",
							Body:         []byte(ev.Payload),
							DeliveryMode: amqp.Persistent,
						})
						if err == nil {
							tx.Model(&ev).Update("status", "SENT")
						} else {
							log.Printf("Failed to publish outbox event %d to payout queue: %v\n", ev.ID, err)
							nextRetry := time.Now().Add(5 * time.Second)
							tx.Model(&ev).Updates(map[string]interface{}{"retry_count": ev.RetryCount + 1, "next_retry_at": nextRetry})
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
							nextRetry := time.Now().Add(5 * time.Second)
							tx.Model(&ev).Updates(map[string]interface{}{"retry_count": ev.RetryCount + 1, "next_retry_at": nextRetry})
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

func handleWebhookFailure(tx *gorm.DB, ev *domain.OutboxEvent, err error) {
	log.Printf("Webhook delivery failed for event %d: %v", ev.ID, err)
	ev.RetryCount++
	
	if ev.RetryCount > 5 {
		tx.Model(ev).Updates(map[string]interface{}{
			"status": "FAILED",
		})
		return
	}

	// Exponential backoff: 5s, 1m, 5m, 1h, 5h
	var delay time.Duration
	switch ev.RetryCount {
	case 1:
		delay = 5 * time.Second
	case 2:
		delay = 1 * time.Minute
	case 3:
		delay = 5 * time.Minute
	case 4:
		delay = 1 * time.Hour
	default:
		delay = 5 * time.Hour
	}

	nextRetry := time.Now().Add(delay)
	tx.Model(ev).Updates(map[string]interface{}{
		"retry_count":   ev.RetryCount,
		"next_retry_at": nextRetry,
	})
}
