package orderservice

import (
	"github.com/google/uuid"
	"shop/src/domain/order"
)

type service struct {
	r order.Repository
}

func NewOrderService(r order.Repository) order.Service {
	return &service{r: r}
}

func (s *service) PlaceOrder(o order.Order) (order.Order, error) {
	if o.ID == uuid.Nil {
		o.ID = uuid.New()
	}
	orderSave, err := s.r.Save(o)
	if err != nil {
		return o, err // Возвращаем ошибку, если не удалось сохранить заказ
	}
	return orderSave, nil
}

func (s *service) CancelOrder(ID, userID uuid.UUID) (uuid.UUID, error) {
	if err := s.r.Remove(ID, userID); err != nil {
		return ID, err // Возвращаем ошибку, если не удалось удалить заказ
	}
	return ID, nil
}

func (s *service) GetByID(id uuid.UUID) (order.Order, error) {
	o, err := s.r.GetByID(id)
	if err != nil {
		return o, err // Возвращаем ошибку, если не удалось найти заказ
	}
	return o, nil
}

func (s *service) OrdersByUserID(userID uuid.UUID) ([]order.Order, error) {
	orders, err := s.r.OrdersByUserID(userID)
	if err != nil {
		return nil, err // Возвращаем ошибку, если не удалось найти заказы пользователя
	}
	return orders, nil
}
