package auth

import (
	"github.com/google/uuid"
)

type Client interface {
	GenerateTokens(userID uuid.UUID) (Tokens, error)
	DeleteAllRefreshTokens(userID uuid.UUID) error
}
