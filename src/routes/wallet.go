package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hanherb/mini-wallet/src/middlewares"
	"github.com/hanherb/mini-wallet/src/modules/transaction"
	"github.com/hanherb/mini-wallet/src/modules/wallet"
)

func walletRoutes(api *gin.RouterGroup) {
	walletModule := wallet.NewModule()
	transactionModule := transaction.NewModule()

	walletRoute := api.Group("/wallet")
	{
		walletRoute.GET("/", middlewares.AuthCustomer, walletModule.Controller.ViewBalance)
		walletRoute.POST("/", middlewares.AuthCustomer, walletModule.Controller.EnableWallet)
		walletRoute.PATCH("/", middlewares.AuthCustomer, walletModule.Controller.DisableWallet)
		walletRoute.GET("/transactions", middlewares.AuthCustomer, transactionModule.Controller.GetListTransaction)
		walletRoute.POST("/deposits", middlewares.AuthCustomer, transactionModule.Controller.Deposit)
		walletRoute.POST("/withdrawals", middlewares.AuthCustomer, transactionModule.Controller.Withdraw)
	}
}
