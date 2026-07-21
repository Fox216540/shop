package basket

import (
	"github.com/Fox216540/shop/apigateway-service/domain/basket"
	"github.com/google/uuid"
)

type UseCase interface {
	GetBasket(userID uuid.UUID) (basket.Basket, error)
	AddItemToBasket(userID, productID uuid.UUID, quantity uint64) (basket.Item, error)
	DeleteBasket(userID uuid.UUID) error
	RemoveItemFromBasket(userID, productID uuid.UUID) (string, error)
}
