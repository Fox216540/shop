package models

import (
	dom "github.com/Fox216540/shop/payment-service/domain/payment"
	"github.com/google/uuid"
	pay "github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
	"github.com/shopspring/decimal"
)

type PaymentORM struct {
	ID          int `gorm:"primaryKey;autoIncrement"`
	IDYuKassa   string
	OrderID     uuid.UUID
	Value       decimal.Decimal
	Currency    string
	Method      string
	Description string
	Status      pay.Status
}

func (PaymentORM) TableName() string {
	return "payments"
}

func FromORM(orm PaymentORM) dom.Payment {
	o := dom.Payment{
		ID:      orm.IDYuKassa,
		OrderID: orm.OrderID,
		Amount: dom.Amount{
			Value:    orm.Value,
			Currency: orm.Currency,
		},
		Method:      orm.Method,
		Description: orm.Description,
		Status:      orm.Status,
	}
	return o
}
