package auth

import (
	"github.com/Fox216540/shop/auth-service/app/jwt"
	"github.com/Fox216540/shop/auth-service/app/tokenstorage"
	jwtDomain "github.com/Fox216540/shop/auth-service/domain/jwt"
	"github.com/google/uuid"
)

type service struct {
	jwtUseCase   jwt.UseCase
	tokenStorage tokenstorage.UseCase
}

func NewService(js jwt.UseCase, ts tokenstorage.UseCase) UseCase {
	return &service{
		jwtUseCase:   js,
		tokenStorage: ts,
	}
}

func (s *service) GenerateTokens(userID uuid.UUID) (tokens jwtDomain.Tokens, err error) {
	refreshToken, jti, err := s.jwtUseCase.GenerateRefresh(userID)
	if err != nil {
		return tokens, err
	}
	accessToken, err := s.jwtUseCase.GenerateAccess(userID)
	if err != nil {
		return tokens, err
	}
	if err = s.tokenStorage.Add(jti, userID); err != nil {
		return tokens, err
	}
	return jwtDomain.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *service) DecodeAccessToken(token string) (jwtUser jwtDomain.JWTUser, err error) {
	jwtUser, err = s.jwtUseCase.DecodeAccess(token)
	if err != nil {
		return jwtUser, err
	}
	return jwtUser, nil
}

func (s *service) DecodeRefreshToken(token string) (jwtUser jwtDomain.JWTUser, err error) {
	jwtUser, err = s.jwtUseCase.DecodeRefresh(token)
	if err != nil {
		return jwtUser, err
	}
	return jwtUser, nil
}

func (s *service) DeleteRefreshToken(token string) error {
	jwtUser, err := s.jwtUseCase.DecodeRefresh(token)
	if err != nil {
		return err
	}
	return s.tokenStorage.Delete(jwtUser.JTI, jwtUser.UserID)
}

func (s *service) DeleteAllTokensByToken(token string) error {
	jwtUser, err := s.jwtUseCase.DecodeRefresh(token)
	if err != nil {
		return err
	}
	return s.DeleteAllTokensByUserID(jwtUser.UserID)
}

func (s *service) DeleteAllTokensByUserID(userID uuid.UUID) error {
	return s.tokenStorage.DeleteAll(userID)
}

func (s *service) RefreshTokens(token string) (tokens jwtDomain.Tokens, err error) {
	userJWT, err := s.DecodeRefreshToken(token)
	if err != nil {
		return tokens, err
	}
	if err := s.tokenStorage.Delete(userJWT.JTI, userJWT.UserID); err != nil {
		return tokens, err
	}
	return s.GenerateTokens(userJWT.UserID)
}
