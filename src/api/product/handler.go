package product

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"shop/src/api/product/di"
	"shop/src/api/product/dto"
	"shop/src/app/product"
)

func Handler(r *gin.Engine) {
	ps := di.GetProductService()
	// GetProductsByCategory
	r.GET("/catalog/", getProductsByCategoryHandler(ps))
	// GetProductByID
	r.GET("/catalog/:id", getProductByIDHandler(ps))
}

func parseCategoryID(c *gin.Context) (*uuid.UUID, error) {
	var req dto.GetProductsByCategoryRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		return nil, err
	}
	if req.CategoryID == "" {
		return nil, nil
	}
	uid, err := uuid.Parse(req.CategoryID)
	if err != nil {
		return nil, err
	}
	return &uid, nil
}

func getProductsByCategoryHandler(ps product.UseCase) gin.HandlerFunc {
	//GetProductsByCategory
	return func(c *gin.Context) {
		categoryID, err := parseCategoryID(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID", "details": err.Error()})
			return
		}

		products, err := ps.ProductsOfCategoryID(categoryID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch product data"})
			return
		}
		var productsDTO []dto.ProductResponse
		for _, p := range products {
			productsDTO = append(productsDTO, dto.ProductResponse{
				ProductID:          p.ID.String(),
				ProductName:        p.Name,
				ProductImg:         p.Img,
				ProductPrice:       p.Price,
				ProductCategoryID:  p.CategoryID.String(),
				ProductDescription: p.Description,
				ProductStock:       p.Stock,
			})
		}

		c.JSON(http.StatusOK, productsDTO)
	}
}

func getProductByIDHandler(ps product.UseCase) gin.HandlerFunc {
	//GetProductByID
	return func(c *gin.Context) {
		var uri dto.GetProductByIDRequest
		if err := c.ShouldBindUri(&uri); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID", "details": err.Error()})
			return
		}

		productID, _ := uuid.Parse(uri.ID)

		Product, err := ps.ProductByID(productID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(http.StatusOK, dto.ProductResponse{
			ProductID:          Product.ID.String(),
			ProductName:        Product.Name,
			ProductImg:         Product.Img,
			ProductPrice:       Product.Price,
			ProductCategoryID:  Product.CategoryID.String(),
			ProductDescription: Product.Description,
			ProductStock:       Product.Stock,
		})
	}
}
