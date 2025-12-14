package jwt

import (
	"github.com/Fox216540/shop/auth-service/app/mapError"
	"github.com/Fox216540/shop/auth-service/domain/jwt"
	"github.com/google/uuid"
)

type service struct {
	jwt jwt.Repository
}

func NewService(jwt jwt.Repository) UseCase {
	return &service{
		jwt: jwt,
	}
}

func (s *service) GenerateRefresh(userID uuid.UUID) (string, uuid.UUID, error) {
	token, jti, err := s.jwt.GenerateRefreshToken(userID)
	if err != nil {
		return "", uuid.Nil, mapError.MapError(err, NewInvalidGenerateRefresh(err))
	}
	return token, jti, nil
}

func (s *service) GenerateAccess(userID uuid.UUID) (string, error) {
	token, err := s.jwt.GenerateAccessToken(userID)
	if err != nil {
		return "", mapError.MapError(err, NewInvalidGenerateAccess(err))
	}
	return token, nil
}

func (s *service) DecodeRefresh(token string) (jwt.JWTUser, error) {
	user, err := s.jwt.DecodeRefreshToken(token)
	if err != nil {
		return jwt.JWTUser{}, mapError.MapError(err, NewInvalidDecodeRefresh(err))
	}
	return user, nil
}

func (s *service) DecodeAccess(token string) (jwt.JWTUser, error) {
	user, err := s.jwt.DecodeAccessToken(token)
	if err != nil {
		return jwt.JWTUser{}, mapError.MapError(err, NewInvalidDecodeAccess(err))
	}
	return user, nil
}
