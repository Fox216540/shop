package tokenstorage

import "github.com/Fox216540/shop/auth-service/app/globalError"

const domain = "TokenStorage"

type TokenStorageError struct {
	*globalError.AppServerError
}

func NewTokenStorageError(msg string, err error) *TokenStorageError {
	return &TokenStorageError{
		AppServerError: globalError.NewAppServerError(msg, domain, err),
	}
}

func (e *TokenStorageError) Error() string {
	return e.AppServerError.Error()
}

func (e *TokenStorageError) Unwrap() error {
	return e.AppServerError
}

type InvalidAdd struct {
	*TokenStorageError
}

func NewInvalidAdd(err error) *InvalidAdd {
	return &InvalidAdd{
		TokenStorageError: NewTokenStorageError("Invalid add", err),
	}
}

func (e *InvalidAdd) Error() string {
	return e.TokenStorageError.Error()
}

func (e *InvalidAdd) Unwrap() error {
	return e.TokenStorageError
}

type InvalidDelete struct {
	*TokenStorageError
}

func NewInvalidDelete(err error) *InvalidDelete {
	return &InvalidDelete{
		TokenStorageError: NewTokenStorageError("Invalid delete", err),
	}
}

func (e *InvalidDelete) Error() string {
	return e.TokenStorageError.Error()
}

func (e *InvalidDelete) Unwrap() error {
	return e.TokenStorageError
}

type InvalidDeleteAll struct {
	*TokenStorageError
}

func NewInvalidDeleteAll(err error) *InvalidDeleteAll {
	return &InvalidDeleteAll{
		TokenStorageError: NewTokenStorageError("Invalid delete all", err),
	}
}

func (e *InvalidDeleteAll) Error() string {
	return e.TokenStorageError.Error()
}

func (e *InvalidDeleteAll) Unwrap() error {
	return e.TokenStorageError
}
