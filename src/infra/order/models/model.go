package ordermodels

import (
	"github.com/google/uuid"
	"shop/src/domain/order"
	"shop/src/domain/product"
)

type ProductItemORM struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	OrderID   uuid.UUID `gorm:"type:uuid;not null;index"`
	ProductID uuid.UUID `gorm:"type:uuid;not null;index"`

	Quantity int `gorm:"not null"`
}

type OrderORM struct {
	ID           int               `gorm:"primaryKey;autoIncrement"`
	OrderID      uuid.UUID         `gorm:"type:uuid;not null;uniqueIndex"`
	OrderNum     string            `gorm:"type:varchar(50);not null;uniqueIndex"`
	Status       string            `gorm:"type:varchar(50);not null"`
	ProductItems []*ProductItemORM `gorm:"foreignKey:OrderID;references:OrderID;constraint:OnDelete:CASCADE"`
}

func FromORM(orm OrderORM, productMap map[uuid.UUID]product.Product) order.Order {
	o := order.Order{
		ID:           orm.OrderID,
		OrderNum:     orm.OrderNum,
		Status:       orm.Status,
		ProductItems: make([]*order.Item, 0, len(orm.ProductItems)),
	}
	for _, item := range orm.ProductItems {
		p := productMap[item.ProductID]
		o.ProductItems = append(o.ProductItems, &order.Item{
			Product:  p,
			Quantity: item.Quantity,
		})
	}

	return o
}
