package product

import (
	"shop/src/infra/globalError"
)

const domain = "Product"

type ProductServerError struct {
	*globalError.InfraServerError
}

func (e *ProductServerError) Error() string {
	return e.InfraServerError.Error()
}

func NewProductServerError(msg string, err error) *ProductServerError {
	return &ProductServerError{
		InfraServerError: globalError.NewInfraServerError(msg, domain, err),
	}
}

type InvalidFindProductsByCategoryID struct {
	*ProductServerError
}

func (e *InvalidFindProductsByCategoryID) Error() string {
	return e.ProductServerError.Error()
}

func NewInvalidFindProductsByCategoryID(err error) error {
	return &InvalidFindProductsByCategoryID{
		ProductServerError: NewProductServerError("Invalid Find Products By Category ID Error", err),
	}
}

type InvalidFindProductByID struct {
	*ProductServerError
}

func (e *InvalidFindProductByID) Error() string {
	return e.ProductServerError.Error()
}

func NewInvalidFindProductByID(err error) error {
	return &InvalidFindProductByID{
		ProductServerError: NewProductServerError("Invalid Find Product By ID Error", err),
	}
}

type InvalidFindProductsByIDs struct {
	*ProductServerError
}

func (e *InvalidFindProductsByIDs) Error() string {
	return e.ProductServerError.Error()
}

func NewInvalidFindProductsByIDs(err error) error {
	return &InvalidFindProductsByIDs{
		ProductServerError: NewProductServerError("Invalid Find Products By IDs Error", err),
	}
}
