package order

import "github.com/google/uuid"

type Service interface {
	PlaceOrder(o Order) (Order, error)
	CancelOrder(ID, userID uuid.UUID) (uuid.UUID, error)
	GetByID(ID uuid.UUID) (Order, error)
	OrdersByUserID(userID uuid.UUID) ([]Order, error) //Для domain User
	CalculateTotal(o Order) float64
	GenerateOrderNum() string
}
