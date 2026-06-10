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

type InvalidOrderItemPrice struct {
	*exception.BadRequestError
}

func (e *InvalidOrderItemPrice) Error() string {
	return e.BadRequestError.Error()
}

func (e *InvalidOrderItemPrice) Unwrap() error {
	return e.BadRequestError
}

func NewInvalidOrderItemPrice(err error) error {
	return &InvalidOrderItemPrice{
		BadRequestError: exception.NewBadRequestError("Invalid order item price", domain, err),
	}
}

type InvalidOrderItemPriceValue struct {
	*exception.BadRequestError
}

func (e *InvalidOrderItemPriceValue) Error() string {
	return e.BadRequestError.Error()
}

func (e *InvalidOrderItemPriceValue) Unwrap() error {
	return e.BadRequestError
}

func NewInvalidOrderItemPriceValue(err error) error {
	return &InvalidOrderItemPriceValue{
		BadRequestError: exception.NewBadRequestError("Invalid order item price value", domain, err),
	}
}

type ProductPriceMismatch struct {
	*exception.BadRequestError
}

func (e *ProductPriceMismatch) Error() string {
	return e.BadRequestError.Error()
}

func (e *ProductPriceMismatch) Unwrap() error {
	return e.BadRequestError
}

func NewProductPriceMismatch(err error) error {
	return &ProductPriceMismatch{
		BadRequestError: exception.NewBadRequestError("Product price mismatch", domain, err),
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
