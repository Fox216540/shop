package jwt

import (
	"github.com/google/uuid"
)

type User struct {
	UserID   uuid.UUID
	Username string
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}
