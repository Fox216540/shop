package order

import (
	"github.com/Fox216540/shop/apigateway-service/domain/order"
	"github.com/google/uuid"
)

type UseCase interface {
	CreateOrder(userID uuid.UUID, items []order.ProductRequest) (order.Order, error)
	GetOrder(userID uuid.UUID, orderID uuid.UUID) (order.OrderWithItems, error)
	GetOrders(userID uuid.UUID) ([]order.Order, error)
	DeleteOrder(userID uuid.UUID, orderID uuid.UUID) (ID uuid.UUID, msg string, status string, err error)
}
