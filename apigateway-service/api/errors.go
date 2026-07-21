package api

import "github.com/Fox216540/shop/apigateway-service/core/exception"

const domain = "API"

type MissingUserIDError struct {
	*exception.UnauthorizedError
}

func (e *MissingUserIDError) Error() string {
	return e.UnauthorizedError.Error()
}

func (e *MissingUserIDError) Unwrap() error {
	return e.UnauthorizedError
}

func NewMissingUserIDError(err error) *MissingUserIDError {
	return &MissingUserIDError{
		UnauthorizedError: exception.NewUnauthorizedError(messages.MissingUserID, domain, err),
	}
}
