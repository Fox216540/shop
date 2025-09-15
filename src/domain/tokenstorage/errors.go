package tokenstorage

import (
	"fmt"
	"shop/src/core/exception"
)

const (
	domain = "Token Storage"
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
		GlobalError: NewGlobalError(msg, err, domain),
	}
}

type NotFoundError struct {
	*DomainError
}

func NewNotFoundError(msg string, err error) *NotFoundError {
	return &NotFoundError{
		DomainError: NewDomainError(msg, err),
	}
}

type NotFoundTokensOfUserError struct {
	*NotFoundError
}

func (e *NotFoundTokensOfUserError) Error() string {
	return e.NotFoundError.Error()
}

func NewNotFoundTokensOfUserError(err error) *NotFoundTokensOfUserError {
	return &NotFoundTokensOfUserError{
		NotFoundError: NewNotFoundError("Tokens of user not found", err),
	}
}
