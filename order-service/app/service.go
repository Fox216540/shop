package order

import (
	"github.com/Fox216540/shop/order-service/app/dto"
	"github.com/Fox216540/shop/order-service/app/mapError"
	c "github.com/Fox216540/shop/order-service/domain/catalog"
	"github.com/Fox216540/shop/order-service/domain/idgenerator"
	"github.com/Fox216540/shop/order-service/domain/order"
	"github.com/google/uuid"
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
	productsOrder := make([]uuid.UUID, len(items))
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

	for _, item := range items {
		if prod, ok := productMapDb[item.Product.ID]; ok {
			item.Product = prod
		} else {
			return order.NewProductOfOrderNotFoundError(nil)
		}
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
	var total float64
	for _, item := range o.OrderItems {
		total += item.Product.Price * float64(item.Quantity)
	}
	o.Total = total
	return o, nil
}

func (s *service) createOrderItems(orderItems []dto.OrderItems) ([]order.Item, error) {
	oI := make([]order.Item, len(orderItems))
	for _, item := range orderItems {
		id, err := uuid.Parse(item.ProductID)
		if err != nil {
			return nil, err
		}
		oI = append(oI, order.Item{
			Product: c.Product{
				ID: id,
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
