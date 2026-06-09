package product

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Product struct {
	ID          uuid.UUID       // Product ID
	Name        string          // Product name
	Img         string          // Product image URL
	Amount      decimal.Decimal // Product amount
	Currency    string          // Product currency
	CategoryID  uuid.UUID       // Product category
	Description string          // Product description
	Stock       uint64          // Product stock quantity
}
