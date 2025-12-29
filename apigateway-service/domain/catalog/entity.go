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
	Price       string
	Currency    string
	CategoryID  uuid.UUID
	Description string
	Stock       uint64
}
