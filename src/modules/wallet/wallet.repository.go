package wallet

import (
	"context"

	uuid "github.com/gofrs/uuid"
	"github.com/hanherb/mini-wallet/src/config"
	"gorm.io/gorm"
)

func NewRepository(model Wallet) *WalletRepositoryImp {
	return &WalletRepositoryImp{
		Model: model,
		Query: config.DB,
	}
}

type WalletRepositoryImp struct {
	Model Wallet
	Query *gorm.DB
}

type WalletGetProps struct {
	WalletId   *uuid.UUID
	CustomerId *uuid.UUID
}

type WalletRepository interface {
	FindOne(ctx context.Context, props *WalletGetProps) (data *Wallet, err error)
	AddWallet(ctx context.Context, req *Wallet) (data *Wallet, err error)
	Enable(ctx context.Context, walletId uuid.UUID) (data *Wallet, err error)
	Disable(ctx context.Context, walletId uuid.UUID) (data *Wallet, err error)
	UpdateBalance(ctx context.Context, req *ReqUpdateBalance) (data *Wallet, err error)
}

func (r *WalletRepositoryImp) FindOne(ctx context.Context, props *WalletGetProps) (data *Wallet, err error) {
	query := r.Query
	if props.WalletId != nil {
		query = query.Where("id = ?", *props.WalletId)
	}
	if props.CustomerId != nil {
		query = query.Where("owned_by = ?", *props.CustomerId)
	}

	if err = query.First(&data).Error; err != nil {
		if err.Error() == "record not found" {
			return nil, nil
		}
	}
	return
}

func (r *WalletRepositoryImp) AddWallet(ctx context.Context, req *Wallet) (data *Wallet, err error) {
	err = r.Query.Create(req).Error
	if err != nil {
		return
	}

	data, err = r.FindOne(ctx, &WalletGetProps{CustomerId: &req.OwnedBy})
	return
}

func (r *WalletRepositoryImp) Enable(ctx context.Context, walletId uuid.UUID) (data *Wallet, err error) {
	err = r.Query.Updates(&Wallet{
		ID:     walletId,
		Status: "enabled",
	}).Error
	if err != nil {
		return
	}

	data, err = r.FindOne(ctx, &WalletGetProps{WalletId: &walletId})
	return
}

func (r *WalletRepositoryImp) Disable(ctx context.Context, walletId uuid.UUID) (data *Wallet, err error) {
	err = r.Query.Updates(&Wallet{
		ID:     walletId,
		Status: "disabled",
	}).Error
	if err != nil {
		return
	}

	data, err = r.FindOne(ctx, &WalletGetProps{WalletId: &walletId})
	return
}

func (r *WalletRepositoryImp) UpdateBalance(ctx context.Context, req *ReqUpdateBalance) (data *Wallet, err error) {
	err = r.Query.Updates(&Wallet{
		ID:      req.WalletID,
		Balance: req.Amount,
	}).Error
	if err != nil {
		return
	}

	data, err = r.FindOne(ctx, &WalletGetProps{WalletId: &req.WalletID})
	return
}
