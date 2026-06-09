package app

import (
	"github.com/Fox216540/shop/basket-service/domain/basket"
	"github.com/Fox216540/shop/basket-service/domain/money"
	"github.com/Fox216540/shop/basket-service/domain/productShort"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type service struct {
	r            basket.Repository
	productShort productShort.Client
}

func NewService(r basket.Repository, productShortClient productShort.Client) UseCase {
	return &service{
		r:            r,
		productShort: productShortClient,
	}
}

func (s *service) GetBasket(userID uuid.UUID) (basket.Basket, error) {
	basketDomain, err := s.r.GetBasket(userID)
	if err != nil {
		return basket.Basket{}, err
	}

	total, err := s.calculateBasketTotal(basketDomain.Products)
	if err != nil {
		return basket.Basket{}, err
	}

	basketDomain.Total = total
	return basketDomain, nil
}

func (s *service) AddItemToBasket(userID, productID uuid.UUID, quantity uint64) (basket.ItemBasket, error) {
	if quantity == 0 {
		return basket.ItemBasket{}, basket.NewInvalidQuantityError(nil)
	}

	product, err := s.productShort.GetProductByID(productID)
	if err != nil {
		return basket.ItemBasket{}, err
	}

	return s.r.AddItemToBasket(userID, product, quantity)
}

func (s *service) DeleteBasket(userID uuid.UUID) error {
	return s.r.DeleteBasket(userID)
}

func (s *service) RemoveItemFromBasket(userID, productID uuid.UUID) error {
	return s.r.RemoveItemFromBasket(userID, productID)
}

func (s *service) calculateBasketTotal(items []basket.ItemBasket) (money.Money, error) {
	total := decimal.Zero
	currency := ""

	for _, item := range items {
		price, err := decimal.NewFromString(item.Product.Price.Amount)
		if err != nil {
			return money.Money{}, NewBasketInvalidItemAmountError(err)
		}

		if currency == "" {
			currency = item.Product.Price.Currency
		} else if currency != item.Product.Price.Currency {
			return money.Money{}, NewBasketContainsMultipleCurrenciesError(nil)
		}

		total = total.Add(price.Mul(decimal.NewFromInt(int64(item.Quantity))))
	}

	return money.Money{
		Amount:   total.String(),
		Currency: currency,
	}, nil
}
