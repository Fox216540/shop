package di

import (
	"github.com/Fox216540/shop/payment-service/app"
	"github.com/Fox216540/shop/payment-service/domain/payment"
	db "github.com/Fox216540/shop/payment-service/infra/db/core"
	paymentrepo "github.com/Fox216540/shop/payment-service/infra/payment/repository"
)

func NewPaymentModule(core *db.Database, provider payment.Provider) app.UseCase {
	repo := paymentrepo.NewRepository(core)
	return app.NewService(repo, provider)
}
