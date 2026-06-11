package payment

import "github.com/google/uuid"

type Provider interface {
	CreatePaymentByOrderID(
		idOrder uuid.UUID,
		value, currency, description, returnURL string,
	) (Payment, error)
}
