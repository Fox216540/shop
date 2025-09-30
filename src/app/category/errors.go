package category

import "shop/src/app/globalError"

const domain = "Category"

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

type InvalidGetCategories struct {
	*CategoryServerError
}

func (e *InvalidGetCategories) Error() string {
	return e.CategoryServerError.Error()
}

func NewInvalidGetCategories(err error) error {
	return &InvalidGetCategories{
		CategoryServerError: NewCategoryServerError("Invalid Get Categories", err),
	}
}
