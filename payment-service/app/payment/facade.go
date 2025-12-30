package payment

import (
	"github.com/Fox216540/shop/payment-service/domain/payment"
	"github.com/google/uuid"
)

type Facade struct {
	repo   payment.Repository
	router *ProviderRouter
}

func NewFacade(
	repo payment.Repository,
	router *ProviderRouter,
) *Facade {
	return &Facade{repo: repo, router: router}
}

func (f *Facade) CreatePayment(
	idOrder uuid.UUID,
	value, currency, description string,
) (payment.Payment, error) {

	provider, err := f.router.ByCurrency(currency)
	if err != nil {
		return payment.Payment{}, err
	}

	uc := NewService(f.repo, provider)
	return uc.CreatePayment(idOrder, value, currency, description)
}
