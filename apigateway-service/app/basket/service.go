package basket

import (
	"github.com/Fox216540/shop/apigateway-service/domain/basket"
	"github.com/google/uuid"
)

type service struct {
	basketClient basket.Client
}

func NewService(basketClient basket.Client) UseCase {
	return &service{
		basketClient: basketClient,
	}
}

func (s *service) GetBasket(userID uuid.UUID) (basket.Basket, error) {
	return s.basketClient.GetBasket(userID)
}

func (s *service) AddItemToBasket(userID, productID uuid.UUID, quantity uint64) (basket.Item, error) {
	return s.basketClient.AddItemToBasket(userID, productID, quantity)
}

func (s *service) DeleteBasket(userID uuid.UUID) error {
	return s.basketClient.DeleteBasket(userID)
}

func (s *service) RemoveItemFromBasket(userID, productID uuid.UUID) (string, error) {
	return s.basketClient.RemoveItemFromBasket(userID, productID)
}
