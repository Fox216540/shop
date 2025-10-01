package tokenstorage

import (
	"shop/src/core/exception"
)

const (
	domain = "Token Storage"
)

type DomainNotFoundError struct {
	*exception.NotFoundError
}

func (e *DomainNotFoundError) Error() string {
	return e.NotFoundError.Error()
}

func (e *DomainNotFoundError) Unwrap() error {
	return e.NotFoundError
}

func NewDomainNotFoundError(msg string, err error) *DomainNotFoundError {
	return &DomainNotFoundError{
		NotFoundError: exception.NewNotFoundError(msg, domain, err),
	}
}

type NotFoundTokenOfUserError struct {
	*DomainNotFoundError
}

func (e *NotFoundTokenOfUserError) Error() string {
	return e.NotFoundError.Error()
}

func (e *NotFoundTokenOfUserError) Unwrap() error {
	return e.NotFoundError
}

func NewNotFoundTokenOfUserError(err error) *NotFoundTokenOfUserError {
	return &NotFoundTokenOfUserError{
		DomainNotFoundError: NewDomainNotFoundError("Token of user not found", err),
	}
}

type NotFoundTokensOfUserError struct {
	*DomainNotFoundError
}

func (e *NotFoundTokensOfUserError) Error() string {
	return e.NotFoundError.Error()
}

func (e *NotFoundTokensOfUserError) Unwrap() error {
	return e.NotFoundError
}

func NewNotFoundTokensOfUserError(err error) *NotFoundTokensOfUserError {
	return &NotFoundTokensOfUserError{
		DomainNotFoundError: NewDomainNotFoundError("Tokens of user not found", err),
	}
}
