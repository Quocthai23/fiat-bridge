package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Quocthai23/fiat-bridge/internal/domain"
	"github.com/Quocthai23/fiat-bridge/internal/infrastructure/db"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// HandleMintCommand handles the mint request from Core Banking
func HandleMintCommand(c *gin.Context) {
	var req domain.MintRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !common.IsHexAddress(req.UserAddress) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid EVM address format"})
		return
	}

	// Save transaction and outbox event to DB within a single transaction
	tx := &domain.Transaction{
		CoreTxId:    req.CoreTxId,
		UserAddress: req.UserAddress,
		Amount:      req.Amount,
		Type:        domain.TxTypeMint,
		Status:      domain.StatusPendingProcessing,
	}

	err := db.DB.Transaction(func(dbTx *gorm.DB) error {
		if err := dbTx.Create(tx).Error; err != nil {
			return err
		}

		payloadBytes, _ := json.Marshal(tx)
		outbox := domain.OutboxEvent{
			Payload: string(payloadBytes),
			Status:  "PENDING",
		}
		return dbTx.Create(&outbox).Error
	})

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			c.JSON(http.StatusConflict, gin.H{"error": "Transaction already processed"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save transaction"})
		return
	}

	fmt.Printf("Tx Published to Queue: Lock Fiat %s for User %s.\n", req.Amount, req.UserAddress)

	c.JSON(http.StatusOK, gin.H{
		"status":     "pending_on_chain",
		"core_tx_id": req.CoreTxId,
		"message":    "Mint transaction published to queue",
	})
}

// HandleBurnCommand handles the burn request from Core Banking
func HandleBurnCommand(c *gin.Context) {
	var req domain.BurnRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !common.IsHexAddress(req.UserAddress) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid EVM address format"})
		return
	}

	// Save transaction and outbox event to DB within a single transaction
	tx := &domain.Transaction{
		CoreTxId:    req.CoreTxId,
		UserAddress: req.UserAddress,
		Amount:      req.Amount,
		Type:        domain.TxTypeBurn,
		Status:      domain.StatusPendingProcessing,
	}

	err := db.DB.Transaction(func(dbTx *gorm.DB) error {
		if err := dbTx.Create(tx).Error; err != nil {
			return err
		}

		payloadBytes, _ := json.Marshal(tx)
		outbox := domain.OutboxEvent{
			Payload: string(payloadBytes),
			Status:  "PENDING",
		}
		return dbTx.Create(&outbox).Error
	})

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			c.JSON(http.StatusConflict, gin.H{"error": "Transaction already processed"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save transaction"})
		return
	}

	fmt.Printf("Tx Published to Queue: Unlock Fiat %s for User %s.\n", req.Amount, req.UserAddress)

	c.JSON(http.StatusOK, gin.H{
		"status":     "pending_on_chain",
		"core_tx_id": req.CoreTxId,
		"message":    "Burn transaction published to queue",
	})
}
