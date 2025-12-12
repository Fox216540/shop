package order

import (
	"github.com/Fox216540/shop/order-service/core/exception"
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

type OrderNumberAlreadyExistsError struct {
	*exception.BadRequestError
}

func (e *OrderNumberAlreadyExistsError) Error() string {
	return e.BadRequestError.Error()
}

func (e *OrderNumberAlreadyExistsError) Unwrap() error {
	return e.BadRequestError
}

func NewOrderNumberAlreadyExistsError(err error) error {
	return &OrderNumberAlreadyExistsError{
		BadRequestError: exception.NewBadRequestError("Order number already exists", domain, err),
	}
}

type ProductOfOrderNotFoundError struct {
	*DomainNotFoundError
}

func (e *ProductOfOrderNotFoundError) Error() string {
	return e.NotFoundError.Error()
}

func (e *ProductOfOrderNotFoundError) Unwrap() error {
	return e.NotFoundError
}

func NewProductOfOrderNotFoundError(err error) error {
	return &ProductOfOrderNotFoundError{
		DomainNotFoundError: NewDomainNotFoundError("Product of order not found", err),
	}
}
