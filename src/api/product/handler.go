package producthandler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	di "shop/src/api/product/di"
	dto "shop/src/api/product/dto"
)

func ProductHandler(r *gin.Engine) {
	// GetProductsByCategory
	r.GET("/catalog/", func(c *gin.Context) {
		var r dto.GetProductsByCategoryRequest

		if err := c.ShouldBindQuery(&r); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters", "details": err.Error()})
			return
		}

		var categoryPtr *string

		if r.Category != "" {
			categoryPtr = &r.Category
		} else {
			categoryPtr = nil
		}

		ps := di.GetProductService()
		products, err := ps.ProductsOfCategory(categoryPtr)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch product data"})
			return
		}
		var productsDTO []dto.ProductResponse
		for _, p := range products {
			productsDTO = append(productsDTO, dto.ProductResponse{
				ProductID:          p.ID,
				ProductName:        p.Name,
				ProductImg:         p.Img,
				ProductPrice:       p.Price,
				ProductCategory:    p.Category,
				ProductDescription: p.Description,
				ProductStock:       p.Stock,
			})
		}

		c.JSON(http.StatusOK, productsDTO)
	})

	// GetProductByID
	r.GET("/catalog/:id", func(c *gin.Context) {
		var uri dto.GetProductByIDRequest
		if err := c.ShouldBindUri(&uri); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID", "details": err.Error()})
			return
		}

		productID, _ := uuid.Parse(uri.ID)

		ps := di.GetProductService()
		product, err := ps.ProductById(productID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(http.StatusOK, dto.ProductResponse{
			ProductID:          product.ID,
			ProductName:        product.Name,
			ProductImg:         product.Img,
			ProductPrice:       product.Price,
			ProductCategory:    product.Category,
			ProductDescription: product.Description,
			ProductStock:       product.Stock,
		})
	})
}
