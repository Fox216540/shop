package order

import (
	"github.com/Fox216540/shop/apigateway-service/domain/order"
	"github.com/google/uuid"
)

type service struct {
	orderClient order.Client
}

func NewService(client order.Client) UseCase {
	return &service{orderClient: client}
}

func (s *service) CreateOrder(userID uuid.UUID, items []order.ProductRequest) (order.Order, error) {
	return s.orderClient.Create(userID, items)
}

func (s *service) GetOrder(userID uuid.UUID, orderID uuid.UUID) (order.OrderWithItems, error) {
	return s.orderClient.Get(userID, orderID)
}

func (s *service) GetOrders(userID uuid.UUID) ([]order.Order, error) {
	return s.orderClient.GetOrders(userID)
}

func (s *service) DeleteOrder(userID uuid.UUID, orderID uuid.UUID) (ID uuid.UUID, message string, status string, err error) {
	return s.orderClient.Delete(userID, orderID)
}
