package user

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id" validate:"required"`
	Username string    `json:"username" validate:"required,min=3,max=20"`
	Email    string    `json:"email" validate:"required,email"`
	Name     string    `json:"name" validate:"required,min=3,max=20"`
	Number   string    `json:"number" validate:"required,min=3,max=20"`
	Password string    `json:"password" validate:"required"`
}
