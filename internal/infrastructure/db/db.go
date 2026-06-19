package db

import (
	"fmt"
	"log"
	"time"

	"github.com/Quocthai23/fiat-bridge/internal/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database connection and runs migrations
func InitDB(dsn string) {
	var err error

	// Retry connection a few times in case the database is still starting up
	for i := 0; i < 5; i++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database. Retrying in 2 seconds... (%v/5)\n", i+1)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatalf("Failed to connect to database after 5 retries: %v", err)
	}

	fmt.Println("Connected to PostgreSQL successfully!")

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get sqlDB: %v", err)
	}

	// Limit connections
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Auto-migrate the schema
	err = DB.AutoMigrate(
		&domain.Transaction{},
		&domain.AddressNonce{},
		&domain.FailedNonce{},
		&domain.OutboxEvent{},
		&domain.SyncState{},
		&domain.FiatOrder{},
		&domain.DappConfig{},
		&domain.PayoutOrder{},
	)
	if err != nil {
		log.Fatalf("Failed to auto-migrate database schema: %v", err)
	}
	fmt.Println("Database schema migrated successfully!")
}



// SaveTransaction saves a new transaction to the database
func SaveTransaction(tx *domain.Transaction) error {
	return DB.Create(tx).Error
}

// UpdateTransaction updates an existing transaction
func UpdateTransaction(tx *domain.Transaction) error {
	return DB.Save(tx).Error
}
