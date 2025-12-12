package globalError

import "github.com/Fox216540/shop/order-service/core/exception"

const (
	layer = "Infra"
)

type InfraServerError struct {
	*exception.ServerError
}

func (e *InfraServerError) Error() string {
	return e.ServerError.Error()
}

func (e *InfraServerError) Unwrap() error {
	return e.ServerError
}

func NewInfraServerError(msg, domain string, err error) *InfraServerError {
	return &InfraServerError{
		ServerError: exception.NewServerError(msg, domain, layer, err),
	}
}

type GRPCErrors struct {
	*InfraServerError
}

func (e *GRPCErrors) Error() string {
	return e.InfraServerError.Error()
}

func (e *GRPCErrors) Unwrap() error {
	return e.InfraServerError
}

func NewGRPCErrors(domain string, err error) *GRPCErrors {
	return &GRPCErrors{
		InfraServerError: NewInfraServerError("Problems with GRPC", domain, err),
	}
}
