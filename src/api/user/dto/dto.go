package dto

import "github.com/google/uuid"

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Address  string `json:"address" binding:"required"`
}

type UserResponse struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Message  string    `json:"message"`
}

type TestDeleteRequest struct {
	ID string `json:"id" binding:"required,uuid"`
}
