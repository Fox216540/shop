package main

import (
	"github.com/gin-gonic/gin"
	"shop/src/api/product"
	"shop/src/api/user"
)

func main() {
	r := gin.Default()
	catalog_handler.CatalogHandler(r)
	user_handler.UserHandler(r)
	r.Run(":8080") // Start the server on port 8080
}
