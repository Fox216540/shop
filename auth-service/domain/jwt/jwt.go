package jwt

import (
	"github.com/google/uuid"
)

type Repository interface {
	GenerateRefreshToken(userID uuid.UUID) (token string, jti uuid.UUID, err error)
	GenerateAccessToken(userID uuid.UUID) (token string, err error)
	DecodeRefreshToken(token string) (userIDWithJTI JWTUser, err error)
	DecodeAccessToken(token string) (userIDWithJTI JWTUser, err error)
}
