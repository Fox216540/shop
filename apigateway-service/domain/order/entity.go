package order

import "github.com/google/uuid"

type ProductRequest struct {
	ID       uuid.UUID
	Price    string
	Currency string
	Quantity uint64
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

type OrderWithItems struct {
	Order Order
	Items []Item
}

type Order struct {
	ID       uuid.UUID
	OrderNum string
	Total    string
	Currency string
	Status   string
}
