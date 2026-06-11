package app

import (
	"github.com/Fox216540/shop/payment-service/domain/payment"
	"github.com/google/uuid"
)

type UseCase interface {
	CreatePayment(
		idOrder uuid.UUID,
		value, currency, description, returnURL string,
	) (payment.Payment, error)
}
