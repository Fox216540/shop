package productShort

import "github.com/google/uuid"

type Client interface {
	GetProductByID(id uuid.UUID) (ProductShort, error)
}
