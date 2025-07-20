package user

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"shop/src/api/user/di"
	"shop/src/api/user/dto"
)

func Handler(r *gin.Engine) {
	us := di.GetUserService()
	// Register
	r.POST("/user", func(c *gin.Context) {
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

	})
	// Login
	// Logout
	// LogoutAll
	// Update
	// Delete
	r.DELETE("/user/", func(c *gin.Context) {
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
	})
	// Orders
	// DeleteOrder
	// CreateOrder
	r.POST("/user/orders", func(c *gin.Context) {
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
	})

}
