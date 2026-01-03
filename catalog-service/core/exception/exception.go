package exception

import (
	"fmt"
)

type Exception struct {
	Domain  string
	Layer   string
	Message string
	Err     error
}

func (e Exception) Error() string {
	if e.Err != nil {
		return fmt.Sprintf(
			"Domain: %s\nLayer: %s\nMessage: %s\nError: %s",
			e.Domain, e.Layer, e.Message, e.Err,
		)
	}

	return fmt.Sprintf(
		"Domain: %s\nLayer: %s\nMessage: %s",
		e.Domain, e.Layer, e.Message,
	)
}

func (e Exception) Unwrap() error {
	return e.Err
}

func NewException(msg, domain, layer string, err error) *Exception {
	return &Exception{
		Message: msg,
		Err:     err,
		Layer:   layer,
		Domain:  domain,
	}
}

const (
	domainLayer = "Domain"
)

type DomainError struct {
	*Exception
}

func (e DomainError) Error() string {
	return e.Exception.Error()
}

func (e DomainError) Unwrap() error { return e.Exception }

func NewDomainException(msg, domain string, err error) *DomainError {
	return &DomainError{
		Exception: NewException(msg, domain, domainLayer, err),
	}
}

type NotFoundError struct {
	*DomainError
}

func (e *NotFoundError) Error() string {
	return e.DomainError.Error()
}

func (e *NotFoundError) Unwrap() error { return e.DomainError }

func NewNotFoundError(msg, domain string, err error) *NotFoundError {
	return &NotFoundError{
		DomainError: NewDomainException(msg, domain, err),
	}
}

type ServerError struct {
	*Exception
}

func (e *ServerError) Error() string {
	return e.Exception.Error()
}

func (e *ServerError) Unwrap() error { return e.Exception }

func NewServerError(msg, domain, layer string, err error) *ServerError {
	return &ServerError{
		Exception: &Exception{
			Message: msg,
			Domain:  domain,
			Layer:   layer,
			Err:     err,
		},
	}
}
