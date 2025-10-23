package migration

import (
	"github.com/Fox216540/shop/user-service/infra/user/models"
	"gorm.io/gorm"
)

func Migration(db *gorm.DB) error {
	//if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "pgcrypto"`).Error; err != nil {
	//	return fmt.Errorf("failed to create extension pgcrypto: %w", err)
	//}

	return db.AutoMigrate(
		&models.UserORM{})
}
