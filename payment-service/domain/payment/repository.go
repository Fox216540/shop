package payment

type Repository interface {
	CreatePayment(value, currency, description string) (Payment, error)
}
