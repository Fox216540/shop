package order

import (
	"github.com/google/uuid"
	"shop/src/domain/order"
)

type UseCase interface {
	Place(userID uuid.UUID, productItems []*order.Item) (order.Order, error)
	Cancel(ID, userID uuid.UUID) (uuid.UUID, error)
	GetByID(ID uuid.UUID) (order.Order, error)
	GetOrdersByUserID(userID uuid.UUID) ([]order.Order, error)
}
