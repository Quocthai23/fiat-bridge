package domain

import "time"

// MintRequest represents the payload from Core Banking
type MintRequest struct {
	CoreTxId    string `json:"core_tx_id" binding:"required"`
	UserAddress string `json:"user_address" binding:"required"`
	Amount      string `json:"amount" binding:"required"`
}

// BurnRequest represents the burn payload from Core Banking
type BurnRequest struct {
	CoreTxId    string `json:"core_tx_id" binding:"required"`
	UserAddress string `json:"user_address" binding:"required"`
	Amount      string `json:"amount" binding:"required"`
}

// OutboxEvent represents an event to be published to a message broker or webhook
type OutboxEvent struct {
	ID          uint       `gorm:"primaryKey"`
	EventType   string     `gorm:"not null;default:'RABBITMQ'"` // "RABBITMQ" or "WEBHOOK" or "PAYOUT"
	Payload     string     `gorm:"type:text;not null"`
	Status      string     `gorm:"not null;default:'PENDING'"`
	RetryCount  int        `gorm:"not null;default:0"`
	NextRetryAt *time.Time // Nullable, if null execute immediately
}

// SyncState tracks the last synced block for different events
type SyncState struct {
	ID            uint   `gorm:"primaryKey"`
	LastMintBlock uint64 `gorm:"not null;default:0"`
	LastBurnBlock uint64 `gorm:"not null;default:0"`
}

// DappConfig stores the banking info and webhook config for each DApp (White-label)
type DappConfig struct {
	ID            string `gorm:"primaryKey"` // dApp API Key
	Name          string
	BankBin       string // VD: "970436" (Vietcombank)
	BankAccount   string // VD: "1012345678"
	AccountName   string // VD: "NGUYEN QUOC THAI"
	WebhookUrl    string // URL để Bridge bắn kết quả về cho DApp
	WebhookSecret string // Secret để ký HMAC gửi cho DApp
}
