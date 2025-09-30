package product

import "shop/src/app/globalError"

const domain = "Product"

type CategoryServerError struct {
	*globalError.AppServerError
}

func (e *CategoryServerError) Error() string {
	return e.AppServerError.Error()
}

func NewCategoryServerError(msg string, err error) *CategoryServerError {
	return &CategoryServerError{
		AppServerError: globalError.NewAppServerError(msg, domain, err),
	}
}

type InvalidProductsOfCategoryID struct {
	*CategoryServerError
}

func (e *InvalidProductsOfCategoryID) Error() string {
	return e.CategoryServerError.Error()
}

func NewInvalidProductsOfCategoryID(err error) error {
	return &InvalidProductsOfCategoryID{
		CategoryServerError: NewCategoryServerError("Invalid Products Of Category ID", err),
	}
}

type InvalidProductByID struct {
	*CategoryServerError
}

func (e *InvalidProductByID) Error() string {
	return e.CategoryServerError.Error()
}

func NewInvalidProductByID(err error) error {
	return &InvalidProductByID{
		CategoryServerError: NewCategoryServerError("Invalid Product By ID", err),
	}
}

type InvalidProductsByIDs struct {
	*CategoryServerError
}

func (e *InvalidProductsByIDs) Error() string {
	return e.CategoryServerError.Error()
}

func NewInvalidProductsByIDs(err error) error {
	return &InvalidProductsByIDs{
		CategoryServerError: NewCategoryServerError("Invalid Products By IDs", err),
	}
}
