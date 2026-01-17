package productShort

import (
	"github.com/Fox216540/shop/basket-service/domain/money"
	"github.com/google/uuid"
)

type ProductShort struct {
	ID    uuid.UUID
	Name  string
	Img   string
	Price money.Money
}
