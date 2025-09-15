package idgenerator

import (
	"shop/src/domain/idgenerator"
)

const layer = "Infra"

type ServerError struct {
	*idgenerator.GlobalError
}

func (e *ServerError) Error() string {
	return e.GlobalError.Error()
}

func NewServerError(msg string, err error) *ServerError {
	return &ServerError{
		GlobalError: idgenerator.NewGlobalError(msg, err, layer),
	}
}

type InvalidGenerateError struct {
	*ServerError
}

func (e *InvalidGenerateError) Error() string {
	return e.ServerError.Error()
}

func NewInvalidGenerateError(err error) error {
	return &InvalidGenerateError{
		ServerError: NewServerError("Invalid Generate New ID Error", err),
	}
}
