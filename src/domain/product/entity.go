package product

import "github.com/google/uuid"

type Product struct {
	ID          uuid.UUID `json:"id"`                                            // Product ID
	Name        string    `json:"name" binding:"required,min=3,max=50"`          // Product name
	Img         string    `json:"img" binding:"required"`                        // Product image URL
	Price       float64   `json:"price" binding:"required,gt=0"`                 // Product price
	Category    string    `json:"category" binding:"required,min=3,max=20"`      // Product category
	Description string    `json:"description" binding:"required,min=10,max=200"` // Product description
	Stock       int       `json:"stock" binding:"required,gt=0"`                 // Product stock quantity
}
