package order

import (
	"github.com/Fox216540/shop/order-service/app/dto"
	"github.com/Fox216540/shop/order-service/domain/order"
	"github.com/google/uuid"
)

type UseCase interface {
	PlaceOrder(userID uuid.UUID, orderItems []dto.OrderItems) (order.Order, error)
	DeleteOrder(ID, userID uuid.UUID) error
	GetOrderByIDAndUserID(ID, userID uuid.UUID) (order.Order, error)
	GetOrdersByUserID(userID uuid.UUID) (orders []order.Order, err error)
}
