package app

import "github.com/Fox216540/shop/basket-service/app/globalError"

const domain = "Basket"

type BasketContainsMultipleCurrenciesError struct {
	*globalError.AppServerError
}

func (e *BasketContainsMultipleCurrenciesError) Error() string {
	return e.AppServerError.Error()
}

func (e *BasketContainsMultipleCurrenciesError) Unwrap() error {
	return e.AppServerError
}

func NewBasketContainsMultipleCurrenciesError(err error) *BasketContainsMultipleCurrenciesError {
	return &BasketContainsMultipleCurrenciesError{
		AppServerError: globalError.NewAppServerError("Basket contains multiple currencies", domain, err),
	}
}

type BasketInvalidItemAmountError struct {
	*globalError.AppServerError
}

func (e *BasketInvalidItemAmountError) Error() string {
	return e.AppServerError.Error()
}

func (e *BasketInvalidItemAmountError) Unwrap() error {
	return e.AppServerError
}

func NewBasketInvalidItemAmountError(err error) *BasketInvalidItemAmountError {
	return &BasketInvalidItemAmountError{
		AppServerError: globalError.NewAppServerError("Invalid basket item amount", domain, err),
	}
}
