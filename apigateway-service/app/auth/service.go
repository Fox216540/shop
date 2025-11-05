package auth

import (
	"github.com/Fox216540/shop/apigateway-service/domain/auth"
	"github.com/google/uuid"
)

type service struct {
	authClient auth.Client
}

func NewService(authClient auth.Client) UseCase {
	return &service{
		authClient: authClient,
	}
}

func (s *service) LogIn(phoneOrEmail, password string) (string, auth.Tokens, string, error) {
	name, tokens, msg, err := s.authClient.LogInUser(phoneOrEmail, password)
	if err != nil {
		return "", auth.Tokens{}, "", err
	}
	return name, tokens, msg, nil
}

func (s *service) LogOut(token string) (string, error) {
	msg, err := s.authClient.LogOutUser(token)
	if err != nil {
		return "", err
	}
	return msg, nil
}

func (s *service) LogOutAll(token string) (string, error) {
	msg, err := s.authClient.LogOutAllUser(token)
	if err != nil {
		return "", err
	}
	return msg, nil
}

func (s *service) RefreshTokens(token string) (auth.Tokens, error) {
	tokens, err := s.authClient.RefreshTokensOfUser(token)
	if err != nil {
		return auth.Tokens{}, err
	}
	return tokens, nil
}

func (s *service) DecodeAccessToken(token string) (uuid.UUID, error) {
	userID, err := s.authClient.DecodeAccessTokenOfUser(token)
	if err != nil {
		return uuid.Nil, err
	}
	return userID, nil
}
