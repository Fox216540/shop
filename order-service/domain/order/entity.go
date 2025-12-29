package order

import (
	"github.com/Fox216540/shop/order-service/domain/catalog"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Item struct {
	Product  catalog.Product
	Quantity uint64
}

type Order struct {
	ID         uuid.UUID       // Unique identifier for the order
	OrderNum   string          // Unique order number
	UserID     uuid.UUID       // User ID who placed the order
	OrderItems []Item          // List of products in the order
	Total      decimal.Decimal // Total order amount
	Currency   string
	Status     string // Order status
}
