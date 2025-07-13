package producthandler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	di "shop/src/api/product/di"
	dto "shop/src/api/product/dto"
)

func CatalogHandler(r *gin.Engine) {
	r.GET("/catalog", func(c *gin.Context) {
		// Simulate fetching product data
		ps := di.GetProductService()
		products, err := ps.ProductsOfCategory("electronics")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch product data"})
			return
		}
		c.JSON(http.StatusOK, dto.GetProductsByCategoryResponse{Products: products})
	})
}
