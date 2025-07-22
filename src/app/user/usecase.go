package user

import (
	"github.com/google/uuid"
	"shop/src/api/user/dto"
	"shop/src/domain/order"
	"shop/src/domain/user"
)

type UseCase interface {
	Register(userDto dto.RegisterRequest) (user.User, error)
	Login(usernameOrEmail, password string) (user.User, error)
	Logout(userID uuid.UUID) error
	LogoutAll(userID uuid.UUID) error
	Update(userID uuid.UUID, u user.User) (user.User, error)
	Delete(userID uuid.UUID) (user.User, error)
	Orders(userID uuid.UUID) ([]order.Order, error)
	DeleteOrder(userDTO dto.TestDeleteOrderRequest) (order.Order, error)
	CreateOrder(userDTO dto.TestCreateOrderRequest) (order.Order, error)
}
