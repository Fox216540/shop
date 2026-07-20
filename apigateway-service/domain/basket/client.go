package basket

import "github.com/google/uuid"

type Client interface {
	GetBasket(userID uuid.UUID) (Basket, error)
	AddItemToBasket(userID, productID uuid.UUID, quantity uint64) (Item, error)
	DeleteBasket(userID uuid.UUID) error
	RemoveItemFromBasket(userID, productID uuid.UUID) (string, error)
}
