package api

import (
	"context"
	"errors"
	"net/http"

	shopApiGen "github.com/Fox216540/shop/api/gen"
	"github.com/Fox216540/shop/apigateway-service/core/exception"
	"github.com/Fox216540/shop/apigateway-service/core/logger"
	domainAuth "github.com/Fox216540/shop/apigateway-service/domain/auth"
	domainBasket "github.com/Fox216540/shop/apigateway-service/domain/basket"
	domainCatalog "github.com/Fox216540/shop/apigateway-service/domain/catalog"
	domainOrder "github.com/Fox216540/shop/apigateway-service/domain/order"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type HTTPMapper struct {
	log logger.Logger
}

func NewHTTPMapper(log logger.Logger) *HTTPMapper {
	return &HTTPMapper{log: log}
}

func (m *HTTPMapper) MapError(ctx context.Context, e error) (statusCode int, message string) {
	var (
		ce *exception.ConflictError
		ue *exception.UnauthorizedError
		se *exception.ServerError
	)

	switch {
	case errors.As(e, &ce):
		return http.StatusConflict, e.Error()
	case errors.As(e, &ue):
		return http.StatusUnauthorized, messages.Unauthorized
	}

	if st, ok := status.FromError(e); ok {
		m.log.Info(ctx, "gRPC error: "+st.Message())

		switch st.Code() {
		case codes.InvalidArgument:
			return http.StatusBadRequest, st.Message()

		case codes.NotFound:
			return http.StatusNotFound, st.Message()

		case codes.Unauthenticated:
			return http.StatusUnauthorized, st.Message()

		case codes.PermissionDenied:
			return http.StatusForbidden, st.Message()

		case codes.Unavailable:
			return http.StatusServiceUnavailable, messages.ServiceUnavailable

		default:
			return http.StatusInternalServerError, st.Message()
		}
	}

	if errors.As(e, &se) {
		m.log.Error(ctx, e)
		return http.StatusInternalServerError, messages.ServerError
	}

	m.log.Info(ctx, "Unknown Server Error")
	m.log.Error(ctx, e)
	return http.StatusInternalServerError, messages.ServerError
}

func moneyResponse(amount, currency string) shopApiGen.Money {
	return shopApiGen.Money{
		Amount:   amount,
		Currency: currency,
	}
}

func userWithTokenResponse(name string, tokens domainAuth.Tokens, message string) shopApiGen.UserWithTokenResponse {
	return shopApiGen.UserWithTokenResponse{
		AccessToken: tokens.AccessToken,
		Message:     message,
		Name:        name,
	}
}

func userResponse(name, message string) shopApiGen.UserResponse {
	return shopApiGen.UserResponse{
		Name:    name,
		Message: message,
	}
}

func categoriesToResponse(categories []domainCatalog.Category) []shopApiGen.CategoryResponse {
	resp := make([]shopApiGen.CategoryResponse, 0, len(categories))
	for _, category := range categories {
		resp = append(resp, shopApiGen.CategoryResponse{
			Id:   category.ID,
			Name: category.Name,
		})
	}
	return resp
}

func productShortResponse(product domainOrder.Product) shopApiGen.ProductShort {
	return shopApiGen.ProductShort{
		Id:    product.ID,
		Img:   product.Img,
		Name:  product.Name,
		Price: moneyResponse(product.Price, product.Currency),
	}
}

func orderStatus(status string) shopApiGen.OrderResponseStatus {
	return shopApiGen.OrderResponseStatus(status)
}

func orderWithItemsStatus(status string) shopApiGen.OrderWithItemsResponseStatus {
	return shopApiGen.OrderWithItemsResponseStatus(status)
}

func orderToResponse(order domainOrder.Order) shopApiGen.OrderResponse {
	return shopApiGen.OrderResponse{
		Id:          order.ID,
		OrderNumber: order.OrderNum,
		Status:      orderStatus(order.Status),
		Total:       moneyResponse(order.Total, order.Currency),
	}
}

func orderWithItemsToResponse(order domainOrder.OrderWithItems) shopApiGen.OrderWithItemsResponse {
	items := make([]shopApiGen.OrderItem, 0, len(order.Items))
	for _, item := range order.Items {
		items = append(items, shopApiGen.OrderItem{
			Product:  productShortResponse(item.Product),
			Quantity: int64(item.Quantity),
		})
	}

	return shopApiGen.OrderWithItemsResponse{
		Id:          order.Order.ID,
		OrderNumber: order.Order.OrderNum,
		Status:      orderWithItemsStatus(order.Order.Status),
		Total:       moneyResponse(order.Order.Total, order.Order.Currency),
		OrderItems:  items,
	}
}

func ordersToResponse(orders []domainOrder.Order) []shopApiGen.OrderResponse {
	resp := make([]shopApiGen.OrderResponse, 0, len(orders))
	for _, order := range orders {
		resp = append(resp, orderToResponse(order))
	}
	return resp
}

func orderDeletedResponse(id uuid.UUID, message, status string) shopApiGen.OrderDeletedResponse {
	return shopApiGen.OrderDeletedResponse{
		Id:      id,
		Message: message,
		Status:  shopApiGen.OrderDeletedResponseStatus(status),
	}
}

func productToResponse(product domainCatalog.Product) shopApiGen.ProductResponse {
	return shopApiGen.ProductResponse{
		Id:          product.ID,
		Name:        product.Name,
		Img:         product.Img,
		Price:       moneyResponse(product.Price, product.Currency),
		CategoryId:  product.CategoryID,
		Description: product.Description,
		Stock:       int64(product.Stock),
	}
}

func productsToResponse(products []domainCatalog.Product) []shopApiGen.ProductResponse {
	resp := make([]shopApiGen.ProductResponse, 0, len(products))
	for _, product := range products {
		resp = append(resp, productToResponse(product))
	}
	return resp
}

func basketItemResponse(item domainBasket.Item) shopApiGen.BasketItem {
	return shopApiGen.BasketItem{
		Product: shopApiGen.ProductShort{
			Id:    item.Product.ID,
			Name:  item.Product.Name,
			Img:   item.Product.Img,
			Price: moneyResponse(item.Product.Price, item.Product.Currency),
		},
		Quantity: int64(item.Quantity),
	}
}

func basketToResponse(basket domainBasket.Basket) shopApiGen.BasketResponse {
	items := make([]shopApiGen.BasketItem, 0, len(basket.Items))
	for _, item := range basket.Items {
		items = append(items, basketItemResponse(item))
	}

	return shopApiGen.BasketResponse{
		Items: items,
		Total: moneyResponse(basket.Total.Amount, basket.Total.Currency),
	}
}
