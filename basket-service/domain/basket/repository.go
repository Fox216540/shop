package basket

import (
	"github.com/Fox216540/shop/basket-service/domain/productShort"
	"github.com/google/uuid"
)

type Repository interface {
	GetBasket(userID uuid.UUID) (Basket, error)
	AddItemToBasket(userID uuid.UUID, product productShort.ProductShort, quantity uint64) (ItemBasket, error)
	DeleteBasket(userID uuid.UUID) error
	RemoveItemFromBasket(userID, productID uuid.UUID) error
}
