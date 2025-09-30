package order

import "shop/src/app/globalError"

const domain = "Order"

type OrderServerError struct {
	*globalError.AppServerError
}

func (e *OrderServerError) Error() string {
	return e.AppServerError.Error()
}

func NewCategoryServerError(msg string, err error) *OrderServerError {
	return &OrderServerError{
		AppServerError: globalError.NewAppServerError(msg, domain, err),
	}
}

type InvalidPlace struct {
	*OrderServerError
}

func (e *InvalidPlace) Error() string {
	return e.OrderServerError.Error()
}

func NewInvalidPlace(err error) error {
	return &InvalidPlace{
		OrderServerError: NewCategoryServerError("Invalid Place", err),
	}
}

type InvalidCancel struct {
	*OrderServerError
}

func (e *InvalidCancel) Error() string {
	return e.OrderServerError.Error()
}

func NewInvalidCancel(err error) error {
	return &InvalidCancel{
		OrderServerError: NewCategoryServerError("Invalid Cancel", err),
	}
}

type InvalidGetByID struct {
	*OrderServerError
}

func (e *InvalidGetByID) Error() string {
	return e.OrderServerError.Error()
}

func NewInvalidGetByID(err error) error {
	return &InvalidGetByID{
		OrderServerError: NewCategoryServerError("Invalid Get By ID", err),
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
		OrderServerError: NewCategoryServerError("Invalid Get Orders By User ID", err),
	}
}
