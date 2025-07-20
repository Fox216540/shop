package order

import (
	"fmt"
	"github.com/google/uuid"
	"shop/src/app/product"
	"shop/src/domain/idgenerator"
	"shop/src/domain/order"
	p "shop/src/domain/product"
)

type service struct {
	r     order.Repository
	ps    product.UseCase
	idGen idgenerator.Generator
}

func NewOrderService(
	r order.Repository,
	ps product.UseCase,
	idGen idgenerator.Generator,
) UseCase {
	return &service{r: r, ps: ps, idGen: idGen}
}

func (s *service) Place(userID uuid.UUID, productItems []*order.Item) (order.Order, error) {
	var err error
	o := order.Order{
		ID:         uuid.New(),
		UserID:     userID,
		OrderItems: productItems,
	}

	productsIDs := make([]uuid.UUID, 0, len(o.OrderItems))
	for _, item := range o.OrderItems {
		productsIDs = append(productsIDs, item.Product.ID)
	}

	products, err := s.ps.ProductsByIDs(productsIDs)
	if err != nil {
		return o, err
	}

	productMap := make(map[uuid.UUID]p.Product)
	for _, prod := range products {
		productMap[prod.ID] = prod
	}

	for _, item := range o.OrderItems {
		if prod, ok := productMap[item.Product.ID]; ok {
			item.Product = prod
		} else {
			return o, fmt.Errorf("product not found: %s", item.Product.ID)
		}
	}

	o, err = s.generateOrderNum(o) // Генерируем уникальный номер заказа
	if err != nil {
		return o, err
	}

	o, err = s.calculateOrderTotal(o) // Вычисляем общую сумму заказа
	if err != nil {
		return o, err
	}

	o, err = s.r.Save(o)
	if err != nil {
		return o, err // Возвращаем ошибку, если не удалось сохранить заказ
	}

	return o, nil
}

func (s *service) Cancel(ID, userID uuid.UUID) (uuid.UUID, error) {
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

func (s *service) calculateOrderTotal(o order.Order) (order.Order, error) {
	var total float64
	for _, item := range o.OrderItems {
		total += item.Product.Price * float64(item.Quantity)
	}
	o.Total = total
	return o, nil
}

func (s *service) generateOrderNum(o order.Order) (order.Order, error) {
	orderNum, err := s.idGen.NewID()
	if err != nil {
		return order.Order{}, err
	}

	o.OrderNum = orderNum

	return o, nil
}
