package domain

import (
	"time"
)

// AddressNonce tracks the nonce for an Ethereum address
type AddressNonce struct {
	Address   string `gorm:"primaryKey"`
	Nonce     uint64 `gorm:"not null;default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
