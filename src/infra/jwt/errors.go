package jwt

import (
	"shop/src/domain/jwt"
)

const layer = "Infra"

type ServerError struct {
	*jwt.GlobalError
}

func (e *ServerError) Error() string {
	return e.GlobalError.Error()
}

func NewServerError(msg string, err error) *ServerError {
	return &ServerError{
		GlobalError: jwt.NewGlobalError(msg, err, layer),
	}
}

type InvalidGenerateAccessToken struct {
	*ServerError
}

func (e *InvalidGenerateAccessToken) Error() string {
	return e.ServerError.Error()
}

func NewInvalidGenerateAccessToken(err error) error {
	return &InvalidGenerateAccessToken{
		ServerError: NewServerError("Invalid Generate Access Token Error", err),
	}
}

type InvalidGenerateRefreshToken struct {
	*ServerError
}

func (e *InvalidGenerateRefreshToken) Error() string {
	return e.ServerError.Error()
}

func NewInvalidGenerateRefreshToken(err error) error {
	return &InvalidGenerateRefreshToken{
		ServerError: NewServerError("Invalid Generate Refresh Token Error", err),
	}
}
