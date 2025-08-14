package dto

import (
	"github.com/google/uuid"
)

// Response
type MessageResponse struct {
	Message string `json:"message"`
}

type CreateOrderResponse struct {
	ID       uuid.UUID `json:"id"`
	UserID   uuid.UUID `json:"user_id"`
	OrderNum string    `json:"order_num"`
	Total    float64   `json:"total"`
	Status   string    `json:"status"`
}
type UserResponse struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Message string    `json:"message"`
}

type TestTokensResponse struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

type UserWithTokensResponse struct {
	ID      uuid.UUID          `json:"id"`
	Name    string             `json:"name"`
	Tokens  TestTokensResponse `json:"tokens"`
	Message string             `json:"message"`
}

// Request

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
}

type ItemRequest struct {
	ProductID string `json:"product_id" binding:"required,uuid"`
	Quantity  int    `json:"quantity" binding:"required"`
}

type CreateOrderRequest struct {
	OrderItems []ItemRequest `json:"order_items" binding:"required"`
}

type DeleteOrderRequest struct {
	OrderID string `json:"order_id" binding:"required,uuid"`
}

type LoginRequest struct {
	PhoneOrEmail string `json:"phone_or_email" binding:"required"`
	Password     string `json:"password" binding:"required"`
}

type LogoutRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type UpdatePasswordRequest struct {
	NewPassword string `json:"new_password" binding:"required"`
}

type UpdateEmailRequest struct {
	NewEmail string `json:"new_email" binding:"required,email"`
}

type UpdatePhoneRequest struct {
	NewPhone string `json:"new_phone" binding:"required"`
}

type UpdateProfileRequest struct {
	NewName    string `json:"new_name"`
	NewAddress string `json:"new_address"`
}
