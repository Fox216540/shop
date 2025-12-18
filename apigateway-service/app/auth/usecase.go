package auth

import (
	"github.com/Fox216540/shop/apigateway-service/domain/auth"
)

type UseCase interface {
	LogIn(phoneOrEmail, password string) (name string, tokens auth.Tokens, message string, err error)
	LogOut(token string) (message string, err error)
	LogOutAll(token string) (message string, err error)
	RefreshTokens(token string) (tokens auth.Tokens, err error)
}
