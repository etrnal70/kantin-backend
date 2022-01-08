package model

import "time"

type Order struct {
	ID        uint64    `json:"id"`
	UserID    uint64    `json:"user_id"`
	SellerID  uint64    `json:"seller_id"`
	Orderlist []int64   `json:"orderlist"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}
