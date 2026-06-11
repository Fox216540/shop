package repository

import (
	"github.com/Fox216540/shop/payment-service/infra/globalError"
)

const layer = "Payment | Repository"

type PaymentServerError struct {
	*globalError.InfraServerError
}

func (e *PaymentServerError) Error() string {
	return e.InfraServerError.Error()
}

func (e *PaymentServerError) Unwrap() error {
	return e.InfraServerError
}

func NewPaymentServerError(msg string, err error) *PaymentServerError {
	return &PaymentServerError{
		InfraServerError: globalError.NewInfraServerError(msg, layer, err),
	}
}

type InvalidSavePayment struct {
	*PaymentServerError
}

func (e *InvalidSavePayment) Error() string {
	return e.PaymentServerError.Error()
}

func (e *InvalidSavePayment) Unwrap() error {
	return e.PaymentServerError
}

func NewInvalidSavePayment(err error) error {
	return &InvalidSavePayment{
		PaymentServerError: NewPaymentServerError("Invalid Save Payment", err),
	}
}
