package api

import (
	"context"
	order "github.com/Fox216540/shop/order-service/app"
	"github.com/Fox216540/shop/order-service/app/dto"
	orderDomain "github.com/Fox216540/shop/order-service/domain/order"
	pb "github.com/Fox216540/shop/proto/order-service/gen"
	"github.com/google/uuid"
)

type GRPCHandler struct {
	orderUC order.UseCase
	pb.UnimplementedInternalOrderServiceServer
}

func (h *GRPCHandler) mapOrder(order orderDomain.Order) *pb.Order {
	items := make([]*pb.Item, 0, len(order.OrderItems))
	for _, item := range order.OrderItems {
		items = append(items, &pb.Item{
			Product: &pb.Product{
				Id:    item.Product.ID.String(),
				Name:  item.Product.Name,
				Img:   item.Product.Img,
				Price: item.Product.Price,
			},
			Quantity: item.Quantity,
		})
	}
	return &pb.Order{
		Id:       order.ID.String(),
		OrderNum: order.OrderNum,
		UserId:   order.UserID.String(),
		Items:    items,
		Total:    order.Total,
		Status:   order.Status,
	}
}

func (h *GRPCHandler) PlaceOrder(ctx context.Context, req *pb.PlaceOrderRequest) (*pb.Order, error) {
	userID, err := uuid.Parse(req.UserId)
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

	o, err := h.orderUC.Place(userID, productsIDs)
	if err != nil {
		return nil, err
	}
	return h.mapOrder(o), nil
}

func (h *GRPCHandler) CancelOrder(ctx context.Context, req *pb.CancelOrderRequest) (*pb.CancelOrderResponse, error) {
	orderID, err := uuid.Parse(req.OrderId)
	if err != nil {
		return nil, err
	}
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}
	orderID, err = h.orderUC.Cancel(orderID, userID)

	if err != nil {
		return nil, err
	}

	return &pb.CancelOrderResponse{
		OrderId: orderID.String(),
	}, nil
}

func (h *GRPCHandler) GetOrderById(ctx context.Context, req *pb.GetOrderByIdRequest) (*pb.Order, error) {
	orderID, err := uuid.Parse(req.OrderId)
	if err != nil {
		return nil, err
	}
	o, err := h.orderUC.GetByID(orderID)
	if err != nil {
		return nil, err
	}
	return h.mapOrder(o), nil
}

func (h *GRPCHandler) mapOrderForListToResponse(orders []orderDomain.Order) []*pb.OrderForList {
	ordersForList := make([]*pb.OrderForList, 0, len(orders))
	for _, o := range orders {
		ordersForList = append(ordersForList, &pb.OrderForList{
			Id:       o.ID.String(),
			OrderNum: o.OrderNum,
			UserId:   o.UserID.String(),
			Total:    o.Total,
			Status:   o.Status,
		})
	}
	return ordersForList
}

func (h *GRPCHandler) GetOrdersByUserId(ctx context.Context, req *pb.GetOrdersByUserIdRequest) (*pb.GetOrdersByUserIdResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}
	orders, err := h.orderUC.GetOrdersByUserID(userID)
	if err != nil {
		return nil, err
	}
	return &pb.GetOrdersByUserIdResponse{
		Orders: h.mapOrderForListToResponse(orders),
	}, nil
}

func NewGRPCHandler(orderUC order.UseCase) *GRPCHandler {
	return &GRPCHandler{
		orderUC: orderUC,
	}
}
