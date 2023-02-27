package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hanherb/mini-wallet/src/middlewares"
	"github.com/hanherb/mini-wallet/src/modules/wallet"
)

func walletRoutes(api *gin.RouterGroup) {
	module := wallet.NewModule()

	walletRoute := api.Group("/")
	{
		walletRoute.POST("/wallet", middlewares.AuthCustomer, module.Controller.EnableWallet)
	}
}
