package app

import (
	"github.com/Fox216540/shop/auth-service/domain/jwt"
	"github.com/google/uuid"
)

type UseCase interface {
	LogIn(phoneOrEmail, password string) (name string, tokens jwt.Tokens, err error)
	GenerateTokens(userID uuid.UUID) (tokens jwt.Tokens, err error)
	DecodeAccessToken(token string) (userJWT jwt.JWTUser, err error)
	DecodeRefreshToken(token string) (userJWT jwt.JWTUser, err error)
	DeleteRefreshToken(token string) error
	DeleteAllTokensByToken(token string) error
	DeleteAllTokensByUserID(userID uuid.UUID) error
	RefreshTokens(token string) (tokens jwt.Tokens, err error)
}
