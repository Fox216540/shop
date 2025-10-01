package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"shop/src/api/category"
	"shop/src/api/product"
	"shop/src/api/user"
	"shop/src/core/logger"
	"shop/src/infra/db"
)

func main() {
	db.InitPostgres()
	db.InitRedis()
	defer db.ClosePostgres()
	defer db.CloseRedis()
	if err := os.MkdirAll("logs", 0755); err != nil {
		log.Fatalf("Failed to create log dir: %v", err)
	}

	file, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()

	// Инициализируем глобальный логгер
	logger.InitLogger(file)

	fileServerError, err := os.OpenFile("logs/server-error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open error log file: %v", err)
	}

	defer fileServerError.Close()

	logger.InitErrorLogger(fileServerError)

	r := gin.Default()
	product.Handler(r)
	category.Handler(r)
	user.Handler(r)
	r.Run(":8080")
}
