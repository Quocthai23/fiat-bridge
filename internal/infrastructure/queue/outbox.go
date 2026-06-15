package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
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
