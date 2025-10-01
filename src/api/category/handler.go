package category

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/src/api/category/di"
	"shop/src/api/category/dto"
	"shop/src/app/category"
	"shop/src/core/mapError"
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
			status, message := mapError.MapError(err)
			c.JSON(status, gin.H{"error": message})
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
