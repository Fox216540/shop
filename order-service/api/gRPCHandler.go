package api

import (
	"context"
	"errors"
	order "github.com/Fox216540/shop/order-service/app"
	"github.com/Fox216540/shop/order-service/app/dto"
	orderDomain "github.com/Fox216540/shop/order-service/domain/order"
	types "github.com/Fox216540/shop/proto/common/gen"
	pbApi "github.com/Fox216540/shop/proto/order-service/gen/api"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCHandler struct {
	orderUC order.UseCase
	pb      pbApi.UnimplementedApiServiceServer
	mapper  *ErrorMapper
}

func (h *GRPCHandler) mapOrderWithItems(
	order orderDomain.Order,
) *pbApi.OrderWithItems {
	items := make([]*pbApi.Item, 0, len(order.OrderItems))
	for _, item := range order.OrderItems {
		items = append(items, &pbApi.Item{
			Product: &types.Product{
				Id:    item.Product.ID.String(),
				Name:  item.Product.Name,
				Img:   item.Product.Img,
				Price: item.Product.Price,
			},
			Quantity: item.Quantity,
		})
	}
	return &pbApi.OrderWithItems{
		Order: &pbApi.Order{
			Id:       order.ID.String(),
			OrderNum: order.OrderNum,
			Total:    order.Total,
			Status:   order.Status,
		},
		Items: items,
	}
}

func (h *GRPCHandler) mapOrderToResponse(
	orders []orderDomain.Order,
) []*pbApi.Order {
	ordersForList := make([]*pbApi.Order, 0, len(orders))
	for _, o := range orders {
		ordersForList = append(ordersForList, &pbApi.Order{
			Id:       o.ID.String(),
			OrderNum: o.OrderNum,
			Total:    o.Total,
			Status:   o.Status,
		})
	}
	return ordersForList
}

func (h *GRPCHandler) getUserIDFromMetadata(
	ctx context.Context,
) (uuid.UUID, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return uuid.Nil, errors.New("metadata not found")
	}
	userID, err := uuid.Parse(md.Get("user_id")[0])
	if err != nil {
		return uuid.Nil, err
	}
	return userID, nil
}

func (h *GRPCHandler) CreteOrder(
	ctx context.Context,
	req *pbApi.CreateOrderRequest,
) (*pbApi.OrderWithItems, error) {
	userID, err := h.getUserIDFromMetadata(ctx)
	if err != nil {
		return nil, h.mapper.MapError(err)
	}

	productsIDs := make([]dto.OrderItems, 0, len(req.Items))
	for _, item := range req.Items {
		productsIDs = append(productsIDs, dto.OrderItems{
			ProductID: item.ProductId,
			Quantity:  item.Quantity,
		})
	}

	o, err := h.orderUC.PlaceOrder(userID, productsIDs)
	if err != nil {
		return nil, h.mapper.MapError(err)
	}
	return h.mapOrderWithItems(o), nil
}

func (h *GRPCHandler) GetOrderById(
	ctx context.Context,
	req *pbApi.OrderId,
) (*pbApi.OrderWithItems, error) {
	userID, err := h.getUserIDFromMetadata(ctx)
	if err != nil {
		return nil, h.mapper.MapError(err)
	}
	orderID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, h.mapper.MapError(err)
	}
	o, err := h.orderUC.GetOrderByIDAndUserID(orderID, userID)

	if err != nil {
		return nil, h.mapper.MapError(err)
	}

	return h.mapOrderWithItems(o), nil
}

func (h *GRPCHandler) GetOrdersByUserId(
	ctx context.Context,
	req *emptypb.Empty,
) (*pbApi.GetOrdersByUserIdResponse, error) {
	userID, err := h.getUserIDFromMetadata(ctx)
	if err != nil {
		return nil, h.mapper.MapError(err)
	}
	orders, err := h.orderUC.GetOrdersByUserID(userID)
	if err != nil {
		return nil, h.mapper.MapError(err)
	}
	return &pbApi.GetOrdersByUserIdResponse{
		Orders: h.mapOrderToResponse(orders),
	}, nil
}

func (h *GRPCHandler) DeleteOrder(
	ctx context.Context,
	req *pbApi.OrderId,
) (*pbApi.DeleteOrderResponse, error) {
	userID, err := h.getUserIDFromMetadata(ctx)
	if err != nil {
		return nil, h.mapper.MapError(err)
	}
	orderID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, h.mapper.MapError(err)
	}
	if err = h.orderUC.DeleteOrder(orderID, userID); err != nil {
		return nil, err
	}

	return &pbApi.DeleteOrderResponse{
		OrderId: &pbApi.OrderId{
			Id: orderID.String(),
		},
		Status: "deleted",
		Message: &types.MessageResponse{
			Message: "Order deleted successfully",
		},
	}, nil
}

func NewGRPCHandler(orderUC order.UseCase) *GRPCHandler {
	return &GRPCHandler{
		orderUC: orderUC,
	}
}
