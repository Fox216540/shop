package user

import (
	"github.com/Fox216540/shop/auth-service/domain/user"
	"github.com/google/uuid"
)

type Service struct {
	u user.Client
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) VerifyUser(phoneOrEmail, password string) (string, uuid.UUID, error) {
	name, id, err := s.u.VerifyCredentialsOfUser(phoneOrEmail, password)
	if err != nil {
		return "", uuid.Nil, err
	}
	return name, id, nil
}
