package wallet

import "github.com/gofrs/uuid"

type ReqUpdateBalance struct {
	WalletID uuid.UUID
	Amount   int
}
