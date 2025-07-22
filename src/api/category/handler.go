package category

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/src/api/category/di"
	"shop/src/api/category/dto"
	"shop/src/app/category"
)

func Handler(r *gin.Engine) {
	cs := di.GetCategoryService()
	// GetCategories
	r.GET("/categories", getCategoriesHandler(cs))
}

func getCategoriesHandler(cs category.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		categories, err := cs.GetCategories()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch category data"})
			return
		}
		categoriesDTO := make([]dto.CategoryResponse, 0, len(categories))
		for _, Category := range categories {
			categoriesDTO = append(categoriesDTO, dto.CategoryResponse{
				ID:   Category.ID,
				Name: Category.Name,
			})
		}
		c.JSON(http.StatusOK, categoriesDTO)
	}
}
