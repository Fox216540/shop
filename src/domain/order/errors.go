package order

import (
	"shop/src/core/exception"
)

const (
	domain = "Order"
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

type NotFoundOrderError struct {
	*DomainNotFoundError
}

func (e *NotFoundOrderError) Error() string {
	return e.NotFoundError.Error()
}

func (e *NotFoundOrderError) Unwrap() error {
	return e.NotFoundError
}

func NewNotFoundOrderError(err error) error {
	return &NotFoundOrderError{
		DomainNotFoundError: NewDomainNotFoundError("Order not found", err),
	}
}
