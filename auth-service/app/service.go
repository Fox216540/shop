package app

import (
	"github.com/Fox216540/shop/auth-service/domain/jwt"
	"github.com/Fox216540/shop/auth-service/domain/tokenstorage"
	"github.com/Fox216540/shop/auth-service/domain/user"
	"github.com/google/uuid"
)

type service struct {
	jwtUseCase   jwt.Repository
	tokenStorage tokenstorage.Repository
	userClient   user.Client
}

func NewService(
	jwtUseCase jwt.Repository,
	tokenStorage tokenstorage.Repository,
	userClient user.Client,
) UseCase {
	return &service{
		jwtUseCase:   jwtUseCase,
		tokenStorage: tokenStorage,
		userClient:   userClient,
	}
}

func (s *service) LogIn(phoneOrEmail, password string) (string, jwt.Tokens, error) {
	name, userID, err := s.userClient.VerifyCredentialsOfUser(phoneOrEmail, password)
	if err != nil {
		return "", jwt.Tokens{}, err
	}

	tokens, err := s.GenerateTokens(userID)
	if err != nil {
		return "", jwt.Tokens{}, err
	}

	return name, tokens, nil
}

func (s *service) GenerateTokens(userID uuid.UUID) (jwt.Tokens, error) {
	refreshToken, jti, err := s.jwtUseCase.GenerateRefreshToken(userID)
	if err != nil {
		return jwt.Tokens{}, err
	}

	accessToken, err := s.jwtUseCase.GenerateAccessToken(userID)
	if err != nil {
		return jwt.Tokens{}, err
	}

	if err = s.tokenStorage.Set(jti, userID); err != nil {
		return jwt.Tokens{}, err
	}

	return jwt.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *service) DecodeAccessToken(token string) (jwt.JWTUser, error) {
	return s.jwtUseCase.DecodeAccessToken(token)
}

func (s *service) DecodeRefreshToken(token string) (jwt.JWTUser, error) {
	return s.jwtUseCase.DecodeRefreshToken(token)
}

func (s *service) DeleteRefreshToken(token string) error {
	jwtUser, err := s.jwtUseCase.DecodeRefreshToken(token)
	if err != nil {
		return err
	}

	return s.tokenStorage.Delete(jwtUser.JTI, jwtUser.UserID)
}

func (s *service) DeleteAllTokensByToken(token string) error {
	jwtUser, err := s.jwtUseCase.DecodeRefreshToken(token)
	if err != nil {
		return err
	}

	return s.tokenStorage.DeleteAll(jwtUser.UserID)
}

func (s *service) DeleteAllTokensByUserID(userID uuid.UUID) error {
	return s.tokenStorage.DeleteAll(userID)
}

func (s *service) RefreshTokens(token string) (jwt.Tokens, error) {
	jwtUser, err := s.DecodeRefreshToken(token)
	if err != nil {
		return jwt.Tokens{}, err
	}

	if err = s.tokenStorage.Delete(jwtUser.JTI, jwtUser.UserID); err != nil {
		return jwt.Tokens{}, err
	}

	return s.GenerateTokens(jwtUser.UserID)
}
