package app

import (
	"github.com/Fox216540/shop/basket-service/domain/basket"
	"github.com/google/uuid"
)

type UseCase interface {
	GetBasket(userID uuid.UUID) (basket.Basket, error)
	AddItemToBasket(userID, productID uuid.UUID, quantity uint64) (basket.ItemBasket, error)
	DeleteBasket(userID uuid.UUID) error
	RemoveItemFromBasket(userID, productID uuid.UUID) error
}
