package jwt

import (
	"shop/src/core/exception"
)

const (
	domain = "JWT"
)

type DomainBadRequestError struct {
	*exception.BadRequestError
}

func (e *DomainBadRequestError) Error() string {
	return e.BadRequestError.Error()
}

func NewDomainBadRequestError(msg string, err error) *DomainBadRequestError {
	return &DomainBadRequestError{
		BadRequestError: exception.NewBadRequestError(msg, domain, err),
	}
}

type BadRefreshTokenError struct {
	*DomainBadRequestError
}

func (e *BadRefreshTokenError) Error() string {
	return e.BadRequestError.Error()
}

func NewBadRefreshTokenError(err error) error {
	return &BadRefreshTokenError{
		DomainBadRequestError: NewDomainBadRequestError("Bad refresh token", err),
	}
}

type BadAccessTokenError struct {
	*DomainBadRequestError
}

func (e *BadAccessTokenError) Error() string {
	return e.BadRequestError.Error()
}

func NewBadAccessTokenError(err error) error {
	return &BadAccessTokenError{
		DomainBadRequestError: NewDomainBadRequestError("Bad access token", err),
	}
}

type DomainUnauthorizedError struct {
	*exception.UnauthorizedError
}

func (e *DomainUnauthorizedError) Error() string {
	return e.UnauthorizedError.Error()
}

func NewDomainUnauthorizedError(msg string, err error) *DomainUnauthorizedError {
	return &DomainUnauthorizedError{
		UnauthorizedError: exception.NewUnauthorizedError(msg, domain, err),
	}
}

type NoValidRefreshTokenError struct {
	*DomainUnauthorizedError
}

func (e *NoValidRefreshTokenError) Error() string {
	return e.DomainUnauthorizedError.Error()
}

func NewNoValidRefreshTokenError(err error) error {
	return &NoValidRefreshTokenError{
		DomainUnauthorizedError: NewDomainUnauthorizedError("No valid refresh token", err),
	}
}

type NoValidAccessTokenError struct {
	*DomainUnauthorizedError
}

func (e *NoValidAccessTokenError) Error() string {
	return e.DomainUnauthorizedError.Error()
}

func NewNoValidAccessTokenError(err error) error {
	return &NoValidAccessTokenError{
		DomainUnauthorizedError: NewDomainUnauthorizedError("No valid access token", err),
	}
}
