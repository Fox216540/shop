package category

import (
	"shop/src/infra/globalError"
)

const domain = "Category"

type DomainServerError struct {
	*globalError.InfraServerError
}

func (e *DomainServerError) Error() string {
	return e.InfraServerError.Error()
}

func (e *DomainServerError) Unwrap() error {
	return e.InfraServerError
}

func NewDomainServerError(msg string, err error) *DomainServerError {
	return &DomainServerError{
		InfraServerError: globalError.NewInfraServerError(msg, domain, err),
	}
}

type InvalidFindAllError struct {
	*DomainServerError
}

func (e *InvalidFindAllError) Error() string {
	return e.InfraServerError.Error()
}

func (e *InvalidFindAllError) Unwrap() error {
	return e.DomainServerError
}

func NewInvalidFindAllError(err error) error {
	return &InvalidFindAllError{
		DomainServerError: NewDomainServerError("Invalid Find All Error", err),
	}
}
