package jwt

import (
	"github.com/google/uuid"
)

type Service interface {
	GenerateRefreshToken(userID uuid.UUID) (string, uuid.UUID, error)
	GenerateAccessToken(userID uuid.UUID, username string) (string, error)
	DecodeRefreshToken(token string) (User, error)
	DecodeAccessToken(token string) (User, error)
}
