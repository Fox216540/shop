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
	ID         uuid.UUID
	OrderNum   string
	UserID     uuid.UUID
	OrderItems []Item
	Total      decimal.Decimal
	Status     string
	Currency   string
}
