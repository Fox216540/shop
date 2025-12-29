package payment

import (
	domain "github.com/Fox216540/shop/payment-service/domain/payment"
	"github.com/google/uuid"
)

type service struct {
	r domain.Repository
	p domain.Provider
}

func NewService(repo domain.Repository, provider domain.Provider) UseCase {
	return &service{r: repo, p: provider}
}

func (s *service) CreatePayment(idOrder uuid.UUID, value, currency, description string) (payment domain.Payment, error error) {
	p, err := s.p.CreatePaymentByOrderID(idOrder, value, currency, description)
	if err != nil {
		return domain.Payment{}, err
	}
	if err = s.r.Save(p); err != nil {
		return domain.Payment{}, err
	}
	return p, nil
}
