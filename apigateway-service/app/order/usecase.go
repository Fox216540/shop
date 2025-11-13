package order

import (
	"github.com/Fox216540/shop/apigateway-service/domain/order"
	"github.com/google/uuid"
)

type UseCase interface {
	Create(userID uuid.UUID, items []order.ProductRequest) (order.Order, error)
	Get(userID uuid.UUID, orderID uuid.UUID) (order.Order, error)
	GetOrders(userID uuid.UUID) ([]order.Order, error)
	Delete(userID uuid.UUID, orderID uuid.UUID) (ID uuid.UUID, msg string, status string, err error)
}
