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
	o, err := s.orderClient.Create(userID, items)
	if err != nil {
		return order.Order{}, err
	}
	return o, nil
}

func (s *service) GetOrder(userID uuid.UUID, orderID uuid.UUID) (order.OrderWithItems, error) {
	o, err := s.orderClient.Get(userID, orderID)
	if err != nil {
		return order.OrderWithItems{}, err
	}
	return o, nil
}

func (s *service) GetOrders(userID uuid.UUID) ([]order.Order, error) {
	orders, err := s.orderClient.GetOrders(userID)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (s *service) DeleteOrder(userID uuid.UUID, orderID uuid.UUID) (ID uuid.UUID, message string, status string, err error) {
	id, msg, stat, err := s.orderClient.Delete(userID, orderID)
	if err != nil {
		return uuid.Nil, "", "", err
	}
	return id, msg, stat, nil
}
