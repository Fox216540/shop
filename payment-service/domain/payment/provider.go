package payment

import "github.com/google/uuid"

type Provider interface {
	CreatePaymentByOrderID(idOrder uuid.UUID, value, currency, description string) (Payment, error)
}
