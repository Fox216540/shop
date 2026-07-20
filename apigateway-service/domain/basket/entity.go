package basket

import "github.com/google/uuid"

type Money struct {
	Amount   string
	Currency string
}

type Product struct {
	ID       uuid.UUID
	Name     string
	Img      string
	Price    string
	Currency string
}

type Item struct {
	Product  Product
	Quantity uint64
}

type Basket struct {
	Items []Item
	Total Money
}
