package model

import (
	"github.com/gofrs/uuid"
	"github.com/jackc/pgtype"
)

type User struct {
	ID        uuid.UUID   `json:"id"`
	Firstname string      `json:"firstname"`
	Lastname  pgtype.Text `json:"lastname"`
	Email     string      `json:"email"`
	Password  string      `json:"password"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRegister struct {
	Firstname string      `json:"firstname"`
	Lastname  pgtype.Text `json:"lastname"`
	Email     string      `json:"email"`
	Password  string      `json:"password"`
}
