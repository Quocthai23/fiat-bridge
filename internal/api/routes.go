package api

import (
	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes the API routes
func SetupRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		// Protected Routes
		protected := v1.Group("/")
		protected.Use(RequireApiKey())
		{
			bridge := protected.Group("/bridge")
			{
				bridge.POST("/mint", HandleMintCommand)
				bridge.POST("/burn", HandleBurnCommand)
			}

			fiat := protected.Group("/fiat")
			{
				fiat.POST("/orders", HandleCreateFiatOrder)
			}
		}

		// Unprotected Webhook Routes
		webhooks := v1.Group("/webhooks")
		{
			webhooks.POST("/bank", HandleBankWebhook)
		}
	}
}
