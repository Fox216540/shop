package user

import (
	"github.com/google/uuid"
	"shop/src/domain/order"
)

type Service interface {
	Register(u User) (User, error)
	Login(usernameOrEmail, password string) (User, error)
	Logout(userID uuid.UUID) error
	LogoutAll(userID uuid.UUID) error
	Update(u User) (User, error)
	Delete(userID uuid.UUID) error
	Orders(userID uuid.UUID) ([]order.Order, error)
	DeleteOrder(userID, orderID uuid.UUID) (uuid.UUID, error)
	CreateOrder(o order.Order) (order.Order, error)
}
