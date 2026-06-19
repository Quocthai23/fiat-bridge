package api

import (
	"context"
	"net/http"
	"time"

	"github.com/Quocthai23/fiat-bridge/internal/domain"
	"github.com/Quocthai23/fiat-bridge/internal/infrastructure/db"
	"github.com/gin-gonic/gin"
)

// RequireApiKey middleware checks for a valid X-API-KEY header
func RequireApiKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-KEY")
		if apiKey == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing X-API-KEY"})
			return
		}

		var config domain.DappConfig
		if err := db.DB.First(&config, "id = ?", apiKey).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized DApp"})
			return
		}

		c.Set("dappId", config.ID)
		c.Next()
	}
}

// RateLimiterMiddleware uses Redis to limit requests per API key
func RateLimiterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-KEY")
		if apiKey == "" {
			c.Next()
			return
		}

		// Skip if redis is not configured
		if db.RedisClient == nil {
			c.Next()
			return
		}

		ctx := context.Background()
		key := "rate_limit:" + apiKey

		// Token bucket / Window approach: 50 requests per 10 seconds
		// Using Pipeline to ensure atomicity
		pipe := db.RedisClient.Pipeline()
		incr := pipe.Incr(ctx, key)
		pipe.Expire(ctx, key, 10*time.Second)
		
		_, err := pipe.Exec(ctx)
		if err != nil {
			c.Next()
			return
		}

		count := incr.Val()

		if count > 50 {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded. Try again later."})
			return
		}

		c.Next()
	}
}
