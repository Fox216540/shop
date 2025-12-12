package db

import (
	"fmt"
	"github.com/Fox216540/shop/catalog-service/infra/config"
	"github.com/Fox216540/shop/catalog-service/infra/db/core"
	mig "github.com/Fox216540/shop/catalog-service/infra/db/migration"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

// InitPostgres инициализирует глобальный объект Database
func InitPostgres(conf *config.Config) *gorm.DB {
	user := conf.PostgresUser
	password := conf.PostgresPassword
	host := conf.PostgresHost
	port := conf.PostgresPort
	databaseName := conf.PostgresDatabase
	fmt.Println(user, password, host, port, databaseName)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
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

// ClosePostgres завершает подключение к базе
func ClosePostgres(database *core.Database) {
	if database == nil || database.DB == nil {
		return
	}
	sqlDB, err := database.DB.DB()
	if err != nil {
		log.Printf("failed to get raw DB: %v", err)
		return
	}
	if err = sqlDB.Close(); err != nil {
		return
	}
}
