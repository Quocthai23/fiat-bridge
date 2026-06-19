package api

import (
	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes the API routes
func SetupRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		bridge := v1.Group("/bridge")
		{
			bridge.POST("/mint", HandleMintCommand)
			bridge.POST("/burn", HandleBurnCommand)
		}

		fiat := v1.Group("/fiat")
		{
			fiat.POST("/orders", HandleCreateFiatOrder)
		}

		webhooks := v1.Group("/webhooks")
		{
			webhooks.POST("/bank", HandleBankWebhook)
		}
	}
}
