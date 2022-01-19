package model

import (
	"github.com/gofrs/uuid"
	"github.com/jackc/pgtype"
)

type Seller struct {
	ID       uuid.UUID   `json:"id"`
	Name     string      `json:"name"`
	Email    string      `json:"email"`
	Password string      `json:"password"`
	Logo     pgtype.Text `json:"logo"`
}

type SellerLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SellerRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SellerPublic struct {
	ID   uuid.UUID   `json:"id"`
	Name string      `json:"name"`
	Logo pgtype.Text `json:"logo"`
}
