package order

import (
	"github.com/Fox216540/shop/apigateway-service/core/transport"
	"github.com/Fox216540/shop/apigateway-service/domain/order"
	"github.com/Fox216540/shop/apigateway-service/infra/client"
	common "github.com/Fox216540/shop/proto/common/gen"
	pb "github.com/Fox216540/shop/proto/order-service/gen/api"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPClient struct {
	conn     *client.GRPCClient
	pbClient pb.ApiServiceClient
}

func NewGRPClient(client *client.GRPCClient) *GRPClient {
	return &GRPClient{
		conn:     client,
		pbClient: pb.NewApiServiceClient(client.Conn()),
	}
}

func (c *GRPClient) mapOrder(o *pb.Order) (order.Order, error) {
	if o == nil {
		return order.Order{}, NewOrderResponseNilError(nil)
	}
	ID, err := uuid.Parse(o.Id)
	if err != nil {
		return order.Order{}, err
	}
	price := o.GetTotal()
	amount := ""
	currency := ""
	if price != nil {
		amount = price.GetAmount()
		currency = price.GetCurrency()
	}
	return order.Order{
		ID:       ID,
		OrderNum: o.OrderNumber,
		Total:    amount,
		Currency: currency,
		Status:   o.Status,
	}, nil
}

func (c *GRPClient) mapOrderWithItems(o *pb.OrderWithItems) (order.OrderWithItems, error) {
	if o == nil {
		return order.OrderWithItems{}, NewOrderResponseNilError(nil)
	}
	if o.Order == nil {
		return order.OrderWithItems{}, NewOrderPayloadNilError(nil)
	}
	items := make([]order.Item, 0, len(o.Items))
	for _, item := range o.Items {
		if item == nil || item.Product == nil {
			return order.OrderWithItems{}, NewOrderItemNilError(nil)
		}
		productID, err := uuid.Parse(item.Product.Id)
		if err != nil {
			return order.OrderWithItems{}, err
		}
		price := item.Product.GetPrice()
		amount := ""
		currency := ""
		if price != nil {
			amount = price.GetAmount()
			currency = price.GetCurrency()
		}
		items = append(items, order.Item{
			Product: order.Product{
				ID:       productID,
				Name:     item.Product.Name,
				Img:      item.Product.Img,
				Price:    amount,
				Currency: currency,
			},
			Quantity: item.Quantity,
		})
	}
	orderWithoutItems, err := c.mapOrder(o.Order)
	if err != nil {
		return order.OrderWithItems{}, err
	}
	return order.OrderWithItems{
		Order: orderWithoutItems,
		Items: items,
	}, nil
}

func (c *GRPClient) Create(userID uuid.UUID, items []order.ProductRequest) (order.Order, error) {
	ctx := c.conn.Context()
	ctx = metadata.AppendToOutgoingContext(ctx, transport.UserIDKey, userID.String())
	dtoReq := make([]*pb.ItemRequest, 0, len(items))
	for _, item := range items {
		dtoReq = append(dtoReq, &pb.ItemRequest{
			ProductId: item.ID.String(),
			Quantity:  item.Quantity,
			ExpectedPrice: &common.Money{
				Amount:   item.Price,
				Currency: item.Currency,
			},
		})
	}
	req := &pb.CreateOrderRequest{
		Items: dtoReq,
	}
	o, err := c.pbClient.CreateOrder(ctx, req)
	if err != nil {
		return order.Order{}, err
	}
	if o == nil {
		return order.Order{}, NewOrderResponseNilError(nil)
	}
	return c.mapOrder(o)
}

func (c *GRPClient) Get(userID uuid.UUID, orderID uuid.UUID) (orderWithItems order.OrderWithItems, err error) {
	ctx := c.conn.Context()
	ctx = metadata.AppendToOutgoingContext(ctx, transport.UserIDKey, userID.String())
	req := &pb.OrderId{
		Id: orderID.String(),
	}
	o, err := c.pbClient.GetOrderById(ctx, req)
	if err != nil {
		return order.OrderWithItems{}, err
	}
	return c.mapOrderWithItems(o)
}

func (c *GRPClient) GetOrders(userID uuid.UUID) ([]order.Order, error) {
	ctx := c.conn.Context()
	ctx = metadata.AppendToOutgoingContext(ctx, transport.UserIDKey, userID.String())
	req := &emptypb.Empty{}
	resp, err := c.pbClient.GetOrdersByUserId(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, NewOrdersResponseNilError(nil)
	}
	orders := make([]order.Order, 0, len(resp.Orders))
	for _, o := range resp.Orders {
		orderMap, err := c.mapOrder(o)
		if err != nil {
			return nil, err
		}
		orders = append(orders, orderMap)
	}
	return orders, nil
}

func (c *GRPClient) Delete(userID uuid.UUID, orderID uuid.UUID) (ID uuid.UUID, msg string, status string, err error) {
	ctx := c.conn.Context()
	ctx = metadata.AppendToOutgoingContext(ctx, transport.UserIDKey, userID.String())
	req := &pb.OrderId{
		Id: orderID.String(),
	}
	resp, err := c.pbClient.DeleteOrder(ctx, req)
	if err != nil {
		return uuid.Nil, "", "", err
	}
	if resp == nil || resp.OrderId == nil || resp.Message == nil {
		return uuid.Nil, "", "", NewDeleteOrderResponseNilError(nil)
	}
	OrderID, err := uuid.Parse(resp.OrderId.Id)
	if err != nil {
		return uuid.Nil, "", "", err
	}
	return OrderID, resp.Message.Message, resp.Status, nil
}
