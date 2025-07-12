package orderrepository

import (
	"github.com/google/uuid"
	"shop/src/domain/order"
	product "shop/src/infra/product"
)

type ProductItemORM struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	OrderID   uuid.UUID `gorm:"type:uuid;not null;index"`
	ProductID uuid.UUID `gorm:"type:uuid;not null;index"`
	Quantity  int       `gorm:"not null"`

	Product product.ProductORM `gorm:"foreignKey:ProductID;references:ProductID;constraint:OnDelete:CASCADE"`
}

type OrderORM struct {
	ID           int              `gorm:"primaryKey;autoIncrement"`
	OrderID      uuid.UUID        `gorm:"type:uuid;not null;uniqueIndex"`
	OrderNum     string           `gorm:"type:varchar(50);not null;uniqueIndex"`
	Status       string           `gorm:"type:varchar(50);not null"`
	ProductItems []ProductItemORM `gorm:"foreignKey:OrderID;references:OrderID;constraint:OnDelete:CASCADE"`
}

func fromORM(ormOrder OrderORM) order.Order {
	o := order.Order{
		ID:           ormOrder.OrderID,
		OrderNum:     ormOrder.OrderNum,
		Status:       ormOrder.Status,
		ProductItems: make([]order.ProductItem, len(ormOrder.ProductItems)),
	}

	for i, item := range ormOrder.ProductItems {
		o.ProductItems[i] = order.ProductItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		}
	}

	return o
}
