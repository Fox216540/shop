package basket

import (
	"github.com/Fox216540/shop/basket-service/domain/money"
	"github.com/Fox216540/shop/basket-service/domain/productShort"
)

type ItemBasket struct {
	Product  productShort.ProductShort
	Quantity uint64
}

type Basket struct {
	Products []ItemBasket
	Total    money.Money
}
