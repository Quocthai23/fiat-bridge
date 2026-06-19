package api

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// RequireApiKey middleware checks for a valid X-API-KEY header
func RequireApiKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-KEY")
		
		// For this MVP, we validate against a master DAPP_API_KEY environment variable.
		// In a full production scenario, we would query the database to find the DApp
		// associated with this API key.
		expectedKey := os.Getenv("DAPP_API_KEY")
		if expectedKey == "" {
			// If no key is configured on the server, deny access by default for safety
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Server API key not configured"})
			return
		}

		if apiKey == "" || apiKey != expectedKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized DApp"})
			return
		}

		// Store a mock DApp ID in the context. In a real scenario, this would be the actual ID from DB.
		c.Set("dappId", "dapp-authenticated")
		
		c.Next()
	}
}
