package dto

import (
	"github.com/google/uuid"
	"shop/src/api/product/dto"
)

type ItemsResponse struct {
	Product  dto.ProductResponse `json:"products"`
	Quantity int                 `json:"quantity"`
}

type OrderResponse struct {
	ID       uuid.UUID       `json:"id"`
	UserID   uuid.UUID       `json:"user_id"`
	OrderNum string          `json:"order_num"`
	Total    float64         `json:"total"`
	Status   string          `json:"status"`
	Items    []ItemsResponse `json:"items"`
}
