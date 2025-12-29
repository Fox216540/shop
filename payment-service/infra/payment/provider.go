package payment

import (
	domain "github.com/Fox216540/shop/payment-service/domain/payment"
	db "github.com/Fox216540/shop/payment-service/infra/db/core"
	"github.com/google/uuid"
	yo "github.com/rvinnie/yookassa-sdk-go/yookassa"
	cm "github.com/rvinnie/yookassa-sdk-go/yookassa/common"
	"github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
	"github.com/shopspring/decimal"
)

type provider struct {
	client    yo.Client
	handler   yo.PaymentHandler
	returnURL string
	db        *db.Database
}

func NewProvider(
	client yo.Client,
	returnUrl string,
	paymentHand yo.PaymentHandler,
) domain.Provider {
	return &provider{
		client:    client,
		returnURL: returnUrl,
		handler:   paymentHand,
	}
}

func (r *provider) CreatePaymentByOrderID(idOrder uuid.UUID, value, currency, description string) (domain.Payment, error) {
	test := &yoopayment.Payment{
		Amount: &cm.Amount{
			Value:    value,
			Currency: currency,
		},
		Confirmation: yoopayment.Redirect{
			Type:      "redirect",
			ReturnURL: r.returnURL,
		},
		Description: description,
	}
	newPayment, err := r.handler.CreatePayment(test)
	if err != nil {
		return domain.Payment{}, err
	}
	val, err := decimal.NewFromString(newPayment.Amount.Value)
	if err != nil {
		return domain.Payment{}, err
	}
	url, ok := newPayment.Confirmation.(*yoopayment.Redirect)
	if !ok {
		return domain.Payment{}, nil
	}

	return domain.Payment{
		ID:      newPayment.ID,
		OrderID: idOrder,
		Amount: domain.Amount{
			Value:    val,
			Currency: newPayment.Amount.Currency,
		},
		Method:    newPayment.PaymentMethodID,
		ReturnURL: url.ReturnURL,
		Status:    newPayment.Status,
	}, nil
}
