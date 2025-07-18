package category

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/src/api/category/di"
	"shop/src/api/category/dto"
)

func Handler(r *gin.Engine) {
	cs := di.GetCategoryService()
	r.GET("/categories", func(c *gin.Context) {
		categories, err := cs.GetCategories()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch category data"})
			return
		}
		categoriesDTO := make([]dto.CategoryResponse, 0, len(categories))
		for _, category := range categories {
			categoriesDTO = append(categoriesDTO, dto.CategoryResponse{
				ID:   category.ID,
				Name: category.Name,
			})
		}
		c.JSON(http.StatusOK, categoriesDTO)
	})
}
