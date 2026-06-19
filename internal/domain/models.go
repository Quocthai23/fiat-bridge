package domain

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
	ID        uint   `gorm:"primaryKey"`
	EventType string `gorm:"not null;default:'RABBITMQ'"` // "RABBITMQ" or "WEBHOOK"
	Payload   string `gorm:"type:text;not null"`
	Status    string `gorm:"not null;default:'PENDING'"`
}

// SyncState tracks the last synced block for different events
type SyncState struct {
	ID            uint   `gorm:"primaryKey"`
	LastMintBlock uint64 `gorm:"not null;default:0"`
	LastBurnBlock uint64 `gorm:"not null;default:0"`
}
