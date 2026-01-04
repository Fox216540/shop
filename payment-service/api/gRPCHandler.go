package api

import (
	"context"
	"github.com/Fox216540/shop/payment-service/app/payment"
	pbInterservice "github.com/Fox216540/shop/proto/payment-service/gen/interservice"
	"github.com/google/uuid"
)

type GRPCHandler struct {
	paymentUS payment.UseCase
	pbInterservice.UnimplementedInterserviceServiceServer
	mapper *ErrorMapper
}

func NewGRPCHandler(useCase payment.UseCase) *GRPCHandler {
	return &GRPCHandler{
		paymentUS: useCase,
	}
}

func (h *GRPCHandler) CreatePayment(
	ctx context.Context,
	req *pbInterservice.CreatePaymentRequest,
) (*pbInterservice.CreatePaymentResponse, error) {
	orderID, err := uuid.Parse(req.OrderId)
	if err != nil {
		return nil, h.mapper.MapError(ctx, err)
	}
	newPayment, err := h.paymentUS.CreatePayment(
		orderID, req.Value,
		req.Currency, req.Description,
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
