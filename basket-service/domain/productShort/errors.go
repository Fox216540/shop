package productShort

import "github.com/Fox216540/shop/basket-service/core/exception"

const domain = "ProductShort"

func NewNotFoundError(err error) *exception.NotFoundError {
	return exception.NewNotFoundError("Product not found", domain, err)
}

func NewInvalidUUIDError(err error) *exception.DomainError {
	return exception.NewDomainException("Invalid product UUID", domain, err)
}
