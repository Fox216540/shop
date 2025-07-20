package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"shop/src/domain/order"
	productORM "shop/src/infra/product/models"
	user "shop/src/infra/user/models"
)

type OrderItemORM struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	OrderID   uuid.UUID `gorm:"type:uuid;not null;index"`
	ProductID uuid.UUID `gorm:"type:uuid;not null;index"`

	Product productORM.ProductORM `gorm:"references:ProductID"`

	Quantity int `gorm:"not null"`
}

func (OrderItemORM) TableName() string {
	return "order_product"
}

type OrderORM struct {
	gorm.Model
	ID         int             `gorm:"primaryKey;autoIncrement"`
	OrderID    uuid.UUID       `gorm:"type:uuid;not null;uniqueIndex"`
	UserID     uuid.UUID       `gorm:"type:uuid;not null;index"`
	OrderNum   string          `gorm:"type:varchar(50);not null;uniqueIndex"`
	Status     string          `gorm:"type:varchar(50);not null"`
	Total      float64         `gorm:"type:decimal(10,2);not null"` // Total order amount
	OrderItems []*OrderItemORM `gorm:"foreignKey:OrderID;references:OrderID;constraint:OnDelete:CASCADE"`

	User user.UserORM `gorm:"references:UserID"`
}

func (OrderORM) TableName() string {
	return "orders"
}

func FromORM(orm OrderORM) order.Order {
	o := order.Order{
		ID:         orm.OrderID,
		OrderNum:   orm.OrderNum,
		Status:     orm.Status,
		UserID:     orm.UserID,
		Total:      orm.Total,
		OrderItems: make([]*order.Item, 0, len(orm.OrderItems)),
	}
	for _, item := range orm.OrderItems {
		o.OrderItems = append(o.OrderItems, &order.Item{
			Product:  productORM.FromORM(item.Product),
			Quantity: item.Quantity,
		})
	}

	return o
}
