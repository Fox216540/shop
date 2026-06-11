package app

import (
	"github.com/Fox216540/shop/payment-service/domain/payment"
	"github.com/google/uuid"
)

type service struct {
	r payment.Repository
	p payment.Provider
}

func NewService(repo payment.Repository, provider payment.Provider) UseCase {
	return &service{
		r: repo,
		p: provider,
	}
}

func (s *service) CreatePayment(
	idOrder uuid.UUID,
	value, currency, description, returnURL string,
) (payment.Payment, error) {
	p, err := s.p.CreatePaymentByOrderID(idOrder, value, currency, description, returnURL)
	if err != nil {
		return payment.Payment{}, err
	}

	if err := s.r.Save(p); err != nil {
		return payment.Payment{}, err
	}

	return p, nil
}
