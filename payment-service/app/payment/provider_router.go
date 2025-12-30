package payment

import (
	"fmt"
	"github.com/Fox216540/shop/payment-service/domain/payment"
)

type ProviderRouter struct {
	yoo    payment.Provider
	testYo payment.Provider
}

func NewProviderRouter(yoo, test payment.Provider) *ProviderRouter {
	return &ProviderRouter{yoo: yoo, testYo: test}
}

func (r *ProviderRouter) ByCurrency(currency string) (payment.Provider, error) {
	switch currency {
	case "RUB":
		return r.yoo, nil
	case "EUR":
		return nil, fmt.Errorf(fmt.Sprintf("test"))
	default:
		//TODO: Исправить
		return nil, fmt.Errorf(fmt.Sprintf("test"))
	}
}
