package jwt

import (
	"github.com/google/uuid"
)

type Service interface {
	GenerateRefreshToken(userID uuid.UUID) (string, error)
	GenerateAccessToken(userID uuid.UUID, username string) (string, error)
	DecodeRefreshToken(token string) (JWT, error)
	DecodeAccessToken(token string) (JWT, error)
}
