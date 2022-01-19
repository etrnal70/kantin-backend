package model

import (
	"github.com/gofrs/uuid"
	"github.com/jackc/pgtype"
)

type OrderDetails struct {
	ID          uuid.UUID   `json:"id"`
	OrderID     uuid.UUID   `json:"order_id"`
	Menuitem_ID uuid.UUID   `json:"menuitem_id"`
	Quantity    int         `json:"quantity"`
	Comment     pgtype.Text `json:"comment"`
}
