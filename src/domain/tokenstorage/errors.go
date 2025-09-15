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

func NewDomainNotFoundError(msg string, err error) *DomainNotFoundError {
	return &DomainNotFoundError{
		NotFoundError: exception.NewNotFoundError(msg, domain, err),
	}
}

type NotFoundTokensOfUserError struct {
	*DomainNotFoundError
}

func (e *NotFoundTokensOfUserError) Error() string {
	return e.NotFoundError.Error()
}

func NewNotFoundTokensOfUserError(err error) *NotFoundTokensOfUserError {
	return &NotFoundTokensOfUserError{
		DomainNotFoundError: NewDomainNotFoundError("Tokens of user not found", err),
	}
}
