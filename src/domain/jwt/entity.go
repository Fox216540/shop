package jwt

import (
	"github.com/google/uuid"
)

type JWTUser struct {
	UserID   uuid.UUID
	Username string
	JTI      uuid.UUID
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}
