package user

import (
	"shop/src/core/exception"
)

const (
	domain = "User"
)

type DomainNotFoundError struct {
	*exception.NotFoundError
}

func (e *DomainNotFoundError) Error() string {
	return e.NotFoundError.Error()
}

func (e *DomainNotFoundError) Unwrap() error {
	return e.NotFoundError
}

func NewDomainNotFoundError(msg string, err error) *DomainNotFoundError {
	return &DomainNotFoundError{
		NotFoundError: exception.NewNotFoundError(msg, domain, err),
	}
}

type NotFoundUserError struct {
	*DomainNotFoundError
}

func (e *NotFoundUserError) Error() string {
	return e.NotFoundError.Error()
}

func (e *NotFoundUserError) Unwrap() error {
	return e.NotFoundError
}

func NewNotFoundUserError(err error) error {
	return &NotFoundUserError{
		DomainNotFoundError: NewDomainNotFoundError("User not found", err),
	}
}

type DomainBadRequestError struct {
	*exception.BadRequestError
}

func (e *DomainBadRequestError) Error() string {
	return e.BadRequestError.Error()
}

func (e *DomainBadRequestError) Unwrap() error {
	return e.BadRequestError
}

func NewDomainBadRequestError(msg string, err error) *DomainBadRequestError {
	return &DomainBadRequestError{
		BadRequestError: exception.NewBadRequestError(msg, domain, err),
	}
}

type ExistingEmailError struct {
	*DomainBadRequestError
}

func (e *ExistingEmailError) Error() string {
	return e.DomainBadRequestError.Error()
}

func (e *ExistingEmailError) Unwrap() error {
	return e.DomainBadRequestError
}

func NewExistingEmailError(err error) error {
	return &ExistingEmailError{
		DomainBadRequestError: NewDomainBadRequestError("Existing email", err),
	}
}

type ExistingPhoneError struct {
	*DomainBadRequestError
}

func (e *ExistingPhoneError) Error() string {
	return e.DomainBadRequestError.Error()
}

func (e *ExistingPhoneError) Unwrap() error {
	return e.DomainBadRequestError
}

func NewExistingPhoneError(err error) error {
	return &ExistingPhoneError{
		DomainBadRequestError: NewDomainBadRequestError("Existing phone", err),
	}
}

type ExistingPasswordError struct {
	*DomainBadRequestError
}

func (e *ExistingPasswordError) Error() string {
	return e.DomainBadRequestError.Error()
}

func (e *ExistingPasswordError) Unwrap() error {
	return e.DomainBadRequestError
}

func NewExistingPasswordError(err error) error {
	return &ExistingPasswordError{
		DomainBadRequestError: NewDomainBadRequestError("Existing password", err),
	}
}
