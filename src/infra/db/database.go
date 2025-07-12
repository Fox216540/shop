package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type Database struct {
	DB *gorm.DB
}

// Глобальный объект Database
var database *Database

// Init инициализирует глобальный объект Database
func Init() {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	database = &Database{DB: db}
}

// Close завершает подключение к базе
func Close() {
	if database == nil || database.DB == nil {
		return
	}
	sqlDB, err := database.DB.DB()
	if err != nil {
		log.Printf("failed to get raw DB: %v", err)
		return
	}
	sqlDB.Close()
}
