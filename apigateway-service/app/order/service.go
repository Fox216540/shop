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

func (s *service) Create(userID uuid.UUID, items []order.ProductRequest) (order.Order, error) {
	o, err := s.orderClient.CreateOrder(userID, items)
	if err != nil {
		return order.Order{}, err
	}
	return o, nil
}

func (s *service) Get(userID uuid.UUID, orderID uuid.UUID) (order.Order, error) {
	o, err := s.orderClient.GetOrder(userID, orderID)
	if err != nil {
		return order.Order{}, err
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

func (s *service) Delete(userID uuid.UUID, orderID uuid.UUID) (ID uuid.UUID, message string, status string, err error) {
	id, msg, stat, err := s.orderClient.DeleteOrder(userID, orderID)
	if err != nil {
		return uuid.Nil, "", "", err
	}
	return id, msg, stat, nil
}
