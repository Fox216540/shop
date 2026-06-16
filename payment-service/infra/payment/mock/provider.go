package mock

import (
	"github.com/Fox216540/shop/payment-service/domain/payment"
	"github.com/google/uuid"
	pay "github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
	"github.com/shopspring/decimal"
)

type provider struct{}

func NewProvider() payment.Provider {
	return &provider{}
}

func (p *provider) CreatePaymentByOrderID(
	idOrder uuid.UUID,
	value, currency, description, returnURL string,
) (payment.Payment, error) {
	amount, err := decimal.NewFromString(value)
	if err != nil {
		return payment.Payment{}, err
	}

	return payment.Payment{
		ID:          uuid.NewString(),
		OrderID:     idOrder,
		Amount:      payment.Amount{Value: amount, Currency: currency},
		Method:      "mock",
		ReturnURL:   returnURL,
		Description: description,
		Status:      pay.Pending,
	}, nil
}
