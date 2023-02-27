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

type WalletRepository interface {
	FindOne(ctx context.Context, customerId uuid.UUID) (data *Wallet, err error)
	Enable(ctx context.Context, req *Wallet) (data *Wallet, err error)
}

func (r *WalletRepositoryImp) FindOne(ctx context.Context, customerId uuid.UUID) (data *Wallet, err error) {
	if err = r.Query.Where("owned_by = ?", customerId).First(&data).Error; err != nil {
		if err.Error() == "record not found" {
			return nil, nil
		}
	}
	return
}

func (r *WalletRepositoryImp) Enable(ctx context.Context, req *Wallet) (data *Wallet, err error) {
	err = r.Query.Create(req).Error
	if err != nil {
		return
	}

	data, err = r.FindOne(ctx, req.OwnedBy)
	return
}
