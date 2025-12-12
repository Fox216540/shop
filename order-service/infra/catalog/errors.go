package catalog

import "github.com/Fox216540/shop/order-service/infra/globalError"

const domain = "Catalog"

type GRPCError struct {
	*globalError.GRPCErrors
}

func (e *GRPCError) Error() string {
	return e.GRPCErrors.Error()
}

func (e *GRPCError) Unwrap() error {
	return e.GRPCErrors
}

func NewGRPCError(err error) *GRPCError {
	return &GRPCError{
		GRPCErrors: globalError.NewGRPCErrors(domain, err),
	}
}

type InvalidUUID struct {
	*globalError.InfraServerError
}

func (e *InvalidUUID) Error() string {
	return e.InfraServerError.Error()
}

func (e *InvalidUUID) Unwrap() error {
	return e.InfraServerError
}

func NewInvalidUUID(err error) *InvalidUUID {
	return &InvalidUUID{
		InfraServerError: globalError.NewInfraServerError("Invalid UUID", domain, err),
	}
}
