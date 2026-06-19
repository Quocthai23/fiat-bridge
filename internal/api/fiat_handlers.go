package api

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/Quocthai23/fiat-bridge/internal/domain"
	"github.com/Quocthai23/fiat-bridge/internal/infrastructure/db"
	"github.com/Quocthai23/fiat-bridge/internal/infrastructure/queue"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func generateCoreTxId() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("FIAT-%d", rand.Intn(1000000))
}

// HandleCreateFiatOrder handles DApp requests to create a deposit order
func HandleCreateFiatOrder(c *gin.Context) {
	var req domain.CreateFiatOrderRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dappId := c.GetString("dappId")
	coreTxId := generateCoreTxId()

	var config domain.DappConfig
	if err := db.DB.First(&config, "id = ?", dappId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DApp config not found"})
		return
	}

	order := domain.FiatOrder{
		CoreTxId:   coreTxId,
		DappId:     dappId,
		Amount:     req.Amount,
		Wallet:     req.UserAddress,
		Status:     domain.FiatStatusWaiting,
		WebhookUrl: config.WebhookUrl,
	}

	if err := db.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	// Generate VietQR URL using DappConfig
	qrUrl := fmt.Sprintf(
		"https://img.vietqr.io/image/%s-%s-compact2.png?amount=%s&addInfo=%s&accountName=%s",
		config.BankBin,
		config.BankAccount,
		order.Amount,
		order.CoreTxId,
		config.AccountName,
	)

	c.JSON(http.StatusOK, gin.H{
		"core_tx_id": order.CoreTxId,
		"qr_url":     qrUrl,
		"status":     order.Status,
	})
}

// verifyBankSignature mocks checking the webhook signature from the bank
func verifyBankSignature(payload domain.BankWebhookPayload) bool {
	// For production, implement actual signature verification
	// e.g. HMAC-SHA256 with a shared secret
	expectedSecret := "secret_key"
	mac := hmac.New(sha256.New, []byte(expectedSecret))
	mac.Write([]byte(payload.Amount + payload.Description))
	expectedSignature := hex.EncodeToString(mac.Sum(nil))

	// Allow bypass for testing if signature is "test_signature"
	if payload.Signature == "test_signature" {
		return true
	}

	return payload.Signature == expectedSignature
}

// extractCoreTxId finds the CoreTxId in the transfer description
func extractCoreTxId(description string) string {
	// Improve regex to handle bank garbage characters
	re := regexp.MustCompile(`(?i)FIAT-\d+`)
	match := re.FindString(description)
	return strings.ToUpper(match)
}

// HandleBankWebhook handles incoming payment notifications from the bank
func HandleBankWebhook(c *gin.Context) {
	var payload domain.BankWebhookPayload
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !verifyBankSignature(payload) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid signature"})
		return
	}

	coreTxId := extractCoreTxId(payload.Description)
	if coreTxId == "" {
		// Not a valid transaction for our bridge
		c.JSON(http.StatusOK, gin.H{"message": "Ignored: No CoreTxId found"})
		return
	}

	// Use Database Transaction to prevent concurrency issues (Race Condition Fix)
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		var order domain.FiatOrder

		// Lock the row for update to prevent concurrent webhooks processing the same order
		if err := tx.Set("gorm:query_option", "FOR UPDATE").
			Where("core_tx_id = ? AND status = ?", coreTxId, domain.FiatStatusWaiting).
			First(&order).Error; err != nil {
			return fmt.Errorf("Order not found or already processed")
		}

		// Update order status
		order.Status = domain.FiatStatusPaid
		if err := tx.Save(&order).Error; err != nil {
			return err
		}

		// Create Mint Transaction (1:1 conversion for this MVP)
		mintTx := domain.Transaction{
			CoreTxId:    order.CoreTxId,
			UserAddress: order.Wallet,
			Amount:      order.Amount, // 1 VND = 1 Token
			Type:        domain.TxTypeMint,
			Status:      domain.StatusPendingProcessing,
		}

		if err := tx.Create(&mintTx).Error; err != nil {
			// This will fail if a Transaction with the same CoreTxId already exists due to unique index
			return err
		}

		// Push to RabbitMQ (Using existing worker logic)
		if err := queue.PublishMintTransaction(&mintTx); err != nil {
			return fmt.Errorf("failed to publish to queue: %v", err)
		}

		return nil
	})

	if err != nil {
		if strings.Contains(err.Error(), "Order not found or already processed") {
			// Return 200 so bank doesn't retry
			c.JSON(http.StatusOK, gin.H{"message": "Order already processed or invalid"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Webhook processed successfully"})
}

// HandleCreatePayoutOrder handles DApp requests to create a payout order (Off-ramp)
func HandleCreatePayoutOrder(c *gin.Context) {
	var req domain.CreatePayoutOrderRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dappId := c.GetString("dappId")
	coreTxId := generateCoreTxId() // reuse the same generator

	order := domain.PayoutOrder{
		CoreTxId:    coreTxId,
		DappId:      dappId,
		Amount:      req.Amount,
		Wallet:      req.UserAddress,
		Status:      "WAITING_FOR_BURN",
	}

	if err := db.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create payout order"})
		return
	}

	// Save sensitive bank information to Redis (Ephemeral storage) instead of Database
	if db.RedisClient != nil {
		ctx := c.Request.Context()
		redisKey := "payout_bank:" + coreTxId
		bankInfo, _ := json.Marshal(map[string]string{
			"bank_account": req.BankAccount,
			"bank_bin":     req.BankBin,
		})
		db.RedisClient.Set(ctx, redisKey, bankInfo, 24*time.Hour)
	}

	c.JSON(http.StatusOK, gin.H{
		"core_tx_id": order.CoreTxId,
		"status":     order.Status,
	})
}
