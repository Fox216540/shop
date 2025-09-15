package tokenstorage

import (
	"shop/src/infra/globalError"
)

const domain = "Token Storage"

type TokenStorageServerError struct {
	*globalError.InfraServerError
}

func (e *TokenStorageServerError) Error() string {
	return e.InfraServerError.Error()
}

func NewTokenStorageServerError(msg string, err error) *TokenStorageServerError {
	return &TokenStorageServerError{
		InfraServerError: globalError.NewInfraServerError(msg, domain, err),
	}
}

type InvalidSet struct {
	*TokenStorageServerError
}

func (e *InvalidSet) Error() string {
	return e.TokenStorageServerError.Error()
}

func NewInvalidSet(err error) error {
	return &InvalidSet{
		TokenStorageServerError: NewTokenStorageServerError("Invalid Set Error", err),
	}
}

type InvalidExists struct {
	*TokenStorageServerError
}

func (e *InvalidExists) Error() string {
	return e.TokenStorageServerError.Error()
}

func NewInvalidExists(err error) error {
	return &InvalidExists{
		TokenStorageServerError: NewTokenStorageServerError("Invalid Exists Error", err),
	}
}

type InvalidDelete struct {
	*TokenStorageServerError
}

func (e *InvalidDelete) Error() string {
	return e.TokenStorageServerError.Error()
}

func NewInvalidDelete(err error) error {
	return &InvalidDelete{
		TokenStorageServerError: NewTokenStorageServerError("Invalid Delete Error", err),
	}
}

type InvalidDeleteAll struct {
	*TokenStorageServerError
}

func (e *InvalidDeleteAll) Error() string {
	return e.TokenStorageServerError.Error()
}

func NewInvalidDeleteAll(err error) error {
	return &InvalidDeleteAll{
		TokenStorageServerError: NewTokenStorageServerError("Invalid Delete All Error", err),
	}
}
