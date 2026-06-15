package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/Quocthai23/fiat-bridge/internal/api"
	"github.com/Quocthai23/fiat-bridge/internal/infrastructure/blockchain"
	"github.com/Quocthai23/fiat-bridge/internal/infrastructure/db"
	"github.com/Quocthai23/fiat-bridge/internal/infrastructure/kms"
	"github.com/Quocthai23/fiat-bridge/internal/infrastructure/nonce"
	"github.com/Quocthai23/fiat-bridge/internal/infrastructure/queue"
	"github.com/Quocthai23/fiat-bridge/internal/infrastructure/reconciliation"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

var client *ethclient.Client

func main() {
	// 0. Initialize Database & Queue
	dsn := "host=localhost user=root password=secretpassword dbname=bridge_db port=54320 sslmode=disable"
	db.InitDB(dsn)

	rabbitUrl := "amqp://guest:guest@localhost:56720/"
	queue.InitRabbitMQ(rabbitUrl)

	var err error
	client, err = ethclient.Dial("https://rpc.sepolia.org")
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum node: %v", err)
	}
	fmt.Println("Connected to Ethereum RPC successfully!")

	// 1. Initialize Web3 Adapter (Nonce Manager, Emitter, Listener)
	nm := nonce.NewDBNonceManager(client, db.DB)

	contractAddrHex := os.Getenv("CONTRACT_ADDRESS")
	if contractAddrHex == "" {
		contractAddrHex = "0x0000000000000000000000000000000000000000"
	}

	var signer kms.Signer
	kmsKeyID := os.Getenv("KMS_KEY_ID")

	if kmsKeyID != "" {
		fmt.Println("Initializing AWS KMS Signer...")
		cfg, err := awsconfig.LoadDefaultConfig(context.Background())
		if err != nil {
			log.Fatalf("Failed to load AWS config: %v", err)
		}
		signer, err = kms.NewAWSKMSSigner(context.Background(), cfg, kmsKeyID)
		if err != nil {
			log.Fatalf("Failed to initialize AWS KMS Signer: %v", err)
		}
	} else {
		fmt.Println("Initializing Mock KMS Signer...")
		privKeyHex := os.Getenv("PRIVATE_KEY_HEX")
		if privKeyHex == "" {
			privKeyHex = "0000000000000000000000000000000000000000000000000000000000000001"
		}
		signer, err = kms.NewMockKMSSigner(privKeyHex)
		if err != nil {
			log.Fatalf("Failed to initialize Mock KMS Signer: %v", err)
		}
	}

	emitter, err := blockchain.NewEmitter(client, signer, contractAddrHex, nm)
	if err != nil {
		log.Fatalf("Failed to initialize Emitter: %v", err)
	}

	listener, err := blockchain.NewListener(client, contractAddrHex)
	if err != nil {
		log.Fatalf("Failed to initialize Listener: %v", err)
	}
	
	appCtx, appCancel := context.WithCancel(context.Background())
	defer appCancel()

	listener.StartListening(appCtx, 11155000)

	// Initialize Reconciliation Engine
	reconEngine, err := reconciliation.NewEngine(client, contractAddrHex)
	if err != nil {
		log.Fatalf("Failed to initialize Reconciliation Engine: %v", err)
	}
	reconEngine.Start(appCtx, 1*time.Minute) // Run every minute for testing

	// Initialize Gas Bumper Worker
	gasBumper := blockchain.NewGasBumper(emitter)
	gasBumper.Start(appCtx, 5*time.Minute) // Scan every 5 minutes

	// Initialize Consumer Worker
	queue.StartConsumerWorker(appCtx, emitter)
	queue.StartBurnConsumerWorker(appCtx, emitter)

	// Initialize Outbox Worker
	queue.StartOutboxRelay(appCtx)

	r := gin.Default()
	api.SetupRoutes(r)

	fmt.Println("Inbound Gateway is running on port 8080...")
	
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server gracefully...")

	// Cancel all background workers
	appCancel()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
