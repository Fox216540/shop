package product

import "github.com/google/uuid"

type Product struct {
	ID          uuid.UUID // Product ID
	Name        string    // Product name
	Img         string    // Product image URL
	Price       float64   // Product price
	CategoryID  uuid.UUID // Product category
	Description string    // Product description
	Stock       int       // Product stock quantity
}
