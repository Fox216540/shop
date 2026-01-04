package di

import (
	"github.com/Fox216540/shop/payment-service/app/payment"
	db "github.com/Fox216540/shop/payment-service/infra/db/core"
	paymentrepo "github.com/Fox216540/shop/payment-service/infra/payment/repository"
	"github.com/Fox216540/shop/payment-service/infra/payment/yokassa"
)

func NewPaymentModule(core *db.Database) payment.UseCase {
	repo := paymentrepo.NewRepository(core)
	provYo := yokassa.NewProvider()
	testprovYo := yokassa.NewProvider()
	providers := payment.NewProviderRouter(provYo, testprovYo)
	facade := payment.NewFacade(repo, providers)

	return facade
}
