package orderservice

import (
	"github.com/google/uuid"
	"shop/src/app/product"
	"shop/src/domain/idgenerator"
	"shop/src/domain/order"
)

type service struct {
	r     order.Repository
	ps    productservice.Service
	idGen idgenerator.Generator
}

func NewOrderService(
	r order.Repository,
	ps productservice.Service,
	idGen idgenerator.Generator,
) Service {
	return &service{r: r, ps: ps, idGen: idGen}
}

func (s *service) PlaceOrder(o order.Order) (order.Order, error) {
	var err error
	if o.ID == uuid.Nil {
		o.ID = uuid.New()
	}

	o = s.GenerateOrderNum(o) // Генерируем уникальный номер заказа

	o, err = s.CalculateTotalByIDs(o) // Вычисляем общую сумму заказа

	if err != nil {
		return o, err
	}

	o, err = s.r.Save(o)
	if err != nil {
		return o, err // Возвращаем ошибку, если не удалось сохранить заказ
	}
	return o, nil
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

func (s *service) CalculateTotalByIDs(o order.Order) (order.Order, error) {
	var total float64
	for _, item := range o.ProductItems {
		p, err := s.ps.ProductById(item.Product.ID)
		if err != nil {
			return order.Order{}, err
		}
		total += p.Price * float64(item.Quantity)
	}
	o.Total = total
	return o, nil
}

func (s *service) GenerateOrderNum(o order.Order) order.Order {
	orderNum, err := s.idGen.NewID()
	if err != nil {
		return order.Order{}
	}

	o.OrderNum = orderNum

	return o
}
