package db

import (
	"fmt"
	"github.com/Fox216540/shop/order-service/infra/config"
	"github.com/Fox216540/shop/order-service/infra/db/core"
	mig "github.com/Fox216540/shop/order-service/infra/db/migration"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

// Init инициализирует глобальный объект Database
func InitPostgres(config *config.Config) *gorm.DB {
	user := config.PostgresUser
	password := config.PostgresPassword
	host := config.PostgresHost
	port := config.PostgresPort
	databaseName := config.PostgresDatabase
	fmt.Println(user, password, host, port, databaseName)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, databaseName, port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	err = mig.Migration(db)
	if err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}
	return db
}

// Close завершает подключение к базе
func ClosePostgres(database *core.Database) {
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
