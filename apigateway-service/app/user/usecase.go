package user

import (
	"github.com/Fox216540/shop/apigateway-service/app/dto"
	"github.com/Fox216540/shop/apigateway-service/domain/auth"
	"github.com/google/uuid"
)

type UseCase interface {
	RegisterUser(user dto.User) (name string, tokens auth.Tokens, message string, err error)
	DeleteUser(id uuid.UUID) (message string, err error)
	UpdateEmailOfUser(id uuid.UUID, email string) (message string, err error)
	UpdatePasswordOfUser(id uuid.UUID, password string) (message string, err error)
	UpdatePhoneOfUser(id uuid.UUID, phone string) (message string, err error)
	UpdateProfileOfUser(id uuid.UUID, name *string, address *string) (message string, newName string, err error)
}
