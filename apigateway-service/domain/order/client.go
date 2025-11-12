package order

import "github.com/google/uuid"

type Client interface {
	CreateOrder(userID uuid.UUID, items []ProductRequest) (order Order, err error)
	GetOrder(userID uuid.UUID, orderID uuid.UUID) (order Order, err error)
	GetOrders(userID uuid.UUID) (orders []Order, err error)
	DeleteOrder(userID uuid.UUID, orderID uuid.UUID) (ID uuid.UUID, msg string, status string, err error)
}
