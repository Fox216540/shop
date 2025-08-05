package user

import (
	"github.com/google/uuid"
	"shop/src/api/user/dto"
	"shop/src/domain/jwt"
	"shop/src/domain/order"
	"shop/src/domain/user"
)

type UseCase interface {
	Register(userDto dto.RegisterRequest) (user.User, jwt.Tokens, error)
	Login(usernameOrEmail, password string) (user.User, jwt.Tokens, error)
	Logout(token string) error
	LogoutAll(token string) error
	UpdateWithoutUsername(userID uuid.UUID, u user.User) (user.User, error)
	UpdateUsername(userID uuid.UUID, username string) (user.User, error)
	Delete(userID uuid.UUID) (user.User, error)
	Orders(userID uuid.UUID) ([]order.Order, error)
	DeleteOrder(userDTO dto.TestDeleteOrderRequest) (order.Order, error)
	CreateOrder(userDTO dto.TestCreateOrderRequest) (order.Order, error)
}
