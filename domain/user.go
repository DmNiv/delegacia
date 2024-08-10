package domain

import (
	"time"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type User struct {
	Id           uuid.UUID `json:"id" gorm:"primaryKey"`
	FirstName    string    `json:"firstName" validate:"required"`
	LastName     string    `json:"lastName" validate:"required"`
	Email        string    `json:"email" validate:"required"`
	Password     string    `json:"password" validate:"required" gorm:"-"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

func (User) TableName() string {
	return "delegacia_facil.users"
}

type UserLoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
