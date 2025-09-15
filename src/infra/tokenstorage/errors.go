package tokenstorage

import (
	"shop/src/domain/tokenstorage"
)

const layer = "Infra"

type ServerError struct {
	*tokenstorage.GlobalError
}

func (e *ServerError) Error() string {
	return e.GlobalError.Error()
}

func NewServerError(msg string, err error) *ServerError {
	return &ServerError{
		GlobalError: tokenstorage.NewGlobalError(msg, err, layer),
	}
}

type InvalidSet struct {
	*ServerError
}

func (e *InvalidSet) Error() string {
	return e.ServerError.Error()
}

func NewInvalidSet(err error) error {
	return &InvalidSet{
		ServerError: NewServerError("Invalid Set Error", err),
	}
}

type InvalidExists struct {
	*ServerError
}

func (e *InvalidExists) Error() string {
	return e.ServerError.Error()
}

func NewInvalidExists(err error) error {
	return &InvalidExists{
		ServerError: NewServerError("Invalid Exists Error", err),
	}
}

type InvalidDelete struct {
	*ServerError
}

func (e *InvalidDelete) Error() string {
	return e.ServerError.Error()
}

func NewInvalidDelete(err error) error {
	return &InvalidDelete{
		ServerError: NewServerError("Invalid Delete Error", err),
	}
}

type InvalidDeleteAll struct {
	*ServerError
}

func (e *InvalidDeleteAll) Error() string {
	return e.ServerError.Error()
}

func NewInvalidDeleteAll(err error) error {
	return &InvalidDeleteAll{
		ServerError: NewServerError("Invalid Delete All Error", err),
	}
}
