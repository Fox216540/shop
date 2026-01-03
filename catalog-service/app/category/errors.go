package category

import "github.com/Fox216540/shop/catalog-service/app/globalError"

const domain = "Category"

type CategoryServerError struct {
	*globalError.AppServerError
}

func (e *CategoryServerError) Error() string {
	return e.AppServerError.Error()
}

func (e *CategoryServerError) Unwrap() error {
	return e.AppServerError
}

func NewCategoryServerError(msg string, err error) *CategoryServerError {
	return &CategoryServerError{
		AppServerError: globalError.NewAppServerError(msg, domain, err),
	}
}

type InvalidGetCategories struct {
	*CategoryServerError
}

func (e *InvalidGetCategories) Error() string {
	return e.CategoryServerError.Error()
}

func (e *InvalidGetCategories) Unwrap() error {
	return e.CategoryServerError
}

func NewInvalidGetCategories(err error) error {
	return &InvalidGetCategories{
		CategoryServerError: NewCategoryServerError("Invalid Get Categories", err),
	}
}
