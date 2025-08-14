package jwt

import (
	"github.com/google/uuid"
)

type Service interface {
	GenerateRefreshToken(userID uuid.UUID) (string, uuid.UUID, error)
	GenerateAccessToken(userID uuid.UUID) (string, error)
	DecodeRefreshToken(token string) (JWTUser, error)
	DecodeAccessToken(token string) (JWTUser, error)
}
