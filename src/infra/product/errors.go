package product

import (
	"shop/src/domain/product"
)

const layer = "Infra"

type ServerError struct {
	*product.GlobalError
}

func (e *ServerError) Error() string {
	return e.GlobalError.Error()
}

func NewServerError(msg string, err error) *ServerError {
	return &ServerError{
		GlobalError: product.NewGlobalError(msg, err, layer),
	}
}

type InvalidFindProductsByCategoryID struct {
	*ServerError
}

func (e *InvalidFindProductsByCategoryID) Error() string {
	return e.ServerError.Error()
}

func NewInvalidFindProductsByCategoryID(err error) error {
	return &InvalidFindProductsByCategoryID{
		ServerError: NewServerError("Invalid Find Products By Category ID Error", err),
	}
}

type InvalidFindProductByID struct {
	*ServerError
}

func (e *InvalidFindProductByID) Error() string {
	return e.ServerError.Error()
}

func NewInvalidFindProductByID(err error) error {
	return &InvalidFindProductByID{
		ServerError: NewServerError("Invalid Find Product By ID Error", err),
	}
}

type InvalidFindProductsByIDs struct {
	*ServerError
}

func (e *InvalidFindProductsByIDs) Error() string {
	return e.ServerError.Error()
}

func NewInvalidFindProductsByIDs(err error) error {
	return &InvalidFindProductsByIDs{
		ServerError: NewServerError("Invalid Find Products By IDs Error", err),
	}
}
