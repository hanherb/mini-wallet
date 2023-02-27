package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/hanherb/mini-wallet/src/helper"
	"github.com/hanherb/mini-wallet/src/lib"
)

func NewCustomerController(repo CustomerRepository) *CustomerControllerImp {
	return &CustomerControllerImp{
		Repo: repo,
	}
}

type CustomerControllerImp struct {
	Repo CustomerRepository
}

func (ctr *CustomerControllerImp) InitCustomer(c *gin.Context) {
	var req ReqInitCustomer
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseISE(err.Error()))
		return
	}

	id := uuid.Must(uuid.FromString(req.CustomerXid))

	// check if exist
	customer, err := ctr.Repo.FindOne(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseISE(err.Error()))
		return
	}

	// if not exist
	if customer == nil {
		reqCustomer := &Customer{
			ID: id,
		}
		if err := ctr.Repo.Create(c, reqCustomer); err != nil {
			c.JSON(http.StatusInternalServerError, helper.ResponseISE(err.Error()))
			return
		}
	}

	token, err := lib.CreateAccessToken(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseISE(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseSuccess(gin.H{"token": token}))
}
