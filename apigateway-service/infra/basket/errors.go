package basket

import "github.com/Fox216540/shop/apigateway-service/infra/globalError"

const domain = "BasketGRPCClient"

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

func NewBasketItemNilError(err error) *GRPCResponseError {
	return NewGRPCResponseError("Basket item is nil", err)
}

func NewBasketResponseNilError(err error) *GRPCResponseError {
	return NewGRPCResponseError("Basket response is nil", err)
}

func NewRemoveBasketItemResponseNilError(err error) *GRPCResponseError {
	return NewGRPCResponseError("Remove basket item response is nil", err)
}
