package order

import (
	"github.com/google/uuid"
	"shop/src/domain/product"
)

type Item struct {
	Product  product.Product
	Quantity int
}

type Order struct {
	ID           uuid.UUID `json:"id"`
	OrderNum     string    `json:"order-num" validate:"required"` // Unique order number
	UserId       uuid.UUID `json:"user-id" validate:"required"`
	ProductItems []*Item   `json:"product-items" validate:"required"` // List of products in the order
	Status       string    `json:"status"`
}
