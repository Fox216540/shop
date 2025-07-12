package productdto

import (
	"github.com/google/uuid"
	"shop/src/domain/product"
)

type GetProductByIDRequest struct {
	ProductID string `json:"id" validate:"required,uuid"`
}

type GetProductByIDResponse struct {
	ProductID          uuid.UUID `json:"id"`                                            // Product ID
	ProductName        string    `json:"name" binding:"required,min=3,max=50"`          // Product name
	ProductImg         string    `json:"img" binding:"required"`                        // Product image URL
	ProductPrice       float64   `json:"price" binding:"required,gt=0"`                 // Product price
	ProductCategory    string    `json:"category" binding:"required,min=3,max=20"`      // Product category
	ProductDescription string    `json:"description" binding:"required,min=10,max=200"` // Product description
	ProductStock       int       `json:"stock" binding:"required,gt=0"`                 // Product stock quantity
}

type GetProductsByCategoryRequest struct {
	Category string `json:"category" binding:"required,min=3,max=20"` // Product category
}

type GetProductsByCategoryResponse struct {
	Products []product.Product `json:"products"` // List of products in the category
}
