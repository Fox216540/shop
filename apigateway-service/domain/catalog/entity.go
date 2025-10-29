package catalog

import "github.com/google/uuid"

type Category struct {
	ID   uuid.UUID
	Name string
}

type Product struct {
	ID          uuid.UUID
	Name        string
	Img         string
	Price       float64
	CategoryID  uuid.UUID
	Description string
	Stock       uint64
}
