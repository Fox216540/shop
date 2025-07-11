package user

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id" binding:"required"`
	Username string    `json:"username" binding:"required,min=3,max=20"`
	Email    string    `json:"email" binding:"required,email"`
	Name     string    `json:"name" binding:"required,min=3,max=20"`
	Number   string    `json:"number" binding:"required,min=3,max=20"`
	Password string    `json:"password" binding:"required"`
}
