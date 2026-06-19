package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Quocthai23/fiat-bridge/internal/api"
	"github.com/Quocthai23/fiat-bridge/internal/domain"
	"github.com/Quocthai23/fiat-bridge/internal/infrastructure/db"
	"github.com/gin-gonic/gin"
)

// TestOnRampFlow simulates a POST request to /api/v1/fiat/orders
func TestOnRampFlow(t *testing.T) {
	testDb := setupTestDB(t)
	if testDb == nil {
		return
	}
	db.DB = testDb

	// Seed DappConfig
	testDb.Create(&domain.DappConfig{
		ID:          "test-api-key",
		Name:        "Test DApp",
		BankBin:     "970436",
		BankAccount: "123456789",
		AccountName: "TEST CORP",
	})

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	api.SetupRoutes(router)

	payload := map[string]interface{}{
		"dapp_id":      "test-api-key",
		"user_address": "0x1234567890abcdef1234567890abcdef12345678",
		"amount":       "100000",
		"webhook_url":  "https://dapp.example.com/webhook",
	}

	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/api/v1/fiat/orders", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-KEY", "test-api-key")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %v: %s", w.Code, w.Body.String())
	}

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	if response["qr_url"] == nil || response["core_tx_id"] == nil {
		t.Errorf("Response missing qr_url or core_tx_id: %v", response)
	}

	var order domain.FiatOrder
	testDb.First(&order, "core_tx_id = ?", response["core_tx_id"])
	if order.Status != "WAITING_FOR_PAYMENT" {
		t.Errorf("Expected status WAITING_FOR_PAYMENT, got %s", order.Status)
	}
}

// TestOffRampFlow checks the payout order creation endpoint
func TestOffRampFlow(t *testing.T) {
	testDb := setupTestDB(t)
	if testDb == nil {
		return
	}
	db.DB = testDb

	// Seed DappConfig
	testDb.Create(&domain.DappConfig{
		ID: "test-api-key",
	})

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	api.SetupRoutes(router)

	payload := map[string]interface{}{
		"user_address": "0x12345",
		"bank_account": "99999999",
		"bank_bin":     "970407",
		"amount":       "500000",
	}

	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/api/v1/fiat/payout-orders", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-KEY", "test-api-key")
	
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %v", w.Code)
	}

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	var order domain.PayoutOrder
	testDb.First(&order, "core_tx_id = ?", response["core_tx_id"])
	if order.Status != "WAITING_FOR_BURN" {
		t.Errorf("Expected status WAITING_FOR_BURN, got %s", order.Status)
	}
}
