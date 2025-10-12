package jwt

import (
	"github.com/Fox216540/shop/auth-service/domain/jwt"
	"github.com/google/uuid"
)

type UseCase interface {
	GenerateRefresh(userID uuid.UUID) (token string, jti uuid.UUID, err error)
	GenerateAccess(userID uuid.UUID) (token string, err error)
	DecodeRefresh(token string) (jwt.JWTUser, error)
	DecodeAccess(token string) (jwt.JWTUser, error)
}
