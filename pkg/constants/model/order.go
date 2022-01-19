package model

import (
	"time"

	"github.com/gofrs/uuid"
)

type Order struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	SellerID  uuid.UUID `json:"seller_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}
