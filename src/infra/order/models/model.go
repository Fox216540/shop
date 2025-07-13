package ordermodels

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"shop/src/domain/order"
	productORM "shop/src/infra/product/models"
	user "shop/src/infra/user/models"
)

type ProductItemORM struct {
	gorm.Model
	ID        int       `gorm:"primaryKey;autoIncrement"`
	OrderID   uuid.UUID `gorm:"type:uuid;not null;index"`
	ProductID uuid.UUID `gorm:"type:uuid;not null;index"`

	Product productORM.ProductORM `gorm:"references:ProductID"`

	Quantity int `gorm:"not null"`
}

type OrderORM struct {
	gorm.Model
	ID           int               `gorm:"primaryKey;autoIncrement"`
	OrderID      uuid.UUID         `gorm:"type:uuid;not null;uniqueIndex"`
	UserID       uuid.UUID         `gorm:"type:uuid;not null;index"`
	OrderNum     string            `gorm:"type:varchar(50);not null;uniqueIndex"`
	Status       string            `gorm:"type:varchar(50);not null"`
	ProductItems []*ProductItemORM `gorm:"foreignKey:OrderID;references:OrderID;constraint:OnDelete:CASCADE"`

	User user.UserORM `gorm:"references:UserID"`
}

func FromORM(orm OrderORM) order.Order {
	o := order.Order{
		ID:           orm.OrderID,
		OrderNum:     orm.OrderNum,
		Status:       orm.Status,
		UserId:       orm.UserID,
		ProductItems: make([]*order.Item, 0, len(orm.ProductItems)),
	}
	for _, item := range orm.ProductItems {
		o.ProductItems = append(o.ProductItems, &order.Item{
			Product:  productORM.FromORM(item.Product),
			Quantity: item.Quantity,
		})
	}

	return o
}
