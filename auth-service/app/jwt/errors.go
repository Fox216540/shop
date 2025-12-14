package jwt

import (
	"github.com/Fox216540/shop/auth-service/app/globalError"
)

const domain = "JWT"

type JWTError struct {
	*globalError.AppServerError
}

func NewJWTError(msg string, err error) *JWTError {
	return &JWTError{
		AppServerError: globalError.NewAppServerError(msg, domain, err),
	}
}

func (e *JWTError) Error() string {
	return e.AppServerError.Error()
}

func (e *JWTError) Unwrap() error {
	return e.AppServerError
}

type InvalidGenerateRefresh struct {
	*JWTError
}

func NewInvalidGenerateRefresh(err error) *InvalidGenerateRefresh {
	return &InvalidGenerateRefresh{
		JWTError: NewJWTError("Invalid generate refresh", err),
	}
}

func (e *InvalidGenerateRefresh) Error() string {
	return e.JWTError.Error()
}

func (e *InvalidGenerateRefresh) Unwrap() error {
	return e.JWTError
}

type InvalidGenerateAccess struct {
	*JWTError
}

func NewInvalidGenerateAccess(err error) *InvalidGenerateAccess {
	return &InvalidGenerateAccess{
		JWTError: NewJWTError("Invalid generate access", err),
	}
}

func (e *InvalidGenerateAccess) Error() string {
	return e.JWTError.Error()
}

func (e *InvalidGenerateAccess) Unwrap() error {
	return e.JWTError
}

type InvalidDecodeRefresh struct {
	*JWTError
}

func NewInvalidDecodeRefresh(err error) *InvalidDecodeRefresh {
	return &InvalidDecodeRefresh{
		JWTError: NewJWTError("Invalid decode refresh", err),
	}
}

func (e *InvalidDecodeRefresh) Error() string {
	return e.JWTError.Error()
}

func (e *InvalidDecodeRefresh) Unwrap() error {
	return e.JWTError
}

type InvalidDecodeAccess struct {
	*JWTError
}

func NewInvalidDecodeAccess(err error) *InvalidDecodeAccess {
	return &InvalidDecodeAccess{
		JWTError: NewJWTError("Invalid decode access", err),
	}
}

func (e *InvalidDecodeAccess) Error() string {
	return e.JWTError.Error()
}

func (e *InvalidDecodeAccess) Unwrap() error {
	return e.JWTError
}
