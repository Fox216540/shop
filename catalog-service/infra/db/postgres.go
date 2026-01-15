package db

import (
	"context"
	"fmt"
	"github.com/Fox216540/shop/catalog-service/core/logger"
	"github.com/Fox216540/shop/catalog-service/infra/config"
	"github.com/Fox216540/shop/catalog-service/infra/db/core"
	"github.com/Fox216540/shop/catalog-service/infra/db/errors"
	mig "github.com/Fox216540/shop/catalog-service/infra/db/migration"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitPostgres инициализирует глобальный объект Database
func InitPostgres(
	ctx context.Context,
	log logger.Logger,
	conf *config.Config,
) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		conf.PostgresHost,
		conf.PostgresUser,
		conf.PostgresPassword,
		conf.PostgresDatabase,
		conf.PostgresPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error(ctx, errors.NewConnectionError(err))
		return nil, errors.NewConnectionError(err)
	}

	if err := mig.Migration(db); err != nil {
		log.Error(ctx, errors.NewMigrationError(err))
		return nil, errors.NewMigrationError(err)
	}

	return db, nil
}

// ClosePostgres завершает подключение к базе
func ClosePostgres(
	ctx context.Context,
	log logger.Logger,
	database *core.Database,
) {
	if database == nil || database.DB == nil {
		return
	}

	sqlDB, err := database.DB.DB()
	if err != nil {
		log.Error(ctx, errors.NewGetRawDBError(err))
		return
	}

	if err = sqlDB.Close(); err != nil {
		log.Error(ctx, errors.NewCloseError(err))
		return
	}
}
