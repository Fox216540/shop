package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"shop/src/infra/db/core"
	mig "shop/src/infra/db/migration"
)

// Init инициализирует глобальный объект Database
func InitPostgres() {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	databaseName := os.Getenv("POSTGRES_DB")
	fmt.Println(user, password, host, port, databaseName)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, databaseName, port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	core.InitDatabase(db)
	core.GetDatabase()
	err = mig.Migration(db)
	if err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}
}

// Close завершает подключение к базе
func ClosePostgres() {
	database := core.GetDatabase()
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
