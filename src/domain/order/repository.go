package order

import "github.com/google/uuid"

type Repository interface {
	Save(o Order) (Order, error)
	Remove(ID, userID uuid.UUID) error
	GetByID(ID uuid.UUID) (Order, error)
	CheckOrderNum(orderNum string) error
	GetOrdersByUserID(userID uuid.UUID) ([]Order, error)
}
