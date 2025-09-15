package jwt

import (
	"shop/src/infra/globalError"
)

const domain = "JWT"

type JWTServerError struct {
	*globalError.InfraServerError
}

func (e *JWTServerError) Error() string {
	return e.InfraServerError.Error()
}

func NewJWTServerError(msg string, err error) *JWTServerError {
	return &JWTServerError{
		InfraServerError: globalError.NewInfraServerError(msg, domain, err),
	}
}

type InvalidGenerateAccessToken struct {
	*JWTServerError
}

func (e *InvalidGenerateAccessToken) Error() string {
	return e.JWTServerError.Error()
}

func NewInvalidGenerateAccessToken(err error) error {
	return &InvalidGenerateAccessToken{
		JWTServerError: NewJWTServerError("Invalid Generate Access Token Error", err),
	}
}

type InvalidGenerateRefreshToken struct {
	*JWTServerError
}

func (e *InvalidGenerateRefreshToken) Error() string {
	return e.JWTServerError.Error()
}

func NewInvalidGenerateRefreshToken(err error) error {
	return &InvalidGenerateRefreshToken{
		JWTServerError: NewJWTServerError("Invalid Generate Refresh Token Error", err),
	}
}
