package migration

import (
	"github.com/Fox216540/shop/basket-service/infra/basket/models"
	"gorm.io/gorm"
)

func Migration(db *gorm.DB) error {
	return db.AutoMigrate(&models.BasketItemORM{})
}
