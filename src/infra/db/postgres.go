package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"shop/src/core/settings"
	"shop/src/infra/db/core"
	mig "shop/src/infra/db/migration"
)

// Init инициализирует глобальный объект Database
func InitPostgres() {
	config := settings.Config
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
