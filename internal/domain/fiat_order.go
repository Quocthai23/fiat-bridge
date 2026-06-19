package domain

import (
	"time"

	"gorm.io/gorm"
)

// FiatOrder status constants
const (
	FiatStatusWaiting = "WAITING_FOR_PAYMENT"
	FiatStatusPaid    = "PAID"
	FiatStatusFailed  = "FAILED"
)

// FiatOrder represents a deposit request from a DApp
type FiatOrder struct {
	ID         uint   `gorm:"primaryKey"`
	CoreTxId   string `gorm:"uniqueIndex;not null"`
	DappId     string `gorm:"not null"`
	Amount     string `gorm:"not null"`
	Wallet     string `gorm:"not null"`
	Status     string `gorm:"not null;default:'WAITING_FOR_PAYMENT'"`
	WebhookUrl string `gorm:"not null"` // Where to send the success callback
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

// CreateFiatOrderRequest represents the API payload from DApp
type CreateFiatOrderRequest struct {
	DappId      string `json:"dapp_id" binding:"required"`
	UserAddress string `json:"user_address" binding:"required"`
	Amount      string `json:"amount" binding:"required"` // Fiat amount (e.g. VND)
	WebhookUrl  string `json:"webhook_url" binding:"required"`
}

// BankWebhookPayload represents the expected payload from the bank/payment gateway
type BankWebhookPayload struct {
	Amount      string `json:"amount"`
	Description string `json:"description"`
	Signature   string `json:"signature"` // Used to verify authenticity
}

// DappWebhookPayload represents the payload sent back to the DApp
type DappWebhookPayload struct {
	CoreTxId string `json:"core_tx_id"`
	TxHash   string `json:"tx_hash"`
	Status   string `json:"status"`
	Amount   string `json:"amount"`
}

// PayoutOrder represents an off-ramp payout request
type PayoutOrder struct {
	ID          uint   `gorm:"primaryKey"`
	CoreTxId    string `gorm:"uniqueIndex;not null"`
	DappId      string `gorm:"not null"`
	Amount      string `gorm:"not null"`
	Wallet      string `gorm:"not null"` // user_address
	Status      string `gorm:"not null;default:'WAITING_FOR_BURN'"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

// CreatePayoutOrderRequest represents the API payload from DApp for off-ramp
type CreatePayoutOrderRequest struct {
	UserAddress string `json:"user_address" binding:"required"`
	BankAccount string `json:"bank_account" binding:"required"`
	BankBin     string `json:"bank_bin" binding:"required"`
	Amount      string `json:"amount" binding:"required"`
}
