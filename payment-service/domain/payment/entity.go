package payment

import (
	"github.com/google/uuid"
	pay "github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
	"github.com/shopspring/decimal"
)

type Amount struct {
	Value    decimal.Decimal
	Currency string
}

type Payment struct {
	ID          string
	OrderID     uuid.UUID
	Amount      Amount
	Method      string
	ReturnURL   string
	Description string
	Status      pay.Status
}
