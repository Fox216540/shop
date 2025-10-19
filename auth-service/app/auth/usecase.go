package auth

import (
	"github.com/Fox216540/shop/auth-service/domain/auth"
	"github.com/Fox216540/shop/auth-service/domain/jwt"
)

type UseCase interface {
	SignUp(auth auth.Auth) (hash string, tokens jwt.Tokens, err error)
	Login(auth auth.Auth, hash string) (tokens jwt.Tokens, err error)
	DecodeRefreshToken(token string) (userJWT jwt.JWTUser, err error)
	DecodeAccessToken(token string) (userJWT jwt.JWTUser, err error)
	DeleteRefreshToken(token string) error
	DeleteAllTokens(token string) error
	NewPassword(newPassword string) (hash string, err error)
	RefreshTokens(token string) (tokens jwt.Tokens, err error)
}
