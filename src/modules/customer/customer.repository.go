package customer

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/hanherb/mini-wallet/src/config"
	"gorm.io/gorm"
)

func NewRepository(model Customer) *CustomerRepositoryImp {
	return &CustomerRepositoryImp{
		Model: model,
		Query: config.DB,
	}
}

type CustomerRepositoryImp struct {
	Model Customer
	Query *gorm.DB
}

type CustomerRepository interface {
	FindOne(ctx context.Context, id uuid.UUID) (data *Customer, err error)
	Create(ctx context.Context, req *Customer) error
}

func (r *CustomerRepositoryImp) FindOne(ctx context.Context, id uuid.UUID) (data *Customer, err error) {
	if err = r.Query.Where("id = ?", id).First(&data).Error; err != nil {
		if err.Error() == "record not found" {
			return nil, nil
		}
	}
	return
}

func (r *CustomerRepositoryImp) Create(ctx context.Context, req *Customer) (err error) {
	err = r.Query.Create(req).Error
	return
}
