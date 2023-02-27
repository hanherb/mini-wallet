package transaction

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

type Transaction struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	WalletID     uuid.UUID `json:"wallet_id"`
	Status       string    `json:"status"`
	TransactedAt time.Time `json:"transacted_at"`
	Type         string    `json:"type"`
	Amount       int       `json:"amount"`
	ReferenceID  uuid.UUID `json:"reference_id"`
}

func (Transaction) TableName() string {
	return "transactions"
}
