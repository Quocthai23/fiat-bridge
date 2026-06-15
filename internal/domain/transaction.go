package domain

import (
	"time"

	"gorm.io/gorm"
)

// Transaction status constants
const (
	StatusPendingProcessing = "pending_processing"
	StatusPendingOnChain   = "pending_on_chain"
	StatusCompleted        = "completed"
	StatusFailed           = "failed"
	StatusFailedRollbacked = "failed_rollbacked"
)

// Transaction type constants
const (
	TxTypeMint = "mint"
	TxTypeBurn = "burn"
)

// Transaction represents a lock-and-mint request from Core Bank
type Transaction struct {
	ID          uint   `gorm:"primaryKey"`
	CoreTxId    string `gorm:"uniqueIndex;not null"`
	UserAddress string `gorm:"not null"`
	Amount      string `gorm:"not null"`
	Type        string `gorm:"not null;default:'mint'"` // "mint" or "burn"
	Status      string `gorm:"not null;default:'pending_on_chain'"`
	TxHash      string // Blockchain transaction hash
	Nonce       uint64 // Transaction nonce for bumping
	BlockNumber uint64 // Block number when the transaction was completed
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
