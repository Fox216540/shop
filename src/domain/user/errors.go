package user

import (
	"shop/src/core/exception"
)

const (
	domain = "User"
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

type NotFoundUserError struct {
	*DomainNotFoundError
}

func (e *NotFoundUserError) Error() string {
	return e.NotFoundError.Error()
}

func NewNotFoundUserError(err error) error {
	return &NotFoundUserError{
		DomainNotFoundError: NewDomainNotFoundError("User not found", err),
	}
}
