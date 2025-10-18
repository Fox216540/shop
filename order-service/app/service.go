package order

import (
	"errors"
	"fmt"
	"github.com/Fox216540/shop/order-service/app/dto"
	"github.com/Fox216540/shop/order-service/core/exception"
	"github.com/Fox216540/shop/order-service/domain/idgenerator"
	"github.com/Fox216540/shop/order-service/domain/order"
	p "github.com/Fox216540/shop/order-service/domain/product"
	"github.com/google/uuid"
)

type service struct {
	r             order.Repository
	idGen         idgenerator.Generator
	productClient p.Client
}

func NewOrderService(
	r order.Repository,
	idGen idgenerator.Generator,
	productClient p.Client,
) UseCase {
	return &service{r: r, idGen: idGen, productClient: productClient}
}

func (s *service) mapError(err, appServerError error) error {
	var appError *exception.ServerError
	var domainError *exception.DomainError
	if errors.As(err, &domainError) {
		return err
	}
	if errors.As(err, &appError) {
		return err
	}
	return appServerError
}

func (s *service) createNewOrder(userID uuid.UUID, items []*order.Item) order.Order {
	return order.Order{
		ID:         uuid.New(),
		UserID:     userID,
		OrderItems: items,
	}
}

func (s *service) enrichProducts(items []*order.Item) error {
	productsOrder := make([]uuid.UUID, len(items))
	for _, item := range items {
		productsOrder = append(productsOrder, item.Product.ID)
	}

	productsDb, err := s.productClient.GetProductsByIDs(productsOrder)
	if err != nil {
		return err
	}

	productMapDb := make(map[uuid.UUID]*p.Product, len(productsDb))
	for _, prod := range productsDb {
		productMapDb[prod.ID] = &prod
	}

	for _, item := range items {
		if prod, ok := productMapDb[item.Product.ID]; ok {
			item.Product = prod
		} else {
			// TODO: разобраться какую ошибку возвращать
			return fmt.Errorf("product not found: %s", item.Product.ID)
		}
	}

	return nil
}

func (s *service) generateOrderNum(o order.Order) (order.Order, error) {
	orderNum, err := s.idGen.NewID()
	if err != nil {
		return order.Order{}, err
	}
	// TODO: Проверить не занят ли номер заказа
	o.OrderNum = orderNum

	return o, nil
}

func (s *service) calculateOrderTotal(o order.Order) (order.Order, error) {
	var total float64
	for _, item := range o.OrderItems {
		total += item.Product.Price * float64(item.Quantity)
	}
	o.Total = total
	return o, nil
}

func (s *service) createOrderItems(orderItems []dto.OrderItems) ([]*order.Item, error) {
	oI := make([]*order.Item, len(orderItems))
	for _, item := range orderItems {
		id, err := uuid.Parse(item.ProductID)
		if err != nil {
			return nil, s.mapError(err, NewInvalidPlace(err))
		}
		oI = append(oI, &order.Item{
			Product: &p.Product{
				ID: id,
			},
			Quantity: item.Quantity,
		})
	}
	return oI, nil
}

func (s *service) Place(userID uuid.UUID, orderItems []dto.OrderItems) (order.Order, error) {
	oI, err := s.createOrderItems(orderItems)
	if err != nil {
		return order.Order{}, s.mapError(err, NewInvalidPlace(err))
	}

	o := s.createNewOrder(userID, oI)

	if err = s.enrichProducts(o.OrderItems); err != nil {
		return o, s.mapError(err, NewInvalidPlace(err))
	}

	o, err = s.generateOrderNum(o) // Генерируем уникальный номер заказа
	if err != nil {
		return o, s.mapError(err, NewInvalidPlace(err))
	}

	o, err = s.calculateOrderTotal(o) // Вычисляем общую сумму заказа
	if err != nil {
		return o, s.mapError(err, NewInvalidPlace(err))
	}

	o, err = s.r.Save(o)
	if err != nil {
		return o, s.mapError(err, NewInvalidPlace(err)) // Возвращаем ошибку, если не удалось сохранить заказ
	}

	return o, nil
}

func (s *service) Cancel(ID, userID uuid.UUID) (uuid.UUID, error) {
	if err := s.r.Remove(ID, userID); err != nil {
		return ID, s.mapError(err, NewInvalidCancel(err)) // Возвращаем ошибку, если не удалось удалить заказ
	}
	return ID, nil
}

func (s *service) GetByID(id uuid.UUID) (order.Order, error) {
	o, err := s.r.GetByID(id)
	if err != nil {
		return o, s.mapError(err, NewInvalidGetByID(err)) // Возвращаем ошибку, если не удалось найти заказ
	}
	return o, nil
}

func (s *service) GetOrdersByUserID(userID uuid.UUID) ([]order.Order, error) {
	orders, err := s.r.GetOrdersByUserID(userID)
	if err != nil {
		return nil, s.mapError(err, NewInvalidGetOrdersByUserID(err)) // Возвращаем ошибку, если не удалось найти заказы пользователя
	}
	return orders, nil
}
