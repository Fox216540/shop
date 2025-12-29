package payment

import (
	domain "github.com/Fox216540/shop/payment-service/domain/payment"
	yo "github.com/rvinnie/yookassa-sdk-go/yookassa"
	cm "github.com/rvinnie/yookassa-sdk-go/yookassa/common"
	"github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
	"strconv"
)

type repository struct {
	client    yo.Client
	handler   yo.PaymentHandler
	returnURL string
}

func NewRepository(client yo.Client, returnUrl string, paymentHand yo.PaymentHandler) payment.Repository {
	return &repository{
		client:    client,
		returnURL: returnUrl,
		handler:   paymentHand,
	}
}

func (r *repository) CreatePayment(value, currency, description string) (domain.Payment, error) {
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
	val, err := strconv.Atoi(newPayment.Amount.Value)
	if err != nil {
		return domain.Payment{}, err
	}
	return domain.Payment{
		Amount: domain.Amount{
			Value:    val,
			Currency: newPayment.Amount.Currency,
		},
		Method: newPayment.PaymentMethodID,
		//TODO: Разобраться
		ReturnURL: newPayment.Confirmation,
		Status:    newPayment.Status,
	}, nil
}
