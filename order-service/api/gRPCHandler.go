package api

import (
	"context"

	order "github.com/Fox216540/shop/order-service/app"
	"github.com/Fox216540/shop/order-service/app/dto"
	orderDomain "github.com/Fox216540/shop/order-service/domain/order"
	types "github.com/Fox216540/shop/proto/common/gen"
	pbApi "github.com/Fox216540/shop/proto/order-service/gen/api"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCHandler struct {
	orderUC order.UseCase
	pbApi.UnimplementedApiServiceServer
	mapper *ErrorMapper
}

func NewGRPCHandler(orderUC order.UseCase, mapper *ErrorMapper) *GRPCHandler {
	return &GRPCHandler{
		orderUC: orderUC,
		mapper:  mapper,
	}
}

func mapMoney(amount, currency string) *types.Money {
	return &types.Money{
		Amount:   amount,
		Currency: currency,
	}
}

func (h *GRPCHandler) mapOrder(order orderDomain.Order) *pbApi.Order {
	return &pbApi.Order{
		Id:       order.ID.String(),
		OrderNum: order.OrderNum,
		Price:    mapMoney(order.Total.StringFixed(2), order.Currency),
		Status:   order.Status,
	}
}

func (h *GRPCHandler) mapOrderWithItems(order orderDomain.Order) *pbApi.OrderWithItems {
	items := make([]*pbApi.Item, 0, len(order.OrderItems))
	for _, item := range order.OrderItems {
		items = append(items, &pbApi.Item{
			Product: &types.Product{
				Id:   item.Product.ID.String(),
				Name: item.Product.Name,
				Img:  item.Product.Img,
				Price: mapMoney(
					item.Product.Price.Amount.StringFixed(2),
					item.Product.Price.Currency,
				),
			},
			Quantity: item.Quantity,
		})
	}

	return &pbApi.OrderWithItems{
		Order: h.mapOrder(order),
		Items: items,
	}
}

func (h *GRPCHandler) mapOrders(orders []orderDomain.Order) []*pbApi.Order {
	items := make([]*pbApi.Order, 0, len(orders))
	for _, o := range orders {
		items = append(items, h.mapOrder(o))
	}
	return items
}

func (h *GRPCHandler) getUserIDFromMetadata(ctx context.Context) (uuid.UUID, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return uuid.Nil, status.Error(codes.InvalidArgument, messages.MetadataNotFound)
	}

	values := md.Get("user_id")
	if len(values) == 0 {
		return uuid.Nil, status.Error(codes.InvalidArgument, messages.UserIDMetadataNotFound)
	}

	userID, err := uuid.Parse(values[0])
	if err != nil {
		return uuid.Nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return userID, nil
}

func (h *GRPCHandler) CreateOrder(ctx context.Context, req *pbApi.CreateOrderRequest) (*pbApi.Order, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, messages.InvalidArgument)
	}
	userID, err := h.getUserIDFromMetadata(ctx)
	if err != nil {
		return nil, h.mapper.MapError(ctx, err)
	}

	items := make([]dto.OrderItems, 0, len(req.Items))
	for _, item := range req.Items {
		items = append(items, dto.OrderItems{
			ProductID: item.ProductId,
			Quantity:  item.Quantity,
			Value:     item.Value,
			Currency:  item.Currency,
		})
	}

	o, err := h.orderUC.PlaceOrder(userID, items)
	if err != nil {
		return nil, h.mapper.MapError(ctx, err)
	}

	return h.mapOrder(o), nil
}

func (h *GRPCHandler) GetOrderById(ctx context.Context, req *pbApi.OrderId) (*pbApi.OrderWithItems, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, messages.InvalidArgument)
	}
	userID, err := h.getUserIDFromMetadata(ctx)
	if err != nil {
		return nil, h.mapper.MapError(ctx, err)
	}

	orderID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, h.mapper.MapError(ctx, status.Error(codes.InvalidArgument, err.Error()))
	}

	o, err := h.orderUC.GetOrderByIDAndUserID(orderID, userID)
	if err != nil {
		return nil, h.mapper.MapError(ctx, err)
	}

	return h.mapOrderWithItems(o), nil
}

func (h *GRPCHandler) GetOrdersByUserId(ctx context.Context, _ *emptypb.Empty) (*pbApi.GetOrdersByUserIdResponse, error) {
	userID, err := h.getUserIDFromMetadata(ctx)
	if err != nil {
		return nil, h.mapper.MapError(ctx, err)
	}

	orders, err := h.orderUC.GetOrdersByUserID(userID)
	if err != nil {
		return nil, h.mapper.MapError(ctx, err)
	}

	return &pbApi.GetOrdersByUserIdResponse{
		Orders: h.mapOrders(orders),
	}, nil
}

func (h *GRPCHandler) DeleteOrder(ctx context.Context, req *pbApi.OrderId) (*pbApi.DeleteOrderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, messages.InvalidArgument)
	}
	userID, err := h.getUserIDFromMetadata(ctx)
	if err != nil {
		return nil, h.mapper.MapError(ctx, err)
	}

	orderID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, h.mapper.MapError(ctx, status.Error(codes.InvalidArgument, err.Error()))
	}

	if err := h.orderUC.DeleteOrder(orderID, userID); err != nil {
		return nil, h.mapper.MapError(ctx, err)
	}

	return &pbApi.DeleteOrderResponse{
		OrderId: &pbApi.OrderId{Id: orderID.String()},
		Status:  "deleted",
		Message: &types.MessageResponse{
			Message: messages.OrderDeletedSuccessfully,
		},
	}, nil
}
