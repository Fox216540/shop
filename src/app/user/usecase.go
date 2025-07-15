package userservice

import (
	"github.com/google/uuid"
	"shop/src/domain/order"
	"shop/src/domain/user"
)

type UseCase interface {
	Register(u user.User) (user.User, error)
	Login(usernameOrEmail, password string) (user.User, error)
	Logout(userID uuid.UUID) error
	LogoutAll(userID uuid.UUID) error
	Update(u user.User) (user.User, error)
	Delete(userID uuid.UUID) error
	Orders(userID uuid.UUID) ([]order.Order, error)
	DeleteOrder(userID, orderID uuid.UUID) (uuid.UUID, error)
	CreateOrder(o order.Order) (order.Order, error)
}
