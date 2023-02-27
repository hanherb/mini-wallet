package customer

import "github.com/gofrs/uuid"

type Customer struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
}

func (Customer) TableName() string {
	return "customers"
}
