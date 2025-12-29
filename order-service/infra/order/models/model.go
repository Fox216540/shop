package models

import (
	"github.com/Fox216540/shop/order-service/domain/catalog"
	"github.com/Fox216540/shop/order-service/domain/order"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type OrderProductORM struct {
	ID              int             `gorm:"primaryKey;autoIncrement"`
	OrderID         uuid.UUID       `gorm:"type:uuid;not null;index"`
	ProductID       uuid.UUID       `gorm:"type:uuid;not null;index"`
	ProductName     string          `gorm:"type:varchar(255);not null"`
	ProductPrice    decimal.Decimal `gorm:"type:decimal(10,2);not null"`
	ProductCurrency string          `gorm:"type:char(3);not null"`
	ProductImg      string          `gorm:"type:varchar(255);not null"`
	Quantity        uint64          `gorm:"not null"`
}

func (OrderProductORM) TableName() string {
	return "order_product"
}

type OrderORM struct {
	ID         int               `gorm:"primaryKey;autoIncrement"`
	OrderID    uuid.UUID         `gorm:"type:uuid;not null;uniqueIndex"`
	UserID     uuid.UUID         `gorm:"type:uuid;not null;index"`
	OrderNum   string            `gorm:"type:varchar(50);not null;uniqueIndex"`
	Status     string            `gorm:"type:varchar(50);not null"`
	Total      decimal.Decimal   `gorm:"type:decimal(10,2);not null"` // Total order amount
	Currency   string            `gorm:"type:char(3);not null"`
	OrderItems []OrderProductORM `gorm:"foreignKey:OrderID;references:OrderID;constraint:OnDelete:CASCADE"`
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
		Currency:   orm.Currency,
		OrderItems: make([]order.Item, 0, len(orm.OrderItems)),
	}
	for _, item := range orm.OrderItems {
		o.OrderItems = append(o.OrderItems, order.Item{
			Product: catalog.Product{
				ID:       item.ProductID,
				Name:     item.ProductName,
				Img:      item.ProductImg,
				Price:    item.ProductPrice,
				Currency: item.ProductCurrency,
			},
			Quantity: item.Quantity,
		})
	}

	return o
}
