package main

import (
	"github.com/gin-gonic/gin"
	"shop/src/api/category"
	"shop/src/api/product"
	"shop/src/api/user"
	"shop/src/infra/db"
)

func main() {
	db.InitPostgres()
	db.InitRedis()
	defer db.ClosePostgres()
	defer db.CloseRedis()
	r := gin.Default()
	product.Handler(r)
	category.Handler(r)
	user.Handler(r)
	r.Run(":8080")
}
