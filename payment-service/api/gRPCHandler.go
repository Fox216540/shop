package api

import (
	"context"
	"github.com/Fox216540/shop/payment-service/app"
	pbInterservice "github.com/Fox216540/shop/proto/payment-service/gen/interservice"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCHandler struct {
	paymentUS app.UseCase
	pbInterservice.UnimplementedInterserviceServiceServer
	mapper *ErrorMapper
}

func NewGRPCHandler(useCase app.UseCase, mapper *ErrorMapper) *GRPCHandler {
	return &GRPCHandler{
		paymentUS: useCase,
		mapper:    mapper,
	}
}

func (h *GRPCHandler) CreatePayment(
	ctx context.Context,
	req *pbInterservice.CreatePaymentRequest,
) (*pbInterservice.CreatePaymentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, messages.InvalidArgument)
	}

	if req.OrderId == "" || req.Value == "" || req.Currency == "" || req.ReturnUrl == "" {
		return nil, status.Error(codes.InvalidArgument, messages.InvalidArgument)
	}

	orderID, err := uuid.Parse(req.OrderId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, messages.InvalidArgument)
	}
	newPayment, err := h.paymentUS.CreatePayment(
		orderID, req.Value,
		req.Currency, req.Description,
		req.ReturnUrl,
	)
	if err != nil {
		return nil, h.mapper.MapError(ctx, err)
	}
	return &pbInterservice.CreatePaymentResponse{
		PaymentId:       newPayment.ID,
		ConfirmationUrl: newPayment.ReturnURL,
		Status:          string(newPayment.Status),
	}, nil
}
