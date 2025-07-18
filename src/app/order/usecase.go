package order

import (
	"github.com/google/uuid"
	"shop/src/domain/order"
)

type UseCase interface {
	PlaceOrder(userID uuid.UUID, productItems []*order.Item) (order.Order, error)
	CancelOrder(ID, userID uuid.UUID) (uuid.UUID, error)
	GetByID(ID uuid.UUID) (order.Order, error)
	OrdersByUserID(userID uuid.UUID) ([]order.Order, error)
	CalculateTotalByProductIDs(o order.Order) (order.Order, error)
	GenerateOrderNum(o order.Order) (order.Order, error)
}
