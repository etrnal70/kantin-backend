package model

import (
	"github.com/gofrs/uuid"
	"github.com/jackc/pgtype"
)

type ImagePath struct {
	ID   uuid.UUID `json:"id"`
	Path string    `json:"image_path"`
}

type MenuItem struct {
	ID          uuid.UUID   `json:"id"`
	Seller_ID   uuid.UUID   `json:"seller_id"`
	Image       ImagePath   `json:"image"`
	Name        string      `json:"name"`
	Description pgtype.Text `json:"description"`
	Category    int         `json:"category"`
	Price       int         `json:"price"`
}
