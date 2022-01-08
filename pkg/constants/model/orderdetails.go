package model

type OrderDetails struct {
	ID          uint64 `json:"id"`
	Menuitem_ID uint64 `json:"menuitem_id"`
	Quantity    int    `json:"quantity"`
	Comment     string `json:"comment"`
}
