package auth

import "github.com/google/uuid"

type Auth struct {
	UserID   uuid.UUID
	Password string
}
