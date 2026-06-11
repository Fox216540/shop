package app

import "github.com/Fox216540/shop/catalog-service/app/globalError"

const domain = "Catalog"

type CatalogServerError struct {
	*globalError.AppServerError
}

func (e *CatalogServerError) Error() string {
	return e.AppServerError.Error()
}

func (e *CatalogServerError) Unwrap() error {
	return e.AppServerError
}

func NewCatalogServerError(msg string, err error) *CatalogServerError {
	return &CatalogServerError{
		AppServerError: globalError.NewAppServerError(msg, domain, err),
	}
}

type InvalidGetCategories struct {
	*CatalogServerError
}

func (e *InvalidGetCategories) Error() string {
	return e.CatalogServerError.Error()
}

func (e *InvalidGetCategories) Unwrap() error {
	return e.CatalogServerError
}

func NewInvalidGetCategories(err error) error {
	return &InvalidGetCategories{
		CatalogServerError: NewCatalogServerError("Invalid Get Categories", err),
	}
}

type InvalidGetAll struct {
	*CatalogServerError
}

func (e *InvalidGetAll) Error() string {
	return e.CatalogServerError.Error()
}

func (e *InvalidGetAll) Unwrap() error {
	return e.CatalogServerError
}

func NewInvalidGetAll(err error) error {
	return &InvalidGetAll{
		CatalogServerError: NewCatalogServerError("Invalid Get All", err),
	}
}

type InvalidGetProductsOfCategoryID struct {
	*CatalogServerError
}

func (e *InvalidGetProductsOfCategoryID) Error() string {
	return e.CatalogServerError.Error()
}

func (e *InvalidGetProductsOfCategoryID) Unwrap() error {
	return e.CatalogServerError
}

func NewInvalidGetProductsOfCategoryID(err error) error {
	return &InvalidGetProductsOfCategoryID{
		CatalogServerError: NewCatalogServerError("Invalid Get Products Of Category ID", err),
	}
}

type InvalidGetProductByID struct {
	*CatalogServerError
}

func (e *InvalidGetProductByID) Error() string {
	return e.CatalogServerError.Error()
}

func (e *InvalidGetProductByID) Unwrap() error {
	return e.CatalogServerError
}

func NewInvalidGetProductByID(err error) error {
	return &InvalidGetProductByID{
		CatalogServerError: NewCatalogServerError("Invalid Get Product By ID", err),
	}
}

type InvalidGetProductsByIDs struct {
	*CatalogServerError
}

func (e *InvalidGetProductsByIDs) Error() string {
	return e.CatalogServerError.Error()
}

func (e *InvalidGetProductsByIDs) Unwrap() error {
	return e.CatalogServerError
}

func NewInvalidGetProductsByIDs(err error) error {
	return &InvalidGetProductsByIDs{
		CatalogServerError: NewCatalogServerError("Invalid Get Products By IDs", err),
	}
}
