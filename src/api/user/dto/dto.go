package dto

import (
	"github.com/google/uuid"
)

// TODO: Отсортировать DTO

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

type CreateOrderResponse struct {
	ID       uuid.UUID `json:"id"`
	UserID   uuid.UUID `json:"user_id"`
	OrderNum string    `json:"order_num"`
	Total    float64   `json:"total"`
	Status   string    `json:"status"`
}

type TestDeleteRequest struct {
	ID string `json:"id" binding:"required,uuid"`
}

type ItemRequest struct {
	ProductID string `json:"product_id" binding:"required,uuid"`
	Quantity  int    `json:"quantity" binding:"required"`
}

// TODO: Поменять product_items на order_items
type TestCreateOrderRequest struct {
	ID         string        `json:"id" binding:"required,uuid"`
	OrderItems []ItemRequest `json:"product_items" binding:"required"`
}

// TODO: Поменять userID на user_id
type TestDeleteOrderRequest struct {
	UserID string `json:"userID" binding:"required,uuid"`
	ID     string `json:"id" binding:"required,uuid"`
}

type TestGetOrdersRequest struct {
	UserID string `json:"userID" binding:"required,uuid"`
}
