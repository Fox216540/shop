package order

import (
	"github.com/Fox216540/shop/order-service/app/dto"
	"github.com/Fox216540/shop/order-service/domain/order"
	"github.com/google/uuid"
)

type UseCase interface {
	Place(userID uuid.UUID, orderItems []dto.OrderItems) (order.Order, error)
	Cancel(ID, userID uuid.UUID) (orderID uuid.UUID, err error)
	GetByID(ID uuid.UUID) (order.Order, error)
	GetOrdersByUserID(userID uuid.UUID) (orders []order.Order, err error)
}
