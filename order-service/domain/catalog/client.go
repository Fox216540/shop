package catalog

import "github.com/google/uuid"

type Client interface {
	GetProductsByIDs(ids []uuid.UUID) ([]Product, error)
}
