package yokassa

import (
	"errors"
	"strings"

	domain "github.com/Fox216540/shop/payment-service/domain/payment"
	"github.com/google/uuid"
	yo "github.com/rvinnie/yookassa-sdk-go/yookassa"
	cm "github.com/rvinnie/yookassa-sdk-go/yookassa/common"
	yoopayment "github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
	"github.com/shopspring/decimal"
)

type provider struct {
	handler *yo.PaymentHandler
}

func NewProvider(handler *yo.PaymentHandler) domain.Provider {
	return &provider{
		handler: handler,
	}
}

func (r *provider) CreatePaymentByOrderID(
	idOrder uuid.UUID,
	value, currency, description, returnURL string,
) (domain.Payment, error) {
	req := &yoopayment.Payment{
		Amount: &cm.Amount{
			Value:    value,
			Currency: currency,
		},
		Confirmation: yoopayment.Redirect{
			Type:      yoopayment.TypeRedirect,
			ReturnURL: returnURL,
		},
		Description: description,
	}

	newPayment, err := r.handler.CreatePayment(req)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "empty confirmation url") {
			return domain.Payment{}, domain.NewNoRedirectConfirmation(err)
		}
		return domain.Payment{}, NewInvalidCreatePayment(err)
	}

	if newPayment == nil || newPayment.Amount == nil {
		return domain.Payment{}, NewInvalidCreatePayment(errors.New("empty payment amount"))
	}

	val, err := decimal.NewFromString(newPayment.Amount.Value)
	if err != nil {
		return domain.Payment{}, NewInvalidParsePaymentAmount(err)
	}

	confirmationURL, err := r.handler.ParsePaymentLink(newPayment)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "empty confirmation url") ||
			strings.Contains(strings.ToLower(err.Error()), "unable to get link") {
			return domain.Payment{}, domain.NewNoRedirectConfirmation(err)
		}
		return domain.Payment{}, NewInvalidCreatePayment(err)
	}

	return domain.Payment{
		ID:      newPayment.ID,
		OrderID: idOrder,
		Amount: domain.Amount{
			Value:    val,
			Currency: newPayment.Amount.Currency,
		},
		Method:      newPayment.PaymentMethodID,
		ReturnURL:   confirmationURL,
		Description: description,
		Status:      newPayment.Status,
	}, nil
}
