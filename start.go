package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"shop/src/api/category"
	"shop/src/api/product"
	"shop/src/api/user"
	"shop/src/infra/db"
)

func main() {
	err := godotenv.Load() // по умолчанию ищет файл .env в текущей папке
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
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
