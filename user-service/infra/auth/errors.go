package auth

import "github.com/Fox216540/shop/user-service/infra/globalError"

const domain = "Auth"

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
