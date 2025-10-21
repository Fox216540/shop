package auth

import (
	"github.com/Fox216540/shop/auth-service/domain/jwt"
	"github.com/google/uuid"
)

type UseCase interface {
	GenerateTokens(userID uuid.UUID) (tokens jwt.Tokens, err error)
	DecodeAccessToken(token string) (userJWT jwt.JWTUser, err error)
	DecodeRefreshToken(token string) (userJWT jwt.JWTUser, err error)
	DeleteRefreshToken(token string) error
	DeleteAllTokens(token string) error
	RefreshTokens(token string) (tokens jwt.Tokens, err error)
}
