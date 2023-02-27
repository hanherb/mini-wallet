package transaction

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/gofrs/uuid"
	"github.com/hanherb/mini-wallet/src/helper"
	"github.com/hanherb/mini-wallet/src/lib"
	"github.com/hanherb/mini-wallet/src/modules/wallet"
)

func NewTransactionController(repo TransactionRepository) *TransactionControllerImp {
	return &TransactionControllerImp{
		Repo: repo,
	}
}

type TransactionControllerImp struct {
	Repo TransactionRepository
}

func (ctr *TransactionControllerImp) GetListTransaction(c *gin.Context) {
	// get customer & wallet
	customerId, err := lib.TokenCustomerId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseISE(err.Error()))
		return
	}

	walletRepository := wallet.NewRepository(wallet.Wallet{})
	wallet, err := walletRepository.FindOne(c, &wallet.WalletGetProps{CustomerId: &customerId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseISE(err.Error()))
		return
	}

	// check wallet status
	if wallet == nil || wallet.Status != "enabled" {
		c.JSON(http.StatusForbidden, helper.ResponseErrorMessage("fail", "Wallet disabled"))
		return
	}

	// get transaction list
	transactions, err := ctr.Repo.FindMany(c, &TransactionGetProps{WalletId: &wallet.ID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseISE(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseSuccess(gin.H{"transactions": transactions}))
}

func (ctr *TransactionControllerImp) Deposit(c *gin.Context) {
	var req ReqAddTransaction
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseISE(err.Error()))
		return
	}

	// get customer & wallet
	customerId, err := lib.TokenCustomerId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseISE(err.Error()))
		return
	}

	walletRepository := wallet.NewRepository(wallet.Wallet{})
	walletData, err := walletRepository.FindOne(c, &wallet.WalletGetProps{CustomerId: &customerId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseISE(err.Error()))
		return
	}

	// check wallet status
	if walletData == nil || walletData.Status != "enabled" {
		c.JSON(http.StatusForbidden, helper.ResponseErrorMessage("fail", "Wallet disabled"))
		return
	}

	// generate reference_id
	referenceId, err := uuid.NewV4()
	if err != nil {
		return
	}

	// add transaction
	reqTransaction := &Transaction{
		WalletID:     walletData.ID,
		Status:       "success",
		TransactedAt: time.Now(),
		Type:         "deposit",
		Amount:       req.Amount,
		ReferenceID:  referenceId,
	}

	transaction, err := ctr.Repo.AddTransaction(c, reqTransaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseISE(err.Error()))
		return
	}

	// update wallet balance
	reqUpdateBalance := &wallet.ReqUpdateBalance{
		WalletID: walletData.ID,
		Amount:   walletData.Balance + req.Amount,
	}

	_, err = walletRepository.UpdateBalance(c, reqUpdateBalance)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseISE(err.Error()))
		return
	}

	// transform to deposit response
	deposit := transformResponseDeposit(transaction, customerId)

	c.JSON(http.StatusOK, helper.ResponseSuccess(gin.H{"deposit": deposit}))
}

func (ctr *TransactionControllerImp) Withdraw(c *gin.Context) {
	var req ReqAddTransaction
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseISE(err.Error()))
		return
	}

	// get cusomter & wallet
	customerId, err := lib.TokenCustomerId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseISE(err.Error()))
		return
	}

	walletRepository := wallet.NewRepository(wallet.Wallet{})
	walletData, err := walletRepository.FindOne(c, &wallet.WalletGetProps{CustomerId: &customerId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseISE(err.Error()))
		return
	}

	// check wallet status
	if walletData == nil || walletData.Status != "enabled" {
		c.JSON(http.StatusForbidden, helper.ResponseErrorMessage("fail", "Wallet disabled"))
		return
	}

	// check balance
	if (walletData.Balance - req.Amount) < 0 {
		c.JSON(http.StatusForbidden, helper.ResponseErrorMessage("fail", "Insufficient balance"))
		return
	}

	// generate reference_id
	referenceId, err := uuid.NewV4()
	if err != nil {
		return
	}

	// add transaction
	reqTransaction := &Transaction{
		WalletID:     walletData.ID,
		Status:       "success",
		TransactedAt: time.Now(),
		Type:         "withdraw",
		Amount:       req.Amount,
		ReferenceID:  referenceId,
	}

	transaction, err := ctr.Repo.AddTransaction(c, reqTransaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseISE(err.Error()))
		return
	}

	// update wallet balance
	reqUpdateBalance := &wallet.ReqUpdateBalance{
		WalletID: walletData.ID,
		Amount:   walletData.Balance - req.Amount,
	}

	_, err = walletRepository.UpdateBalance(c, reqUpdateBalance)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ResponseISE(err.Error()))
		return
	}

	// transform to withdraw response
	withdraw := transformResponseWithdraw(transaction, customerId)

	c.JSON(http.StatusOK, helper.ResponseSuccess(gin.H{"withdrawal": withdraw}))
}

func transformResponseDeposit(transaction *Transaction, customerId uuid.UUID) (res *ResDeposit) {
	res = &ResDeposit{
		ID:          transaction.ID,
		WalletID:    transaction.WalletID,
		DepositedBy: customerId,
		Status:      transaction.Status,
		DepositedAt: transaction.TransactedAt,
		Type:        transaction.Type,
		Amount:      transaction.Amount,
		ReferenceID: transaction.ReferenceID,
	}

	return
}

func transformResponseWithdraw(transaction *Transaction, customerId uuid.UUID) (res *ResWithdraw) {
	res = &ResWithdraw{
		ID:          transaction.ID,
		WalletID:    transaction.WalletID,
		WithdrawnBy: customerId,
		Status:      transaction.Status,
		WithdrawnAt: transaction.TransactedAt,
		Type:        transaction.Type,
		Amount:      transaction.Amount,
		ReferenceID: transaction.ReferenceID,
	}

	return
}
