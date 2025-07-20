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
	ID         uuid.UUID // Unique identifier for the order
	OrderNum   string    // Unique order number
	UserID     uuid.UUID // User ID who placed the order
	OrderItems []*Item   // List of products in the order
	Total      float64   // Total order amount
	Status     string    // Order status
}
