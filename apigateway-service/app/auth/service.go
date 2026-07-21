package auth

import (
	"github.com/Fox216540/shop/apigateway-service/domain/auth"
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
	return s.authClient.LogInUser(phoneOrEmail, password)
}

func (s *service) LogOut(token string) (string, error) {
	return s.authClient.LogOutUser(token)
}

func (s *service) LogOutAll(token string) (string, error) {
	return s.authClient.LogOutAllUser(token)
}

func (s *service) RefreshTokens(token string) (auth.Tokens, error) {
	return s.authClient.RefreshTokensOfUser(token)
}
