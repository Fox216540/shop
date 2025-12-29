package payment

import pay "github.com/rvinnie/yookassa-sdk-go/yookassa/payment"

type Amount struct {
	Value    int
	Currency string
}

type Payment struct {
	Amount      Amount
	Method      string
	ReturnURL   string
	Description string
	Status      pay.Status
}
