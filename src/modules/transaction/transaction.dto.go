package transaction

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

type ReqAddTransaction struct {
	Amount int `json:"amount" binding:"required"`
}

type ResDeposit struct {
	ID          uuid.UUID `json:"id"`
	WalletID    uuid.UUID `json:"wallet_id"`
	DepositedBy uuid.UUID `json:"deposited_by"`
	Status      string    `json:"status"`
	DepositedAt time.Time `json:"deposited_at"`
	Type        string    `json:"type"`
	Amount      int       `json:"amount"`
	ReferenceID uuid.UUID `json:"reference_id"`
}

type ResWithdraw struct {
	ID          uuid.UUID `json:"id"`
	WalletID    uuid.UUID `json:"wallet_id"`
	WithdrawnBy uuid.UUID `json:"withdrawn_by"`
	Status      string    `json:"status"`
	WithdrawnAt time.Time `json:"withdrawn_at"`
	Type        string    `json:"type"`
	Amount      int       `json:"amount"`
	ReferenceID uuid.UUID `json:"reference_id"`
}
