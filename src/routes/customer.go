package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hanherb/mini-wallet/src/modules/customer"
)

func customerRoutes(api *gin.RouterGroup) {
	module := customer.NewModule()

	customerRoute := api.Group("/")
	{
		customerRoute.POST("/init", module.Controller.InitCustomer)
	}
}
