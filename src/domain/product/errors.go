package product

import (
	"shop/src/core/exception"
)

const (
	domain = "Product"
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

type NotFoundProductError struct {
	*DomainNotFoundError
}

func (e *NotFoundProductError) Error() string {
	return e.NotFoundError.Error()
}

func NewNotFoundProductError(err error) error {
	return &NotFoundProductError{
		DomainNotFoundError: NewDomainNotFoundError("Product not found", err),
	}
}

type NotFoundProductsError struct {
	*DomainNotFoundError
}

func (e *NotFoundProductsError) Error() string {
	return e.NotFoundError.Error()
}

func NewNotFoundProductsError(err error) error {
	return &NotFoundProductsError{
		DomainNotFoundError: NewDomainNotFoundError("Products not found", err),
	}
}
