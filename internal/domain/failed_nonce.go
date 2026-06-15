package domain

import (
	"time"
)

// FailedNonce represents a nonce that was acquired but failed to be sent,
// and thus can be reused to prevent nonce gaps.
type FailedNonce struct {
	ID        uint   `gorm:"primaryKey"`
	Address   string `gorm:"index;not null"`
	Nonce     uint64 `gorm:"not null"`
	CreatedAt time.Time
}
