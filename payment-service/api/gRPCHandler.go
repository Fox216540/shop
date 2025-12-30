package api

import (
	"context"
	"github.com/Fox216540/shop/payment-service/app/payment"
	pbInterservice "github.com/Fox216540/shop/proto/payment-service/gen/interservice"
	"github.com/google/uuid"
)

type GRPCHandler struct {
	paymentFacade *payment.Facade
	pbInterservice.UnimplementedInterserviceServiceServer
	mapper *ErrorMapper
}

func NewGRPCHandler(facade *payment.Facade) *GRPCHandler {
	return &GRPCHandler{
		paymentFacade: facade,
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
	newPayment, err := h.paymentFacade.CreatePayment(
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
