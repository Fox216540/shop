package order

import "github.com/google/uuid"

type ProductRequest struct {
	ID       uuid.UUID
	Quantity uint64
}

type Product struct {
	ID    uuid.UUID
	Name  string
	Img   string
	Price float64
}

type Item struct {
	Product  *Product
	quantity uint64
}

type OrderWithItems struct {
	Order *Order
	Items []*Item
}

type Order struct {
	ID       uuid.UUID
	OrderNum string
	Total    float64
	Status   string
}
