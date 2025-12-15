package middleware

import "github.com/Fox216540/shop/apigateway-service/core/exception"

const domain = "MiddleWare"

type MiddlewareError struct {
	*exception.UnauthorizedError
}

func NewMiddlewareError(err error) *MiddlewareError {
	return &MiddlewareError{
		UnauthorizedError: exception.NewUnauthorizedError("Unauthorized", domain, err),
	}
}

func (e *MiddlewareError) Error() string {
	return e.UnauthorizedError.Error()
}

func (e *MiddlewareError) Unwrap() error {
	return e.UnauthorizedError
}
