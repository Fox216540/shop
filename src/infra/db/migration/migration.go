package migration

import (
	"gorm.io/gorm"
	order "shop/src/infra/order/models"
	product "shop/src/infra/product/models"
	user "shop/src/infra/user/models"
)

func Migration(db *gorm.DB) error {
	return db.AutoMigrate(
		&product.ProductORM{},
		&order.ProductItemORM{},
		&order.OrderORM{},
		&user.UserORM{})
}
