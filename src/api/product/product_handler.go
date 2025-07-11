package catalog_handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CatalogHandler(r *gin.Engine) {
	r.GET("/catalog", func(c *gin.Context) {
		// Simulate fetching product data
		catalogData := "This is product data"
		c.JSON(http.StatusOK, catalogData)
	})
}
