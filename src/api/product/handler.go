package producthandler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	di "shop/src/api/product/di"
	dto "shop/src/api/product/dto"
)

func ProductHandler(r *gin.Engine) {
	ps := di.GetProductService()
	// GetProductsByCategory
	r.GET("/catalog/", func(c *gin.Context) {
		var r dto.GetProductsByCategoryRequest

		if err := c.ShouldBindQuery(&r); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters", "details": err.Error()})
			return
		}

		var categoryPtr *uuid.UUID

		if r.CategoryID != "" {
			categoryUUID, _ := uuid.Parse(r.CategoryID)
			categoryPtr = &categoryUUID
		} else {
			categoryPtr = nil
		}

		products, err := ps.ProductsOfCategoryID(categoryPtr)

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
	})

	// GetProductByID
	r.GET("/catalog/:id", func(c *gin.Context) {
		var uri dto.GetProductByIDRequest
		if err := c.ShouldBindUri(&uri); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID", "details": err.Error()})
			return
		}

		productID, _ := uuid.Parse(uri.ID)

		product, err := ps.ProductByID(productID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(http.StatusOK, dto.ProductResponse{
			ProductID:          product.ID.String(),
			ProductName:        product.Name,
			ProductImg:         product.Img,
			ProductPrice:       product.Price,
			ProductCategoryID:  product.CategoryID.String(),
			ProductDescription: product.Description,
			ProductStock:       product.Stock,
		})
	})
}
