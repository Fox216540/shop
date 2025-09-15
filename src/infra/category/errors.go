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

func NewInvalidFindAllError(err error) error {
	return &InvalidFindAllError{
		DomainServerError: NewDomainServerError("Invalid Find All Error", err),
	}
}
