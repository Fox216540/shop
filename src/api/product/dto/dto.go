package dto

// Response

type ProductResponse struct {
	ProductID          string  `json:"id"`          // Product ID
	ProductName        string  `json:"name"`        // Product name
	ProductImg         string  `json:"img"`         // Product image URL
	ProductPrice       float64 `json:"price"`       // Product price
	ProductCategoryID  string  `json:"category_id"` // Product category
	ProductDescription string  `json:"description"` // Product description
	ProductStock       int     `json:"stock"`       // Product stock quantity
}

// Request

type GetProductByIDRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type GetProductsByCategoryRequest struct {
	CategoryID string `form:"category" binding:"omitempty,uuid"` // Product category
}
