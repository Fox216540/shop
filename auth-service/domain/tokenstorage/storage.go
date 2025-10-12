package tokenstorage

import (
	"github.com/google/uuid"
)

type Repository interface {
	Set(jti, userID uuid.UUID) error
	Exists(jti uuid.UUID) error
	Delete(jti, userID uuid.UUID) error
	DeleteAll(userID uuid.UUID) error
}
