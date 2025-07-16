package jwt

import (
	"github.com/google/uuid"
)

type JWT struct {
	UserID   uuid.UUID
	Username string
}
