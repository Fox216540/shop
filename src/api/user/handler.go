package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	jwtdi "shop/src/api/jwt/di"
	orderDTO "shop/src/api/order/dto"
	productDTO "shop/src/api/product/dto"
	"shop/src/api/user/di"
	"shop/src/api/user/dto"
	"shop/src/app/user"
	"shop/src/core/middleware"
	"shop/src/core/settings"
	"shop/src/domain/order"
	"strconv"
)

func Handler(r *gin.Engine) {
	us := di.GetUserService()
	jwtService := jwtdi.GetJwtService()
	// Register
	r.POST(
		"/user",
		registerHandler(us),
	)
	// Login
	r.POST(
		"/user/login",
		loginHandler(us),
	)
	// Logout
	r.POST(
		"/user/logout",
		logoutHandler(us),
	)
	// LogoutAll
	r.POST(
		"/user/logout-all",
		logoutAllHandler(us),
	)
	// Update email
	r.PATCH(
		"/user/email",
		middleware.JWTMiddleware(jwtService),
		updateEmailHandler(us),
	)
	// Update password
	r.PATCH(
		"/user/password",
		middleware.JWTMiddleware(jwtService),
		updatePasswordHandler(us),
	)
	// Update phone
	r.PATCH(
		"/user/phone",
		middleware.JWTMiddleware(jwtService),
		updatePhoneHandler(us),
	)
	// Update profile
	r.PATCH(
		"/user/profile",
		middleware.JWTMiddleware(jwtService),
		updateProfileHandler(us),
	)
	// Refresh tokens
	r.POST(
		"/user/refresh",
		refreshHandler(us),
	)
	// Delete
	r.DELETE(
		"/user",
		middleware.JWTMiddleware(jwtService),
		deleteHandler(us),
	)
	// Orders
	r.GET(
		"/user/orders",
		middleware.JWTMiddleware(jwtService),
		ordersHandler(us),
	)
	// CreateOrder
	r.POST(
		"/user/order/create",
		middleware.JWTMiddleware(jwtService),
		createOrderHandler(us),
	)
	// DeleteOrder
	r.POST(
		"/user/order/delete",
		middleware.JWTMiddleware(jwtService),
		deleteOrderHandler(us),
	)
}

func getUserIDFromContext(c *gin.Context) (uuid.UUID, error) {
	val, exists := c.Get("user_id")
	if !exists {
		return uuid.Nil, fmt.Errorf("user not authenticated")
	}
	id, ok := val.(uuid.UUID)
	if !ok {
		return uuid.Nil, fmt.Errorf("invalid user ID type")
	}
	return id, nil
}

func getBadRequestResponse(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
}

func getNotFoundResponse(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
}

func getInternalServerErrorResponse(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
}

func durationDiff(a, b string) (int, error) {
	t1, err := strconv.Atoi(a)
	if err != nil {
		return 0, err
	}
	t2, err := strconv.Atoi(b)
	if err != nil {
		return 0, err
	}
	return t1 - t2, nil
}

func registerHandler(us user.UseCase) gin.HandlerFunc {
	//Register
	return func(c *gin.Context) {
		var r dto.RegisterRequest

		if err := c.ShouldBindJSON(&r); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		NewUser, tokens, err := us.Register(r)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		ttl := settings.Config.RefreshTokenTTL
		bufferTime := settings.Config.BufferSeconds
		duration, err := durationDiff(ttl, bufferTime)
		if err != nil {
			getInternalServerErrorResponse(c)
			return
		}

		c.SetCookie(
			"refresh_token",     // имя
			tokens.RefreshToken, // значение, например user.RefreshToken
			duration,            // maxAge в секундах
			"/user/refresh",     // path
			"",                  // domain (можно пустое "")
			false,               // secure (true для HTTPS)
			true,                // HttpOnly
		)
		c.JSON(http.StatusOK, dto.UserWithTokenResponse{
			ID:          NewUser.ID,
			Name:        NewUser.Name,
			AccessToken: tokens.AccessToken,
			Message:     "User created successfully",
		})
	}
}

func loginHandler(us user.UseCase) gin.HandlerFunc {
	//Login
	return func(c *gin.Context) {
		var r dto.LoginRequest

		if err := c.ShouldBindJSON(&r); err != nil {
			getBadRequestResponse(c)
			return
		}

		User, tokens, err := us.Login(r.PhoneOrEmail, r.Password)

		if err != nil {
			getInternalServerErrorResponse(c)
			return
		}

		ttl := settings.Config.RefreshTokenTTL
		bufferTime := settings.Config.BufferSeconds
		duration, err := durationDiff(ttl, bufferTime)
		if err != nil {
			getInternalServerErrorResponse(c)
			return
		}

		c.SetCookie(
			"refresh_token",     // имя
			tokens.RefreshToken, // значение, например user.RefreshToken
			duration,            // maxAge в секундах
			"/user/refresh",     // path
			"",                  // domain (можно пустое "")
			false,               // secure (true для HTTPS)
			true,                // HttpOnly
		)
		c.JSON(http.StatusOK, dto.UserWithTokenResponse{
			ID:          User.ID,
			Name:        User.Name,
			AccessToken: tokens.AccessToken,
			Message:     "User created successfully",
		})
	}
}

func logoutHandler(us user.UseCase) gin.HandlerFunc {
	//Logout
	return func(c *gin.Context) {
		var r dto.LogoutRequest

		if err := c.ShouldBindJSON(&r); err != nil {
			getBadRequestResponse(c)
			return
		}

		err := us.Logout(r.RefreshToken)

		if err != nil {
			getInternalServerErrorResponse(c)
			return
		}
		c.JSON(http.StatusOK, dto.MessageResponse{
			Message: "User logged out successfully",
		})
	}
}

func logoutAllHandler(us user.UseCase) gin.HandlerFunc {
	//LogoutAll
	return func(c *gin.Context) {
		var r dto.LogoutRequest

		if err := c.ShouldBindJSON(&r); err != nil {
			getBadRequestResponse(c)
			return
		}

		err := us.LogoutAll(r.RefreshToken)

		if err != nil {
			getInternalServerErrorResponse(c)
			return
		}

		c.JSON(http.StatusOK, dto.MessageResponse{
			Message: "User all logged out successfully",
		})
	}
}
func updatePasswordHandler(us user.UseCase) gin.HandlerFunc {
	//Update password
	return func(c *gin.Context) {
		var r dto.UpdatePasswordRequest
		if err := c.ShouldBindJSON(&r); err != nil {
			getBadRequestResponse(c)
			return
		}

		ID, err := getUserIDFromContext(c)
		if err != nil {
			getBadRequestResponse(c)
			return
		}
		updatedUser, err := us.UpdatePassword(ID, r.NewPassword)
		if err != nil {
			getInternalServerErrorResponse(c)
			return
		}
		c.JSON(http.StatusOK, dto.UserResponse{
			ID:      updatedUser.ID,
			Name:    updatedUser.Name,
			Message: "User password updated successfully",
		})
	}
}

func updateEmailHandler(us user.UseCase) gin.HandlerFunc {
	//Update email
	return func(c *gin.Context) {
		var r dto.UpdateEmailRequest
		if err := c.ShouldBindJSON(&r); err != nil {
			getBadRequestResponse(c)
			return
		}
		ID, err := getUserIDFromContext(c)
		if err != nil {
			getBadRequestResponse(c)
			return
		}
		updatedUser, err := us.UpdateEmail(ID, r.NewEmail)
		if err != nil {
			getInternalServerErrorResponse(c)
			return
		}

		c.JSON(http.StatusOK, dto.UserResponse{
			ID:      updatedUser.ID,
			Name:    updatedUser.Name,
			Message: "User email updated successfully",
		})
	}
}

func updatePhoneHandler(us user.UseCase) gin.HandlerFunc {
	//Update phone
	return func(c *gin.Context) {
		var r dto.UpdatePhoneRequest
		if err := c.ShouldBindJSON(&r); err != nil {
			getBadRequestResponse(c)
			return
		}
		ID, err := getUserIDFromContext(c)
		if err != nil {
			getBadRequestResponse(c)
			return
		}
		updatedUser, err := us.UpdatePhone(ID, r.NewPhone)
		if err != nil {
			getInternalServerErrorResponse(c)
			return
		}

		c.JSON(http.StatusOK, dto.UserResponse{
			ID:      updatedUser.ID,
			Name:    updatedUser.Name,
			Message: "User phone updated successfully",
		})
	}
}

func updateProfileHandler(us user.UseCase) gin.HandlerFunc {
	//Update profile
	return func(c *gin.Context) {
		var r dto.UpdateProfileRequest
		if err := c.ShouldBindJSON(&r); err != nil {
			getBadRequestResponse(c)
			return
		}
		ID, err := getUserIDFromContext(c)
		if err != nil {
			getBadRequestResponse(c)
			return
		}
		updatedUser, err := us.UpdateProfile(ID, r)
		if err != nil {
			getInternalServerErrorResponse(c)
			return
		}

		c.JSON(http.StatusOK, dto.UserResponse{
			ID:      updatedUser.ID,
			Name:    updatedUser.Name,
			Message: "User profile updated successfully",
		})
	}
}

func refreshHandler(us user.UseCase) gin.HandlerFunc {
	//Refresh
	return func(c *gin.Context) {
		token, err := c.Cookie("refresh_token")
		if err != nil {
			getBadRequestResponse(c)
			return
		}
		newTokens, err := us.RefreshTokens(token)
		if err != nil {
			getInternalServerErrorResponse(c)
			return
		}
		ttl := settings.Config.RefreshTokenTTL
		bufferTime := settings.Config.BufferSeconds
		duration, err := durationDiff(ttl, bufferTime)
		if err != nil {
			getInternalServerErrorResponse(c)
			return
		}

		c.SetCookie(
			"refresh_token",        // имя
			newTokens.RefreshToken, // значение, например user.RefreshToken
			duration,               // maxAge в секундах
			"/user/refresh",        // path
			"",                     // domain (можно пустое "")
			false,                  // secure (true для HTTPS)
			true,                   // HttpOnly
		)
		c.JSON(http.StatusOK, dto.TokenAccessResponse{
			AccessToken: newTokens.AccessToken,
			Message:     "Refresh token updated successfully",
		})
	}
}

func deleteHandler(us user.UseCase) gin.HandlerFunc {
	//Delete
	return func(c *gin.Context) {
		ID, err := getUserIDFromContext(c)
		if err != nil {
			getBadRequestResponse(c)
			return
		}

		deletedUser, err := us.Delete(ID)
		if err != nil {
			getInternalServerErrorResponse(c)
			return
		}

		c.JSON(http.StatusOK, dto.UserResponse{
			ID:      deletedUser.ID,
			Name:    deletedUser.Name,
			Message: "User deleted successfully",
		})
	}
}

func mapProductsToResponse(products []*order.Item) []orderDTO.ItemsResponse {
	var productsDTO []orderDTO.ItemsResponse
	for _, p := range products {
		productsDTO = append(productsDTO, orderDTO.ItemsResponse{
			Quantity: p.Quantity,
			Product: productDTO.ProductResponse{
				ProductID:    p.Product.ID.String(),
				ProductName:  p.Product.Name,
				ProductImg:   p.Product.Img,
				ProductPrice: p.Product.Price,
			},
		})
	}
	return productsDTO
}

func mapOrdersToResponse(orders []order.Order) []orderDTO.OrderResponse {
	var ordersDTO []orderDTO.OrderResponse
	for _, o := range orders {
		ordersDTO = append(ordersDTO, orderDTO.OrderResponse{
			ID:       o.ID,
			UserID:   o.UserID,
			OrderNum: o.OrderNum,
			Total:    o.Total,
			Status:   o.Status,
			Items:    mapProductsToResponse(o.OrderItems),
		})
	}
	return ordersDTO
}

func ordersHandler(us user.UseCase) gin.HandlerFunc {
	//Orders
	return func(c *gin.Context) {
		ID, err := getUserIDFromContext(c)
		if err != nil {
			getBadRequestResponse(c)
			return
		}

		orders, err := us.Orders(ID)
		if err != nil {
			getInternalServerErrorResponse(c)
			return
		}

		ordersDTO := mapOrdersToResponse(orders)

		c.JSON(http.StatusOK, ordersDTO)
	}
}

func createOrderHandler(us user.UseCase) gin.HandlerFunc {
	//CreateOrder
	return func(c *gin.Context) {
		var r dto.CreateOrderRequest

		if err := c.ShouldBindJSON(&r); err != nil {
			getBadRequestResponse(c)
			return
		}

		ID, err := getUserIDFromContext(c)
		if err != nil {
			getBadRequestResponse(c)
			return
		}

		o, err := us.CreateOrder(ID, r)
		if err != nil {
			getInternalServerErrorResponse(c)
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
		var r dto.DeleteOrderRequest

		if err := c.ShouldBindJSON(&r); err != nil {
			getBadRequestResponse(c)
			return
		}

		ID, err := getUserIDFromContext(c)
		if err != nil {
			getBadRequestResponse(c)
			return
		}

		orderID, err := uuid.Parse(r.OrderID)
		if err != nil {
			getBadRequestResponse(c)
			return
		}

		o, err := us.DeleteOrder(ID, orderID)
		if err != nil {
			getInternalServerErrorResponse(c)
			return
		}

		c.JSON(http.StatusOK, orderDTO.OrderResponse{
			ID:       o.ID,
			UserID:   o.UserID,
			OrderNum: o.OrderNum,
			Total:    o.Total,
			Status:   o.Status,
		})

	}
}
