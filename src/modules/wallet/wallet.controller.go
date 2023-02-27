package wallet

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hanherb/mini-wallet/src/helper"
	"github.com/hanherb/mini-wallet/src/lib"
)

func NewWalletController(repo WalletRepository) *WalletControllerImp {
	return &WalletControllerImp{
		Repo: repo,
	}
}

type WalletControllerImp struct {
	Repo WalletRepository
}

func (ctr *WalletControllerImp) EnableWallet(c *gin.Context) {
	customerId, err := lib.TokenCustomerId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseISE(err.Error()))
		return
	}

	wallet, err := ctr.Repo.FindOne(c, customerId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseISE(err.Error()))
		return
	}
	if wallet != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"status": "fail",
			"data": gin.H{
				"error": "Already enabled",
			},
		})
		return
	}

	reqWallet := &Wallet{
		OwnedBy:   customerId,
		Status:    "enabled",
		EnabledAt: time.Now(),
	}

	wallet, err = ctr.Repo.Enable(c, reqWallet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseISE(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseSuccess(gin.H{"wallet": wallet}))
}
