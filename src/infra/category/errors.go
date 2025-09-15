package category

import (
	"shop/src/domain/category"
)

const layer = "Infra"

type ServerError struct {
	*category.GlobalError
}

func (e *ServerError) Error() string {
	return e.GlobalError.Error()
}

func NewServerError(msg string, err error) *ServerError {
	return &ServerError{
		GlobalError: category.NewGlobalError(msg, err, layer),
	}
}

type InvalidFindAllError struct {
	*ServerError
}

func (e *InvalidFindAllError) Error() string {
	return e.ServerError.Error()
}

func NewInvalidFindAllError(err error) error {
	return &InvalidFindAllError{
		ServerError: NewServerError("Invalid Find All Error", err),
	}
}
