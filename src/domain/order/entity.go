package order

import "github.com/google/uuid"

type ProductItem struct {
	ProductID uuid.UUID
	Quantity  int
}

type Order struct {
	ID           uuid.UUID     `json:"id"`
	OrderNum     string        `json:"order-num" binding:"required"` // Unique order number
	UserId       uuid.UUID     `json:"user-id" binding:"required"`
	ProductItems []ProductItem `json:"product_items" binding:"required"` // List of products in the order
	Status       string        `json:"status"`
}
