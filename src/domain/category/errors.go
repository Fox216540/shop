package category

import (
	"shop/src/core/exception"
)

const (
	domain = "Category"
)

type DomainNotFoundError struct {
	*exception.NotFoundError
}

func (e *DomainNotFoundError) Error() string {
	return e.NotFoundError.Error()
}

func (e *DomainNotFoundError) Unwrap() error {
	return e.NotFoundError.Unwrap()
}

func NewDomainNotFoundError(msg string, err error) *DomainNotFoundError {
	return &DomainNotFoundError{
		NotFoundError: exception.NewNotFoundError(msg, domain, err),
	}
}

type NotFoundCategoryError struct {
	*DomainNotFoundError
}

func (e *NotFoundCategoryError) Error() string {
	return e.NotFoundError.Error()
}

func (e *NotFoundCategoryError) Unwrap() error {
	return e.NotFoundError
}

func NewNotFoundCategoryError(err error) error {
	return &NotFoundCategoryError{
		DomainNotFoundError: NewDomainNotFoundError("Category not found", err),
	}
}
