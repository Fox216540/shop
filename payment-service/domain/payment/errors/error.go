package category

import (
	"github.com/Fox216540/shop/payment-service/core/exception"
)

const (
	domain = "Payment"
)

type DomainConflictError struct {
	*exception.ConflictError
}

func (e *DomainConflictError) Error() string {
	return e.ConflictError.Error()
}

func (e *DomainConflictError) Unwrap() error {
	return e.ConflictError
}

func NewDomainConflictError(msg string, err error) *DomainConflictError {
	return &DomainConflictError{
		ConflictError: exception.NewConflictError(msg, domain, err),
	}
}

type NoRedirectConfirmation struct {
	*DomainConflictError
}

func (e *NoRedirectConfirmation) Error() string {
	return e.ConflictError.Error()
}

func (e *NoRedirectConfirmation) Unwrap() error {
	return e.ConflictError
}

func NewNoRedirectConfirmation(err error) error {
	return &NoRedirectConfirmation{
		DomainConflictError: NewDomainConflictError("No redirect confirmation", err),
	}
}
