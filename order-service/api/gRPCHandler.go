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
}

func (h *GRPCHandler) mapOrder(
	order orderDomain.Order,
) *types.Order {
	items := make([]*types.Item, 0, len(order.OrderItems))
	for _, item := range order.OrderItems {
		items = append(items, &types.Item{
			Product: &types.Product{
				Id:    item.Product.ID.String(),
				Name:  item.Product.Name,
				Img:   item.Product.Img,
				Price: item.Product.Price,
			},
			Quantity: item.Quantity,
		})
	}
	return &types.Order{
		Id:       order.ID.String(),
		OrderNum: order.OrderNum,
		Items:    items,
		Total:    order.Total,
		Status:   order.Status,
	}
}

func (h *GRPCHandler) mapOrderForListToResponse(
	orders []orderDomain.Order,
) []*pbApi.OrderForList {
	ordersForList := make([]*pbApi.OrderForList, 0, len(orders))
	for _, o := range orders {
		ordersForList = append(ordersForList, &pbApi.OrderForList{
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
) (*types.Order, error) {
	userID, err := h.getUserIDFromMetadata(ctx)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	return h.mapOrder(o), nil
}

func (h *GRPCHandler) GetOrderById(
	ctx context.Context,
	req *types.OrderId,
) (*types.Order, error) {
	userID, err := h.getUserIDFromMetadata(ctx)
	if err != nil {
		return nil, err
	}
	orderID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	o, err := h.orderUC.GetOrderByIDAndUserID(orderID, userID)

	if err != nil {
		return nil, err
	}

	return h.mapOrder(o), nil
}

func (h *GRPCHandler) GetOrdersByUserId(
	ctx context.Context,
	req *emptypb.Empty,
) (*pbApi.GetOrdersByUserIdResponse, error) {
	userID, err := h.getUserIDFromMetadata(ctx)
	if err != nil {
		return nil, err
	}
	orders, err := h.orderUC.GetOrdersByUserID(userID)
	if err != nil {
		return nil, err
	}
	return &pbApi.GetOrdersByUserIdResponse{
		Orders: h.mapOrderForListToResponse(orders),
	}, nil
}

func (h *GRPCHandler) DeleteOrder(
	ctx context.Context,
	req *types.OrderId,
) (*pbApi.DeleteOrderResponse, error) {
	userID, err := h.getUserIDFromMetadata(ctx)
	if err != nil {
		return nil, err
	}
	orderID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	if err = h.orderUC.DeleteOrder(orderID, userID); err != nil {
		return nil, err
	}

	return &pbApi.DeleteOrderResponse{
		OrderId: &types.OrderId{
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
