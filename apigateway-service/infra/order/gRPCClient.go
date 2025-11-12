package order

import (
	"github.com/Fox216540/shop/apigateway-service/domain/order"
	"github.com/Fox216540/shop/apigateway-service/infra/client"
	pb "github.com/Fox216540/shop/proto/order-service/gen/api"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPClient struct {
	conn     *client.GRPCClient
	pbClient pb.ApiServiceClient
}

func (c *GRPClient) mapOrder(o *pb.Order) (order.Order, error) {
	ID, err := uuid.Parse(o.Id)
	if err != nil {
		return order.Order{}, err
	}
	return order.Order{
		ID:       ID,
		OrderNum: o.OrderNum,
		Total:    o.Total,
		Status:   o.Status,
	}, nil
}

func (c *GRPClient) CreateOrder(userID uuid.UUID, items []order.ProductRequest) (order.Order, error) {
	ctx := c.conn.Context()
	ctx = metadata.AppendToOutgoingContext(ctx, "user_id", userID.String())
	dtoReq := make([]*pb.ItemRequest, 0, len(items))
	for _, item := range items {
		dtoReq = append(dtoReq, &pb.ItemRequest{
			ProductId: item.ID.String(),
			Quantity:  item.Quantity,
		})
	}
	req := &pb.CreateOrderRequest{
		Items: dtoReq,
	}
	o, err := c.pbClient.CreateOrder(ctx, req)
	if err != nil {
		return order.Order{}, err
	}
	return c.mapOrder(o)
}

func (c *GRPClient) GetOrder(userID uuid.UUID, orderID uuid.UUID) (order order.Order, err error) {
	ctx := c.conn.Context()
	ctx = metadata.AppendToOutgoingContext(ctx, "user_id", userID.String())
	req := &pb.OrderId{
		Id: orderID.String(),
	}
	o, err := c.pbClient.GetOrderById(ctx, req)
	return c.mapOrder(o)
}

func (c *GRPClient) GetOrders(userID uuid.UUID) ([]order.Order, error) {
	ctx := c.conn.Context()
	ctx = metadata.AppendToOutgoingContext(ctx, "user_id", userID.String())
	req := &emptypb.Empty{}
	resp, err := c.pbClient.GetOrdersByUserId(ctx, req)
	if err != nil {
		return nil, err
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

func (c *GRPClient) DeleteOrder(userID uuid.UUID, orderID uuid.UUID) (ID uuid.UUID, msg string, status string, err error) {
	ctx := c.conn.Context()
	ctx = metadata.AppendToOutgoingContext(ctx, "user_id", userID.String())
	req := &pb.OrderId{
		Id: orderID.String(),
	}
	resp, err := c.pbClient.DeleteOrder(ctx, req)
	if err != nil {
		return uuid.Nil, "", "", err
	}
	OrderID, err := uuid.Parse(resp.OrderId.Id)
	if err != nil {
		return uuid.Nil, "", "", err
	}
	return OrderID, resp.Message.Message, resp.Status, nil
}
