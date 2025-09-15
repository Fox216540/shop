package order

import (
	"shop/src/domain/order"
)

const layer = "Infra"

type ServerError struct {
	*order.GlobalError
}

func (e *ServerError) Error() string {
	return e.GlobalError.Error()
}

func NewServerError(msg string, err error) *ServerError {
	return &ServerError{
		GlobalError: order.NewGlobalError(msg, err, layer),
	}
}

type InvalidSaveOrder struct {
	*ServerError
}

func (e *InvalidSaveOrder) Error() string {
	return e.ServerError.Error()
}

func NewInvalidSaveOrder(err error) error {
	return &InvalidSaveOrder{
		ServerError: NewServerError("Invalid Save Order Error", err),
	}
}

type InvalidRemoveOrder struct {
	*ServerError
}

func (e *InvalidRemoveOrder) Error() string {
	return e.ServerError.Error()
}

func NewInvalidRemoveOrder(err error) error {
	return &InvalidRemoveOrder{
		ServerError: NewServerError("Invalid Remove Order Error", err),
	}
}

type InvalidGetOrderByID struct {
	*ServerError
}

func (e *InvalidGetOrderByID) Error() string {
	return e.ServerError.Error()
}

func NewInvalidGetOrderByID(err error) error {
	return &InvalidGetOrderByID{
		ServerError: NewServerError("Invalid Get Order By ID Error", err),
	}
}

type InvalidGetOrdersByUserID struct {
	*ServerError
}

func (e *InvalidGetOrdersByUserID) Error() string {
	return e.ServerError.Error()
}

func NewInvalidGetOrdersByUserID(err error) error {
	return &InvalidGetOrdersByUserID{
		ServerError: NewServerError("Invalid Get Orders By User ID Error", err),
	}
}

type InvalidCheckOrderNum struct {
	*ServerError
}

func (e *InvalidCheckOrderNum) Error() string {
	return e.ServerError.Error()
}

func NewInvalidCheckOrderNum(err error) error {
	return &InvalidCheckOrderNum{
		ServerError: NewServerError("Invalid Check Order Number Error", err),
	}
}
