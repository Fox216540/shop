package orderservice

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"shop/src/domain/order"
	"time"
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

	o.OrderNum = s.GenerateOrderNum() // Генерируем уникальный номер заказа

	o.Total = s.CalculateTotal(o) // Вычисляем общую сумму заказа

	savedOrder, err := s.r.Save(o)
	if err != nil {
		return o, err // Возвращаем ошибку, если не удалось сохранить заказ
	}
	return savedOrder, nil
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

func (s *service) CalculateTotal(order order.Order) float64 {
	var total float64
	if len(order.ProductItems) == 0 {
		return 0 // Возвращаем 0, если нет товаров в заказе
	}
	for _, item := range order.ProductItems {
		total += item.Product.Price * float64(item.Quantity)
	}
	return total
}

func (s *service) GenerateOrderNum() string {
	// Генерируем уникальный номер заказа на основе текущей даты и случайного числа
	datetime := time.Now().Format("20060102")
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	var orderNum string

	for {
		randomPart := r.Intn(10000)                                 // Генерируем новый случайный номер
		orderNum = fmt.Sprintf("ORD-%s-%04d", datetime, randomPart) // Форматируем новый номер заказа
		err := s.r.CheckOrderNum(orderNum)                          // Проверяем, существует ли уже такой номер заказа
		if err == nil {
			break // Если номер заказа уникален, выходим из цикла
		}
	}

	return orderNum
}
