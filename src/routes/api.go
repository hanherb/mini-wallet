package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hanherb/mini-wallet/src/middlewares"
	"github.com/hanherb/mini-wallet/src/modules"
)

func StartRoute() *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.CORS())
	api := router.Group("/api/v1")
	{
		api.GET("/", modules.HealthCheck)
	}

	go customerRoutes(api)
	go walletRoutes(api)

	return router
}
