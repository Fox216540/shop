package dto

import "github.com/google/uuid"

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

type ProductResponse struct {
	ProductID          uuid.UUID `json:"id"`
	ProductName        string    `json:"name"`
	ProductPrice       float64   `json:"price"`
	ProductCategoryID  uuid.UUID `json:"category_id"`
	ProductDescription string    `json:"description"`
}

type ItemsResponse struct {
	Product  []ProductResponse `json:"products"`
	Quantity int               `json:"quantity"`
}

type OrderResponse struct {
	ID       uuid.UUID       `json:"id"`
	UserID   uuid.UUID       `json:"user_id"`
	OrderNum string          `json:"order_num"`
	Total    float64         `json:"total"`
	Status   string          `json:"status"`
	Items    []ItemsResponse `json:"items"`
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

type TestCreateOrderRequest struct {
	ID         string        `json:"id" binding:"required,uuid"`
	OrderItems []ItemRequest `json:"product_items" binding:"required"`
}

type TestDeleteOrderRequest struct {
	UserID string `json:"userID" binding:"required,uuid"`
	ID     string `json:"id" binding:"required,uuid"`
}
