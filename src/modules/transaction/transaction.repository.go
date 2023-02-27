package transaction

import (
	"context"

	uuid "github.com/gofrs/uuid"
	"github.com/hanherb/mini-wallet/src/config"
	"gorm.io/gorm"
)

func NewRepository(model Transaction) *TransactionRepositoryImp {
	return &TransactionRepositoryImp{
		Model: model,
		Query: config.DB,
	}
}

type TransactionRepositoryImp struct {
	Model Transaction
	Query *gorm.DB
}

type TransactionGetProps struct {
	WalletId    *uuid.UUID
	Type        *string
	ReferenceID *uuid.UUID
}

type TransactionRepository interface {
	FindOne(ctx context.Context, props *TransactionGetProps) (data *Transaction, err error)
	FindMany(ctx context.Context, props *TransactionGetProps) (data []*Transaction, err error)
	AddTransaction(ctx context.Context, req *Transaction) (data *Transaction, err error)
}

func (r *TransactionRepositoryImp) FindOne(ctx context.Context, props *TransactionGetProps) (data *Transaction, err error) {
	query := r.Query
	if props.ReferenceID != nil {
		query = query.Where("reference_id = ?", *props.ReferenceID)
	}

	if err = query.Find(&data).Error; err != nil {
		if err.Error() == "record not found" {
			return nil, nil
		}
	}
	return
}

func (r *TransactionRepositoryImp) FindMany(ctx context.Context, props *TransactionGetProps) (data []*Transaction, err error) {
	query := r.Query
	if props.WalletId != nil {
		query = query.Where("wallet_id = ?", *props.WalletId)
	}
	if props.Type != nil {
		query = query.Where("type = ?", *props.Type)
	}

	if err = query.Find(&data).Error; err != nil {
		if err.Error() == "record not found" {
			return nil, nil
		}
	}
	return
}

func (r *TransactionRepositoryImp) AddTransaction(ctx context.Context, req *Transaction) (data *Transaction, err error) {
	err = r.Query.Create(req).Error
	if err != nil {
		return
	}

	data, err = r.FindOne(ctx, &TransactionGetProps{ReferenceID: &req.ReferenceID})
	return
}
