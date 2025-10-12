package tokenstorage

import (
	"github.com/google/uuid"
)

type UseCase interface {
	Add(jti, userID uuid.UUID) error
	Delete(jti, userID uuid.UUID) error
	DeleteAll(userID uuid.UUID) error
	Exists(jti uuid.UUID) error
}
