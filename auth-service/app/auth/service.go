package auth

import (
	"github.com/Fox216540/shop/auth-service/app/hasher"
	"github.com/Fox216540/shop/auth-service/app/jwt"
	"github.com/Fox216540/shop/auth-service/app/tokenstorage"
	"github.com/Fox216540/shop/auth-service/domain/auth"
	jwtDomain "github.com/Fox216540/shop/auth-service/domain/jwt"
	"github.com/google/uuid"
)

type service struct {
	jwtUseCase   jwt.UseCase
	hasher       hasher.UseCase
	tokenStorage tokenstorage.UseCase
}

func NewService(js jwt.UseCase, hs hasher.UseCase, ts tokenstorage.UseCase) UseCase {
	return &service{
		jwtUseCase:   js,
		hasher:       hs,
		tokenStorage: ts,
	}
}

func (s *service) generateTokens(userID uuid.UUID) (tokens jwtDomain.Tokens, err error) {

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

func (s *service) SignUp(a auth.Auth) (hash string, tokens jwtDomain.Tokens, err error) {
	hash, err = s.NewPassword(a.Password)
	if err != nil {
		return "", tokens, err
	}

	tokens, err = s.generateTokens(a.UserID)
	if err != nil {
		return "", tokens, err
	}
	return hash, tokens, nil
}

func (s *service) Login(a auth.Auth, hash string) (tokens jwtDomain.Tokens, err error) {
	if err := s.hasher.VerifyPass(a.Password, hash); err != nil {
		return tokens, err
	}

	tokens, err = s.generateTokens(a.UserID)
	if err != nil {
		return tokens, err
	}
	return tokens, nil
}

func (s *service) DecodeRefreshToken(token string) (jwtUser jwtDomain.JWTUser, err error) {
	jwtUser, err = s.jwtUseCase.DecodeRefresh(token)
	if err != nil {
		return jwtUser, err
	}
	return jwtUser, nil
}

func (s *service) DecodeAccessToken(token string) (jwtUser jwtDomain.JWTUser, err error) {
	jwtUser, err = s.jwtUseCase.DecodeAccess(token)
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

func (s *service) DeleteAllTokens(token string) error {
	jwtUser, err := s.jwtUseCase.DecodeRefresh(token)
	if err != nil {
		return err
	}
	return s.tokenStorage.DeleteAll(jwtUser.UserID)
}

func (s *service) NewPassword(newPassword string) (hash string, err error) {
	return s.hasher.HashPass(newPassword)
}

func (s *service) RefreshTokens(token string) (tokens jwtDomain.Tokens, err error) {
	jwtUser, err := s.jwtUseCase.DecodeRefresh(token)
	if err != nil {
		return jwtDomain.Tokens{}, err
	}
	if err := s.tokenStorage.Delete(jwtUser.JTI, jwtUser.UserID); err != nil {
		return jwtDomain.Tokens{}, err
	}
	tokens, err = s.generateTokens(jwtUser.UserID)
	return tokens, nil
}
