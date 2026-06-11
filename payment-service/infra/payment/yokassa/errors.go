package yokassa

import "github.com/Fox216540/shop/payment-service/infra/globalError"

const layer = "Payment | Provider | YooKassa"

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

type InvalidCreatePayment struct {
	*PaymentServerError
}

func (e *InvalidCreatePayment) Error() string {
	return e.PaymentServerError.Error()
}

func (e *InvalidCreatePayment) Unwrap() error {
	return e.PaymentServerError
}

func NewInvalidCreatePayment(err error) error {
	return &InvalidCreatePayment{
		PaymentServerError: NewPaymentServerError("Invalid Create Payment", err),
	}
}

type InvalidParsePaymentAmount struct {
	*PaymentServerError
}

func (e *InvalidParsePaymentAmount) Error() string {
	return e.PaymentServerError.Error()
}

func (e *InvalidParsePaymentAmount) Unwrap() error {
	return e.PaymentServerError
}

func NewInvalidParsePaymentAmount(err error) error {
	return &InvalidParsePaymentAmount{
		PaymentServerError: NewPaymentServerError("Invalid Parse Payment Amount", err),
	}
}
