package user

import (
	"fmt"
	"shop/src/core/exception"
)

const (
	layer  = "Domain"
	domain = "User"
)

type GlobalError struct {
	*exception.Exception
	Domain  string
	Layer   string
	Message string
	Err     error
}

func (e *GlobalError) Error() string {
	return fmt.Sprintf("Domain: %s\nLayer: %s\nMessage: %s\nCause: %v",
		e.Domain, e.Layer, e.Message, e.Err)
}

func (e *GlobalError) Unwrap() error {
	return e.Err
}

func NewGlobalError(msg string, err error, layer string) *GlobalError {
	return &GlobalError{
		Exception: &exception.Exception{},
		Domain:    domain,
		Layer:     layer,
		Message:   msg,
		Err:       err,
	}
}

type DomainError struct {
	*GlobalError
}

func (e *DomainError) Error() string {
	return e.GlobalError.Error()
}

func NewDomainError(msg string, err error) *DomainError {
	return &DomainError{
		GlobalError: NewGlobalError(msg, err, layer),
	}
}

type NotFoundError struct {
	*DomainError
}

func (e *NotFoundError) Error() string {
	return e.DomainError.Error()
}

func NewNotFoundError(msg string, err error) *NotFoundError {
	return &NotFoundError{
		DomainError: NewDomainError(msg, err),
	}
}

type NotFoundUserError struct {
	*NotFoundError
}

func (e *NotFoundUserError) Error() string {
	return e.NotFoundError.Error()
}

func NewNotFoundUserError(err error) error {
	return &NotFoundUserError{
		NotFoundError: NewNotFoundError("User not found", err),
	}
}
