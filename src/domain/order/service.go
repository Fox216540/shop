package order

import "github.com/google/uuid"

type Service interface {
	PlaceOrder(o Order) (Order, error)
	CancelOrder(id uuid.UUID) (Order, error)
}
