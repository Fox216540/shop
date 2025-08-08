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
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Message  string    `json:"message"`
}

type TestTokensResponse struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

type UserWithTokensResponse struct {
	ID       uuid.UUID          `json:"id"`
	Username string             `json:"username"`
	Tokens   TestTokensResponse `json:"tokens"`
	Message  string             `json:"message"`
}

// Request

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
}

type TestDeleteRequest struct {
	ID string `json:"id" binding:"required,uuid"`
}

type ItemRequest struct {
	ProductID string `json:"product_id" binding:"required,uuid"`
	Quantity  int    `json:"quantity" binding:"required"`
}

type TestCreateOrderRequest struct {
	ID         string        `json:"id" binding:"required,uuid"`
	OrderItems []ItemRequest `json:"order_items" binding:"required"`
}

type TestDeleteOrderRequest struct {
	UserID string `json:"user_id" binding:"required,uuid"`
	ID     string `json:"id" binding:"required,uuid"`
}

type TestGetOrdersRequest struct {
	UserID string `json:"user_id" binding:"required,uuid"`
}

type TestLoginRequest struct {
	UsernameOrEmail string `json:"username_or_email" binding:"required"`
	Password        string `json:"password" binding:"required"`
}

type TestLogoutRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type TestUpdateUsernameRequest struct {
	ID       string `json:"id" binding:"required,uuid"`
	Username string `json:"username" binding:"required"`
}

type TestUpdatePasswordRequest struct {
	ID       string `json:"id" binding:"required,uuid"`
	Password string `json:"password" binding:"required"`
}

type TestUpdateEmailRequest struct {
	ID    string `json:"id" binding:"required,uuid"`
	Email string `json:"email" binding:"required,email"`
}

type TestUpdatePhoneRequest struct {
	ID    string `json:"id" binding:"required,uuid"`
	Phone string `json:"phone" binding:"required"`
}

type TestUpdateProfileRequest struct {
	ID      string `json:"id" binding:"required,uuid"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
