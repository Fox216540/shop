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
	name, tokens, message, err := s.userClient.Register(u)
	if err != nil {
		return "", auth.Tokens{}, "", err
	}
	return name, tokens, message, err
}

func (s *service) DeleteUser(id uuid.UUID) (string, error) {
	msg, err := s.userClient.Delete(id)
	if err != nil {
		return "", err
	}
	return msg, nil
}

func (s *service) UpdateEmailOfUser(id uuid.UUID, email string) (message string, err error) {
	msg, err := s.userClient.UpdateEmail(id, email)
	if err != nil {
		return "", err
	}
	return msg, nil
}

func (s *service) UpdatePasswordOfUser(id uuid.UUID, password string) (message string, err error) {
	msg, err := s.userClient.UpdatePassword(id, password)
	if err != nil {
		return "", err
	}
	return msg, nil
}

func (s *service) UpdatePhoneOfUser(id uuid.UUID, phone string) (message string, err error) {
	msg, err := s.userClient.UpdatePhone(id, phone)
	if err != nil {
		return "", err
	}
	return msg, nil
}

func (s *service) UpdateProfileOfUser(id uuid.UUID, name *string, address *string) (message string, newName string, err error) {
	newName, msg, err := s.userClient.UpdateProfile(id, name, address)
	if err != nil {
		return "", "", err
	}
	return msg, newName, nil
}
