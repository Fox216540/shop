package product

import (
	"fmt"
	"shop/src/core/exception"
)

const (
	layer  = "Domain"
	domain = "Product"
)

type GlobalError struct {
	*exception.Exception
	Domain  string
	Layer   string
	Message string
	Err     error
}

func (e *GlobalError) Error() string {
	return fmt.Sprintf("Domain: %s\nLayer: %s\nMessage: %s\nCause: %v",
		e.Domain, e.Layer, e.Message, e.Err)
}

func (e *GlobalError) Unwrap() error {
	return e.Err
}

func NewGlobalError(msg string, err error, layer string) *GlobalError {
	return &GlobalError{
		Exception: &exception.Exception{},
		Domain:    domain,
		Layer:     layer,
		Message:   msg,
		Err:       err,
	}
}

type DomainError struct {
	*GlobalError
}

func (e *DomainError) Error() string {
	return e.GlobalError.Error()
}

func NewDomainError(msg string, err error) *DomainError {
	return &DomainError{
		GlobalError: NewGlobalError(msg, err, layer),
	}
}

type NotFoundError struct {
	*DomainError
}

func (e *NotFoundError) Error() string {
	return e.DomainError.Error()
}

func NewNotFoundError(msg string, err error) *NotFoundError {
	return &NotFoundError{
		DomainError: NewDomainError(msg, err),
	}
}

type NotFoundProductError struct {
	*NotFoundError
}

func (e *NotFoundProductError) Error() string {
	return e.NotFoundError.Error()
}

func NewNotFoundProductError(err error) error {
	return &NotFoundProductError{
		NotFoundError: NewNotFoundError("Product not found", err),
	}
}

type NotFoundProductsError struct {
	*NotFoundError
}

func (e *NotFoundProductsError) Error() string {
	return e.NotFoundError.Error()
}

func NewNotFoundProductsError(err error) error {
	return &NotFoundProductsError{
		NotFoundError: NewNotFoundError("Products not found", err),
	}
}
