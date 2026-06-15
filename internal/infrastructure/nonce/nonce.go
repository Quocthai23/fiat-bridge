package nonce

import (
	"context"
	"fmt"
	"strings"

	"github.com/Quocthai23/fiat-bridge/internal/domain"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Manager handles the assignment of nonces for Ethereum transactions
type Manager interface {
	GetAndIncrementNonce(ctx context.Context, address common.Address) (uint64, error)
	ReleaseNonce(address common.Address, nonce uint64)
	ResetNonce(address common.Address)
}

// DBNonceManager implements the Manager interface using a database table to store the state safely across multiple pods
type DBNonceManager struct {
	client *ethclient.Client
	db     *gorm.DB
}

// NewDBNonceManager creates a database-backed nonce manager
func NewDBNonceManager(client *ethclient.Client, db *gorm.DB) *DBNonceManager {
	return &DBNonceManager{
		client: client,
		db:     db,
	}
}

// GetAndIncrementNonce fetches the current nonce from the DB (with row-level lock) and increments it atomically
func (m *DBNonceManager) GetAndIncrementNonce(ctx context.Context, address common.Address) (uint64, error) {
	addrHex := strings.ToLower(address.Hex())
	var assignedNonce uint64

	err := m.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var failedNonce domain.FailedNonce

		// Try to grab an available failed nonce first
		err := tx.Clauses(clause.Locking{Strength: "UPDATE", Options: "SKIP LOCKED"}).
			Where("address = ?", addrHex).Order("nonce ASC").First(&failedNonce).Error

		if err == nil {
			assignedNonce = failedNonce.Nonce
			// Delete it from failed_nonces so it's not reused again
			return tx.Delete(&failedNonce).Error
		}

		var nonceRecord domain.AddressNonce

		// Lock the row for update (SELECT ... FOR UPDATE)
		err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("address = ?", addrHex).First(&nonceRecord).Error

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				// If not found, fetch from blockchain
				chainNonce, err := m.client.PendingNonceAt(ctx, address)
				if err != nil {
					return fmt.Errorf("failed to get pending nonce from chain: %v", err)
				}
				nonceRecord = domain.AddressNonce{
					Address: addrHex,
					Nonce:   chainNonce,
				}
				if err := tx.Create(&nonceRecord).Error; err != nil {
					return err
				}
			} else {
				return err
			}
		}

		assignedNonce = nonceRecord.Nonce

		// Increment nonce in DB
		return tx.Model(&nonceRecord).Update("nonce", assignedNonce+1).Error
	})

	if err != nil {
		return 0, fmt.Errorf("failed to get and increment nonce: %v", err)
	}

	return assignedNonce, nil
}

// ReleaseNonce saves a failed nonce into the DB to be reused later
func (m *DBNonceManager) ReleaseNonce(address common.Address, nonce uint64) {
	addrHex := strings.ToLower(address.Hex())
	failedNonce := domain.FailedNonce{
		Address: addrHex,
		Nonce:   nonce,
	}
	m.db.Create(&failedNonce)
}

// ResetNonce resets the nonce to the current on-chain state
func (m *DBNonceManager) ResetNonce(address common.Address) {
	addrHex := strings.ToLower(address.Hex())
	chainNonce, err := m.client.PendingNonceAt(context.Background(), address)
	if err == nil {
		m.db.Model(&domain.AddressNonce{}).Where("address = ?", addrHex).Update("nonce", chainNonce)
	}
}
