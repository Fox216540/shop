package order

import "github.com/google/uuid"

type Client interface {
	Create(userID uuid.UUID, items []ProductRequest) (order Order, err error)
	Get(userID uuid.UUID, orderID uuid.UUID) (order OrderWithItems, err error)
	GetOrders(userID uuid.UUID) (orders []Order, err error)
	Delete(userID uuid.UUID, orderID uuid.UUID) (ID uuid.UUID, msg string, status string, err error)
}
