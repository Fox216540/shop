package order

import "github.com/google/uuid"

type Repository interface {
	Save(o Order) (Order, error)
	Remove(id uuid.UUID) error
}
