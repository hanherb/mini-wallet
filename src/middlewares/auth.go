package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hanherb/mini-wallet/src/helper"
	"github.com/hanherb/mini-wallet/src/lib"
)

func AuthCustomer(c *gin.Context) {
	jwt, err := lib.ExtractToken(c.Request)

	if err != nil {
		c.JSON(http.StatusUnauthorized, helper.ResponseUnauthorized(err.Error()))
		return
	}

	c.Set("jwtId", jwt.ID)
	c.Next()
}
