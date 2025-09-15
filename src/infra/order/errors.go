package order

import (
	"shop/src/infra/globalError"
)

const domain = "Order"

type OrderServerError struct {
	*globalError.InfraServerError
}

func (e *OrderServerError) Error() string {
	return e.InfraServerError.Error()
}

func NewOrderServerError(msg string, err error) *OrderServerError {
	return &OrderServerError{
		InfraServerError: globalError.NewInfraServerError(msg, domain, err),
	}
}

type InvalidSaveOrder struct {
	*OrderServerError
}

func (e *InvalidSaveOrder) Error() string {
	return e.OrderServerError.Error()
}

func NewInvalidSaveOrder(err error) error {
	return &InvalidSaveOrder{
		OrderServerError: NewOrderServerError("Invalid Save Order Error", err),
	}
}

type InvalidRemoveOrder struct {
	*OrderServerError
}

func (e *InvalidRemoveOrder) Error() string {
	return e.OrderServerError.Error()
}

func NewInvalidRemoveOrder(err error) error {
	return &InvalidRemoveOrder{
		OrderServerError: NewOrderServerError("Invalid Remove Order Error", err),
	}
}

type InvalidGetOrderByID struct {
	*OrderServerError
}

func (e *InvalidGetOrderByID) Error() string {
	return e.OrderServerError.Error()
}

func NewInvalidGetOrderByID(err error) error {
	return &InvalidGetOrderByID{
		OrderServerError: NewOrderServerError("Invalid Get Order By ID Error", err),
	}
}

type InvalidGetOrdersByUserID struct {
	*OrderServerError
}

func (e *InvalidGetOrdersByUserID) Error() string {
	return e.OrderServerError.Error()
}

func NewInvalidGetOrdersByUserID(err error) error {
	return &InvalidGetOrdersByUserID{
		OrderServerError: NewOrderServerError("Invalid Get Orders By User ID Error", err),
	}
}

type InvalidCheckOrderNum struct {
	*OrderServerError
}

func (e *InvalidCheckOrderNum) Error() string {
	return e.OrderServerError.Error()
}

func NewInvalidCheckOrderNum(err error) error {
	return &InvalidCheckOrderNum{
		OrderServerError: NewOrderServerError("Invalid Check Order Number Error", err),
	}
}
