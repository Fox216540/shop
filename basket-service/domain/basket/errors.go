package basket

import "github.com/Fox216540/shop/basket-service/core/exception"

const domain = "Basket"

func NewBasketNotFoundError(err error) *exception.NotFoundError {
	return exception.NewNotFoundError("Basket not found", domain, err)
}

func NewItemNotFoundError(err error) *exception.NotFoundError {
	return exception.NewNotFoundError("Basket item not found", domain, err)
}

func NewInvalidQuantityError(err error) *exception.DomainError {
	return exception.NewDomainException("Invalid quantity", domain, err)
}

func NewInternalError(msg string, err error) *exception.ServerError {
	return exception.NewServerError(msg, domain, "Domain", err)
}
