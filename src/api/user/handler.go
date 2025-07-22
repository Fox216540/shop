package user

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"shop/src/api/user/di"
	"shop/src/api/user/dto"
	"shop/src/app/user"
)

// TODO: разбить на под функции
func Handler(r *gin.Engine) {
	us := di.GetUserService()
	// Register
	r.POST("/user", registerHandler(us))
	// Login
	//r.POST("/user/login", loginHandler(us))
	// Logout
	//r.POST("/user/logout", logoutHandler(us))
	// LogoutAll
	//r.POST("/user/logoutall", logoutAllHandler(us))
	// Update
	//r.PATCH("/user/", updateHandler(us))
	// Delete
	r.DELETE("/user/", deleteHandler(us))
	// Orders
	//r.GET("/user/orders", ordersHandler(us))
	// CreateOrder
	r.POST("/user/order/create", createOrderHandler(us))
	// DeleteOrder
	r.POST("/user/order/delete", deleteOrderHandler(us))
}

func registerHandler(us user.UseCase) gin.HandlerFunc {
	//Register
	return func(c *gin.Context) {
		var r dto.RegisterRequest

		if err := c.ShouldBindJSON(&r); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		NewUser, err := us.Register(r)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		NewUserDTO := dto.UserResponse{
			ID:       NewUser.ID,
			Username: NewUser.Username,
			Message:  "User created successfully",
		}

		c.JSON(http.StatusOK, NewUserDTO)
	}
}

func deleteHandler(us user.UseCase) gin.HandlerFunc {
	//Delete
	return func(c *gin.Context) {
		var r dto.TestDeleteRequest

		if err := c.ShouldBindJSON(&r); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}
		uuidID, err := uuid.Parse(r.ID)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		deletedUser, err := us.Delete(uuidID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
			return
		}

		userDeletedDTO := dto.UserResponse{
			ID:       deletedUser.ID,
			Username: deletedUser.Username,
			Message:  "User deleted successfully",
		}

		c.JSON(http.StatusOK, userDeletedDTO)
	}
}

func createOrderHandler(us user.UseCase) gin.HandlerFunc {
	//CreateOrder
	return func(c *gin.Context) {
		var r dto.TestCreateOrderRequest

		if err := c.ShouldBindJSON(&r); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		o, err := us.CreateOrder(r)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
			return
		}

		c.JSON(http.StatusOK, dto.CreateOrderResponse{
			ID:       o.ID,
			UserID:   o.UserID,
			OrderNum: o.OrderNum,
			Total:    o.Total,
			Status:   o.Status,
		})
	}
}

func deleteOrderHandler(us user.UseCase) gin.HandlerFunc {
	//DeleteOrder
	return func(c *gin.Context) {
		var r dto.TestDeleteOrderRequest

		if err := c.ShouldBindJSON(&r); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		o, err := us.DeleteOrder(r)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})

	}
}
