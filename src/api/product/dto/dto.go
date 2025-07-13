package productdto

import (
	"github.com/google/uuid"
)

type GetProductByIDRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type ProductResponse struct {
	ProductID          uuid.UUID `json:"id"`                                            // Product ID
	ProductName        string    `json:"name" binding:"required,min=3,max=50"`          // Product name
	ProductImg         string    `json:"img" binding:"required"`                        // Product image URL
	ProductPrice       float64   `json:"price" binding:"required,gt=0"`                 // Product price
	ProductCategory    string    `json:"category" binding:"required,min=3,max=20"`      // Product category
	ProductDescription string    `json:"description" binding:"required,min=10,max=200"` // Product description
	ProductStock       int       `json:"stock" binding:"required,gt=0"`                 // Product stock quantity
}

type GetProductsByCategoryRequest struct {
	Category string `form:"category" binding:"omitempty,min=3,max=20"` // Product category
}
