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
	// get customer
	customerId, err := lib.TokenCustomerId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseISE(err.Error()))
		return
	}

	// get wallet
	wallet, err := ctr.Repo.FindOne(c, &WalletGetProps{CustomerId: &customerId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseISE(err.Error()))
		return
	}
	if wallet != nil { // enable disabled wallet
		if wallet.Status == "enabled" {
			c.JSON(http.StatusForbidden, helper.ResponseErrorMessage("fail", "Already enabled"))
			return
		} else {
			wallet, err = ctr.Repo.Enable(c, wallet.ID)

		}
	} else { // add wallet if not exist
		reqWallet := &Wallet{
			OwnedBy:   customerId,
			Status:    "enabled",
			EnabledAt: time.Now(),
		}

		wallet, err = ctr.Repo.AddWallet(c, reqWallet)
		if err != nil {
			c.JSON(http.StatusInternalServerError, helper.ResponseISE(err.Error()))
			return
		}
	}

	c.JSON(http.StatusOK, helper.ResponseSuccess(gin.H{"wallet": wallet}))
}

func (ctr *WalletControllerImp) DisableWallet(c *gin.Context) {
	// get customer
	customerId, err := lib.TokenCustomerId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseISE(err.Error()))
		return
	}

	// get wallet
	wallet, err := ctr.Repo.FindOne(c, &WalletGetProps{CustomerId: &customerId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseISE(err.Error()))
		return
	}

	if wallet == nil { // if already disabled
		c.JSON(http.StatusForbidden, helper.ResponseErrorMessage("fail", "Wallet disabled"))
		return
	}

	// disable enabled wallet
	wallet, err = ctr.Repo.Disable(c, wallet.ID)

	c.JSON(http.StatusOK, helper.ResponseSuccess(gin.H{"wallet": wallet}))
}

func (ctr *WalletControllerImp) ViewBalance(c *gin.Context) {
	// get cusomter
	customerId, err := lib.TokenCustomerId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseISE(err.Error()))
		return
	}

	// get enabled wallet
	wallet, err := ctr.Repo.FindOne(c, &WalletGetProps{CustomerId: &customerId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseISE(err.Error()))
		return
	}

	if wallet == nil || wallet.Status != "enabled" {
		c.JSON(http.StatusForbidden, helper.ResponseErrorMessage("fail", "Wallet disabled"))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseSuccess(gin.H{"wallet": wallet}))
}
