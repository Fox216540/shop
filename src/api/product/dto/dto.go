package productdto

type GetProductByIDRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type NewProductResponse struct {
	ProductID          string  `json:"id"`          // Product ID
	ProductName        string  `json:"name"`        // Product name
	ProductImg         string  `json:"img"`         // Product image URL
	ProductPrice       float64 `json:"price"`       // Product price
	ProductCategory    string  `json:"category"`    // Product category
	ProductDescription string  `json:"description"` // Product description
	ProductStock       int     `json:"stock"`       // Product stock quantity
}

type GetProductsByCategoryRequest struct {
	Category string `form:"category" binding:"omitempty,min=3,max=20"` // Product category
}
