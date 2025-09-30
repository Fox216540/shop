package product

import "shop/src/app/globalError"

const domain = "Product"

type ProductServerError struct {
	*globalError.AppServerError
}

func (e *ProductServerError) Error() string {
	return e.AppServerError.Error()
}

func NewCategoryServerError(msg string, err error) *ProductServerError {
	return &ProductServerError{
		AppServerError: globalError.NewAppServerError(msg, domain, err),
	}
}

type InvalidProductsOfCategoryID struct {
	*ProductServerError
}

func (e *InvalidProductsOfCategoryID) Error() string {
	return e.ProductServerError.Error()
}

func NewInvalidProductsOfCategoryID(err error) error {
	return &InvalidProductsOfCategoryID{
		ProductServerError: NewCategoryServerError("Invalid Products Of Category ID", err),
	}
}

type InvalidProductByID struct {
	*ProductServerError
}

func (e *InvalidProductByID) Error() string {
	return e.ProductServerError.Error()
}

func NewInvalidProductByID(err error) error {
	return &InvalidProductByID{
		ProductServerError: NewCategoryServerError("Invalid Product By ID", err),
	}
}

type InvalidProductsByIDs struct {
	*ProductServerError
}

func (e *InvalidProductsByIDs) Error() string {
	return e.ProductServerError.Error()
}

func NewInvalidProductsByIDs(err error) error {
	return &InvalidProductsByIDs{
		ProductServerError: NewCategoryServerError("Invalid Products By IDs", err),
	}
}
