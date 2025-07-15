package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	category "shop/src/api/category"
	product "shop/src/api/product"
	"shop/src/infra/db"
)

func main() {
	err := godotenv.Load() // по умолчанию ищет файл .env в текущей папке
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	db.Init()
	defer db.Close()
	r := gin.Default()
	product.ProductHandler(r)
	category.CategoryHandler(r)
	r.Run(":8080") // Start the server on port 8080
}
