package order

import (
	"fmt"
	"github.com/Fox216540/shop/order-service/app/dto"
	"github.com/Fox216540/shop/order-service/app/mapError"
	c "github.com/Fox216540/shop/order-service/domain/catalog"
	"github.com/Fox216540/shop/order-service/domain/idgenerator"
	"github.com/Fox216540/shop/order-service/domain/order"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type service struct {
	r             order.Repository
	idGen         idgenerator.Generator
	catalogClient c.Client
}

func NewOrderService(
	r order.Repository,
	idGen idgenerator.Generator,
	catalogClient c.Client,
) UseCase {
	return &service{r: r, idGen: idGen, catalogClient: catalogClient}
}

func (s *service) createNewOrder(userID uuid.UUID, items []order.Item) order.Order {
	return order.Order{
		ID:         uuid.New(),
		UserID:     userID,
		OrderItems: items,
	}
}

func (s *service) enrichProducts(items []order.Item) error {
	productsOrder := make([]uuid.UUID, 0, len(items))

	for _, item := range items {
		productsOrder = append(productsOrder, item.Product.ID)
	}

	productsDb, err := s.catalogClient.GetProductsByIDs(productsOrder)
	if err != nil {
		return err
	}

	productMapDb := make(map[uuid.UUID]c.Product, len(productsDb))
	for _, prod := range productsDb {
		productMapDb[prod.ID] = prod
	}

	for i := range items {
		prod, ok := productMapDb[items[i].Product.ID]
		if !ok {
			return order.NewProductOfOrderNotFoundError(nil)
		}

		//TODO: Поменять ошибку
		if !prod.Price.Equal(items[i].Product.Price) {
			return fmt.Errorf("product %v does not have the same price", items[i].Product.ID)
		}

		items[i].Product = prod
	}

	return nil
}

func (s *service) generateOrderNum(o order.Order) (order.Order, error) {
	orderNum, err := s.idGen.NewID()
	if err != nil {
		return order.Order{}, err
	}
	if err = s.r.CheckOrderNum(orderNum); err != nil {
		return order.Order{}, err
	}
	o.OrderNum = orderNum

	return o, nil
}

func (s *service) calculateOrderTotal(o order.Order) (order.Order, error) {
	total := decimal.Zero

	for _, item := range o.OrderItems {
		lineTotal := item.Product.Price.
			Mul(decimal.NewFromInt(int64(item.Quantity)))

		total = total.Add(lineTotal)
	}

	o.Total = total
	return o, nil
}

func (s *service) createOrderItems(orderItems []dto.OrderItems) ([]order.Item, error) {
	oI := make([]order.Item, 0, len(orderItems))

	for _, item := range orderItems {
		id, err := uuid.Parse(item.ProductID)
		if err != nil {
			return nil, err
		}

		price, err := decimal.NewFromString(item.Value)
		if err != nil {
			//TODO: Поменять ошибку 404
			return nil, fmt.Errorf("invalid price %q: %w", item.Value, err)
		}

		if price.LessThan(decimal.Zero) {
			//TODO: Поменять ошибку 404
			return nil, fmt.Errorf("price must be >= 0")
		}

		oI = append(oI, order.Item{
			Product: c.Product{
				ID:       id,
				Price:    price,
				Currency: item.Currency,
			},
			Quantity: item.Quantity,
		})
	}

	return oI, nil
}

func (s *service) PlaceOrder(userID uuid.UUID, orderItems []dto.OrderItems) (order.Order, error) {
	oI, err := s.createOrderItems(orderItems)
	if err != nil {
		return order.Order{}, mapError.MapError(err, NewInvalidPlace(err))
	}

	o := s.createNewOrder(userID, oI)

	if err = s.enrichProducts(o.OrderItems); err != nil {
		return o, mapError.MapError(err, NewInvalidPlace(err))
	}

	o, err = s.generateOrderNum(o) // Генерируем уникальный номер заказа
	if err != nil {
		return o, mapError.MapError(err, NewInvalidPlace(err))
	}

	o, err = s.calculateOrderTotal(o) // Вычисляем общую сумму заказа
	if err != nil {
		return o, mapError.MapError(err, NewInvalidPlace(err))
	}

	o, err = s.r.Save(o)
	if err != nil {
		return o, mapError.MapError(err, NewInvalidPlace(err)) // Возвращаем ошибку, если не удалось сохранить заказ
	}

	return o, nil
}

func (s *service) DeleteOrder(ID, userID uuid.UUID) error {
	if err := s.r.Remove(ID, userID); err != nil {
		return mapError.MapError(err, NewInvalidCancel(err)) // Возвращаем ошибку, если не удалось удалить заказ
	}
	return nil
}

func (s *service) GetOrderByIDAndUserID(id, userID uuid.UUID) (order.Order, error) {
	o, err := s.r.GetByIDAndUserID(id, userID)
	if err != nil {
		return o, mapError.MapError(err, NewInvalidGetByID(err)) // Возвращаем ошибку, если не удалось найти заказ
	}
	return o, nil
}

func (s *service) GetOrdersByUserID(userID uuid.UUID) ([]order.Order, error) {
	orders, err := s.r.GetOrdersByUserID(userID)
	if err != nil {
		return nil, mapError.MapError(err, NewInvalidGetOrdersByUserID(err)) // Возвращаем ошибку, если не удалось найти заказы пользователя
	}
	return orders, nil
}
