package queue

import (
	"encoding/json"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/Quocthai23/fiat-bridge/internal/domain"
)

var RabbitConn *amqp.Connection
var RabbitChannel *amqp.Channel

const MintQueueName = "mint_transactions"
const BurnQueueName = "burn_transactions"
const PayoutQueueName = "payout_queue"
const DLQName = "mint_transactions_dlq"
const DLXName = "dlx_exchange"

// InitRabbitMQ initializes the RabbitMQ connection and channel
func InitRabbitMQ(url string) {
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	RabbitChannel, err = conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}

	// 1. Declare the Dead Letter Exchange
	err = RabbitChannel.ExchangeDeclare(
		DLXName,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare DLX: %v", err)
	}

	// 2. Declare the Dead Letter Queue
	_, err = RabbitChannel.QueueDeclare(
		DLQName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare DLQ: %v", err)
	}

	// 3. Bind the DLQ to the DLX
	err = RabbitChannel.QueueBind(
		DLQName,
		"dlq_routing_key", // routing key
		DLXName,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to bind DLQ: %v", err)
	}

	// 4. Declare the main queue with DLX args
	args := amqp.Table{
		"x-dead-letter-exchange":    DLXName,
		"x-dead-letter-routing-key": "dlq_routing_key",
	}

	_, err = RabbitChannel.QueueDeclare(
		MintQueueName, // name
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		args,          // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare mint queue: %v", err)
	}

	_, err = RabbitChannel.QueueDeclare(
		BurnQueueName, // name
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		args,          // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare burn queue: %v", err)
	}

	_, err = RabbitChannel.QueueDeclare(
		PayoutQueueName, // name
		true,            // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		args,            // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare payout queue: %v", err)
	}

	fmt.Println("Connected to RabbitMQ successfully!")
}

// PublishMintTransaction publishes a transaction to the RabbitMQ queue
func PublishMintTransaction(tx *domain.Transaction) error {
	body, err := json.Marshal(tx)
	if err != nil {
		return err
	}

	err = RabbitChannel.Publish(
		"",            // exchange
		MintQueueName, // routing key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         body,
			DeliveryMode: amqp.Persistent, // persistent message
		})
	return err
}

// PublishBurnTransaction publishes a transaction to the RabbitMQ burn queue
func PublishBurnTransaction(tx *domain.Transaction) error {
	body, err := json.Marshal(tx)
	if err != nil {
		return err
	}

	err = RabbitChannel.Publish(
		"",            // exchange
		BurnQueueName, // routing key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         body,
			DeliveryMode: amqp.Persistent, // persistent message
		})
	return err
}
