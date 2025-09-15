package hasher

import (
	hasherdomain "shop/src/domain/hasher"
)

const layer = "Infra"

type ServerError struct {
	*hasherdomain.GlobalError
}

func (e *ServerError) Error() string {
	return e.GlobalError.Error()
}

func NewServerError(msg string, err error) *ServerError {
	return &ServerError{
		GlobalError: hasherdomain.NewGlobalError(msg, err, layer),
	}
}

type InvalidHashError struct {
	*ServerError
}

func (e *InvalidHashError) Error() string {
	return e.ServerError.Error()
}

func NewInvalidHashError(err error) error {
	return &InvalidHashError{
		ServerError: NewServerError("Invalid Find All Error", err),
	}
}

type InvalidVerifyError struct {
	*ServerError
}

func (e *InvalidVerifyError) Error() string {
	return e.ServerError.Error()
}

func NewInvalidVerifyError(err error) error {
	return &InvalidVerifyError{
		ServerError: NewServerError("Invalid Verify Error", err),
	}
}
