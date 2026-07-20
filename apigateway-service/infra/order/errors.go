package order

import "github.com/Fox216540/shop/apigateway-service/infra/globalError"

const domain = "OrderGRPCClient"

type GRPCResponseError struct {
	*globalError.InfraServerError
}

func (e *GRPCResponseError) Error() string {
	return e.InfraServerError.Error()
}

func (e *GRPCResponseError) Unwrap() error {
	return e.InfraServerError
}

func NewGRPCResponseError(msg string, err error) *GRPCResponseError {
	return &GRPCResponseError{
		InfraServerError: globalError.NewInfraServerError(msg, domain, err),
	}
}

func NewOrderResponseNilError(err error) *GRPCResponseError {
	return NewGRPCResponseError("Order response is nil", err)
}

func NewOrderPayloadNilError(err error) *GRPCResponseError {
	return NewGRPCResponseError("Order payload is nil", err)
}

func NewOrderItemNilError(err error) *GRPCResponseError {
	return NewGRPCResponseError("Order item is nil", err)
}

func NewOrdersResponseNilError(err error) *GRPCResponseError {
	return NewGRPCResponseError("Orders response is nil", err)
}

func NewDeleteOrderResponseNilError(err error) *GRPCResponseError {
	return NewGRPCResponseError("Delete order response is nil", err)
}
