package wallet

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

type Wallet struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	OwnedBy   uuid.UUID `json:"owned_by"`
	Status    string    `json:"status"`
	EnabledAt time.Time `json:"enabled_at"`
	Balance   int       `json:"balance"`
}

func (Wallet) TableName() string {
	return "wallets"
}
