package orderdto

type CreateOrderRequest struct {
	Products []OrderProduct `json:"products" binding:"required,dive"` // массив объектов
	UserID   string         `json:"user_id" binding:"required,uuid"`
}

type OrderProduct struct {
	ProductID string `json:"product_id" binding:"required,uuid"`
	Quantity  int    `json:"quantity" binding:"required,min=1"`
}
