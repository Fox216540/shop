package orderhandler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	di "shop/src/api/order/di"
	dto "shop/src/api/order/dto"
	"shop/src/domain/order"
	"shop/src/domain/product"
)

func OrderHandler(r *gin.Engine) {
	// Здесь вы можете определить обработчики для маршрутов, связанных с заказами
	// Например:
	//r.GET("/orders", func(c *gin.Context) {
	//
	//})
	orderService := di.GetOrderService()
	//
	r.POST("/orders", func(c *gin.Context) {
		var r dto.CreateOrderRequest

		if err := c.ShouldBindJSON(&r); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters", "details": err.Error()})
			return
		}

		var o order.Order

		if r.UserID != "" {
			userID, err := uuid.Parse(r.UserID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID format", "details": err.Error()})
				return
			}
			o.UserID = userID
			o.ProductItems = make([]*order.Item, len(r.Products))

			for i, p := range r.Products {
				productID, err := uuid.Parse(p.ProductID)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid Product ID format for product %d: %s", i+1, err.Error())})
					return
				}
				if p.Quantity <= 0 {
					c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid quantity for product %d: must be greater than 0", i+1)})
					return
				}
				o.ProductItems[i] = &order.Item{
					Product: product.Product{
						ID: productID,
					},
					Quantity: p.Quantity,
				}
			}

			if len(o.ProductItems) == 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "At least one product must be specified in the order"})
				return
			}

		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
			return
		}
		fmt.Print(&o)
		updatedOrder, err := orderService.PlaceOrder(o)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to place order", "details": err.Error()})
			return
		}
		c.JSON(http.StatusOK, updatedOrder)
	})
	//r.DELETE("/orders/:id", cancelOrder)

	// Пример простого обработчика
	//r.GET("/orders", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "List of orders",
	//	})
	//})

}
