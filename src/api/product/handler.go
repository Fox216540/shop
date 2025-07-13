package producthandler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	di "shop/src/api/product/di"
	dto "shop/src/api/product/dto"
)

func CatalogHandler(r *gin.Engine) {
	r.GET("/catalog/", func(c *gin.Context) {
		category := c.Query("category")
		var categoryPtr *string
		if category != "" {
			categoryPtr = &category
		} else {
			categoryPtr = nil
		}
		ps := di.GetProductService()
		products, err := ps.ProductsOfCategory(categoryPtr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch product data"})
			return
		}
		c.JSON(http.StatusOK, dto.GetProductsByCategoryResponse{Products: products})
	})
	r.GET("/catalog/:id", func(c *gin.Context) {
		productID, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
			return
		}
		ps := di.GetProductService()
		product, err := ps.ProductById(productID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(http.StatusOK, dto.GetProductByIDResponse{
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
