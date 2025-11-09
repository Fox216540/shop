package user

import (
	"github.com/Fox216540/shop/apigateway-service/app/dto"
	"github.com/Fox216540/shop/apigateway-service/domain/auth"
	"github.com/Fox216540/shop/apigateway-service/domain/user"
	userDomain "github.com/Fox216540/shop/apigateway-service/domain/user"
	"github.com/google/uuid"
)

type service struct {
	userClient user.Client
}

func NewService(userClient user.Client) UseCase {
	return &service{
		userClient: userClient,
	}
}

func (s *service) RegisterUser(dto dto.User) (string, auth.Tokens, string, error) {
	u := userDomain.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Phone:    dto.Phone,
		Password: dto.Password,
		Address:  dto.Address,
	}
	name, tokens, message, err := s.userClient.RegisterUser(u)
	if err != nil {
		return "", auth.Tokens{}, "", err
	}
	return name, tokens, message, err
}

func (s *service) DeleteUser(id uuid.UUID) (string, error) {
	msg, err := s.userClient.DeleteUser(id)
	if err != nil {
		return "", err
	}
	return msg, nil
}
