package catalog

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Product struct {
	ID    uuid.UUID
	Name  string
	Img   string
	Price Money
}

type Money struct {
	Amount   decimal.Decimal
	Currency string
}
