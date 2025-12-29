package payment

type Repository interface {
	Save(payment Payment) error
}
