package hasher

import "shop/src/core/exception"

const domain = "Hasher"

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

type BadPasswordError struct {
	*DomainBadRequestError
}

func (e *BadPasswordError) Error() string {
	return e.BadRequestError.Error()
}

func (e *BadPasswordError) Unwrap() error {
	return e.BadRequestError
}

func NewBadPasswordError(err error) *BadPasswordError {
	return &BadPasswordError{
		DomainBadRequestError: NewDomainBadRequestError("Bad password", err),
	}
}
