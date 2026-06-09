package models

import (
	"time"

	"github.com/Fox216540/shop/basket-service/domain/basket"
	"github.com/Fox216540/shop/basket-service/domain/money"
	"github.com/Fox216540/shop/basket-service/domain/productShort"
	"github.com/google/uuid"
)

type BasketItemORM struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID        uuid.UUID `gorm:"type:uuid;index:idx_basket_user_product,priority:1;not null"`
	ProductID     uuid.UUID `gorm:"type:uuid;index:idx_basket_user_product,priority:2;not null"`
	Name          string
	Img           string
	PriceAmount   string
	PriceCurrency string
	Quantity      uint64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewBasketItemORM(userID uuid.UUID, product productShort.ProductShort, quantity uint64) BasketItemORM {
	return BasketItemORM{
		ID:            uuid.New(),
		UserID:        userID,
		ProductID:     product.ID,
		Name:          product.Name,
		Img:           product.Img,
		PriceAmount:   product.Price.Amount,
		PriceCurrency: product.Price.Currency,
		Quantity:      quantity,
	}
}

func (m BasketItemORM) ToDomain() basket.ItemBasket {
	return basket.ItemBasket{
		Product: productShort.ProductShort{
			ID:   m.ProductID,
			Name: m.Name,
			Img:  m.Img,
			Price: money.Money{
				Amount:   m.PriceAmount,
				Currency: m.PriceCurrency,
			},
		},
		Quantity: m.Quantity,
	}
}
