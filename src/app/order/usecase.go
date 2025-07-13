package orderservice

import (
	"github.com/google/uuid"
	"shop/src/domain/order"
)

type Service interface {
	PlaceOrder(o order.Order) (order.Order, error)
	CancelOrder(ID, userID uuid.UUID) (uuid.UUID, error)
	GetByID(ID uuid.UUID) (order.Order, error)
	OrdersByUserID(userID uuid.UUID) ([]order.Order, error) //Для domain User
	CalculateTotalByIDs(o order.Order) (order.Order, error)
	GenerateOrderNum(o order.Order) order.Order
}
