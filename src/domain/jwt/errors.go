package jwt

import (
	"fmt"
	"shop/src/core/exception"
)

const (
	layer  = "Domain"
	domain = "JWT"
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

type BadRequestError struct {
	*DomainError
}

func (e *BadRequestError) Error() string {
	return e.DomainError.Error()
}

func NewBadRequestError(msg string, err error) *BadRequestError {
	return &BadRequestError{
		DomainError: NewDomainError(msg, err),
	}
}

type BadRefreshTokenError struct {
	*BadRequestError
}

func (e *BadRefreshTokenError) Error() string {
	return e.BadRequestError.Error()
}

func NewBadRefreshTokenError(err error) error {
	return &BadRefreshTokenError{
		BadRequestError: NewBadRequestError("Bad refresh token", err),
	}
}

type BadAccessTokenError struct {
	*BadRequestError
}

func (e *BadAccessTokenError) Error() string {
	return e.BadRequestError.Error()
}

func NewBadAccessTokenError(err error) error {
	return &BadAccessTokenError{
		BadRequestError: NewBadRequestError("Bad access token", err),
	}
}

type NoValidError struct {
	*DomainError
}

func (e *NoValidError) Error() string {
	return e.DomainError.Error()
}

func NewNoValidError(msg string, err error) *NoValidError {
	return &NoValidError{
		DomainError: NewDomainError(msg, err),
	}
}

type NoValidRefreshTokenError struct {
	*NoValidError
}

func (e *NoValidRefreshTokenError) Error() string {
	return e.NoValidError.Error()
}

func NewNoValidRefreshTokenError(err error) error {
	return &NoValidRefreshTokenError{
		NoValidError: NewNoValidError("No valid refresh token", err),
	}
}

type NoValidAccessTokenError struct {
	*NoValidError
}

func (e *NoValidAccessTokenError) Error() string {
	return e.NoValidError.Error()
}

func NewNoValidAccessTokenError(err error) error {
	return &NoValidAccessTokenError{
		NoValidError: NewNoValidError("No valid access token", err),
	}
}
