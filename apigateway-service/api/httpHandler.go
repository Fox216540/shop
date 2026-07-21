package api

import (
	"net/http"
	"time"

	shopApiGen "github.com/Fox216540/shop/api/gen"
	aUseCase "github.com/Fox216540/shop/apigateway-service/app/auth"
	bUseCase "github.com/Fox216540/shop/apigateway-service/app/basket"
	cUseCase "github.com/Fox216540/shop/apigateway-service/app/catalog"
	DTO "github.com/Fox216540/shop/apigateway-service/app/dto"
	oUseCase "github.com/Fox216540/shop/apigateway-service/app/order"
	uUseCase "github.com/Fox216540/shop/apigateway-service/app/user"
	"github.com/Fox216540/shop/apigateway-service/core/metrics"
	"github.com/Fox216540/shop/apigateway-service/core/transport"
	domainCatalog "github.com/Fox216540/shop/apigateway-service/domain/catalog"
	domainOrder "github.com/Fox216540/shop/apigateway-service/domain/order"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type HTTPHandler struct {
	authUseCase    aUseCase.UseCase
	basketUseCase  bUseCase.UseCase
	catalogUseCase cUseCase.UseCase
	userUseCase    uUseCase.UseCase
	orderUseCase   oUseCase.UseCase
	m              *HTTPMapper
	metrics        metrics.Metrics
	refreshMaxAge  int
}

func NewHTTPHandler(
	authUseCase aUseCase.UseCase,
	basketUseCase bUseCase.UseCase,
	catalogUseCase cUseCase.UseCase,
	userUseCase uUseCase.UseCase,
	orderUseCase oUseCase.UseCase,
	mapper *HTTPMapper,
	metrics metrics.Metrics,
	refreshMaxAge int,
) *HTTPHandler {
	return &HTTPHandler{
		authUseCase:    authUseCase,
		basketUseCase:  basketUseCase,
		catalogUseCase: catalogUseCase,
		userUseCase:    userUseCase,
		orderUseCase:   orderUseCase,
		m:              mapper,
		metrics:        metrics,
		refreshMaxAge:  refreshMaxAge,
	}
}

func (h *HTTPHandler) writeMappedError(c *gin.Context, err error) {
	c.JSON(h.m.MapError(c.Request.Context(), err))
}

func (h *HTTPHandler) writeMessage(c *gin.Context, status int, message string) {
	c.JSON(status, shopApiGen.MessageResponse{Message: message})
}

func (h *HTTPHandler) bindJSON(c *gin.Context, req any) bool {
	if err := c.ShouldBindJSON(req); err != nil {
		h.writeMessage(c, http.StatusBadRequest, messages.InvalidJSON)
		return false
	}
	return true
}

func (h *HTTPHandler) refreshToken(c *gin.Context) (string, bool) {
	refresh, err := c.Cookie(messages.RefreshCookieName)
	if err != nil {
		h.writeMessage(c, http.StatusUnauthorized, messages.MissingRefreshCookie)
		return "", false
	}
	return refresh, true
}

func (h *HTTPHandler) bearerToken(c *gin.Context) (string, bool) {
	authHeader := c.GetHeader(transport.AuthorizationHeader)
	prefix := transport.BearerPrefix
	if len(authHeader) <= len(prefix) || authHeader[:len(prefix)] != prefix {
		h.writeMessage(c, http.StatusUnauthorized, messages.Unauthorized)
		return "", false
	}
	return authHeader[len(prefix):], true
}

func (h *HTTPHandler) setRefreshCookie(c *gin.Context, token string) {
	http.SetCookie(c.Writer, refreshCookie(token, h.refreshMaxAge))
}

func (h *HTTPHandler) clearRefreshCookie(c *gin.Context) {
	http.SetCookie(c.Writer, expiredRefreshCookie())
}

func (h *HTTPHandler) Login(c *gin.Context) {
	var req shopApiGen.LoginJSONRequestBody
	if !h.bindJSON(c, &req) {
		return
	}

	name, tokens, message, err := h.authUseCase.LogIn(req.PhoneOrEmail, req.Password)
	if err != nil {
		h.metrics.IncLoginFailure()
		h.writeMappedError(c, err)
		return
	}

	h.setRefreshCookie(c, tokens.RefreshToken)
	c.JSON(http.StatusOK, userWithTokenResponse(name, tokens, message))
}

func (h *HTTPHandler) Logout(c *gin.Context) {
	refresh, ok := h.refreshToken(c)
	if !ok {
		return
	}

	msg, err := h.authUseCase.LogOut(refresh)
	if err != nil {
		h.writeMappedError(c, err)
		return
	}

	h.clearRefreshCookie(c)
	h.writeMessage(c, http.StatusOK, msg)
}

func (h *HTTPHandler) LogoutAll(c *gin.Context) {
	token, ok := h.bearerToken(c)
	if !ok {
		return
	}

	msg, err := h.authUseCase.LogOutAll(token)
	if err != nil {
		h.writeMappedError(c, err)
		return
	}

	h.clearRefreshCookie(c)
	h.writeMessage(c, http.StatusOK, msg)
}

func (h *HTTPHandler) RefreshToken(c *gin.Context) {
	refresh, ok := h.refreshToken(c)
	if !ok {
		return
	}

	tokens, err := h.authUseCase.RefreshTokens(refresh)
	if err != nil {
		h.writeMappedError(c, err)
		return
	}

	h.setRefreshCookie(c, tokens.RefreshToken)
	c.JSON(http.StatusOK, shopApiGen.RefreshResponse{
		AccessToken: tokens.AccessToken,
		Message:     messages.RefreshSuccess,
	})
}

func (h *HTTPHandler) GetCategories(c *gin.Context) {
	categories, err := h.catalogUseCase.GetCategories()
	if err != nil {
		h.writeMappedError(c, err)
		return
	}

	c.JSON(http.StatusOK, categoriesToResponse(categories))
}

func (h *HTTPHandler) getIDOfUser(c *gin.Context) (uuid.UUID, error) {
	idValue, exists := c.Get(transport.UserIDKey)
	if !exists {
		return uuid.Nil, NewMissingUserIDError(nil)
	}

	switch id := idValue.(type) {
	case uuid.UUID:
		return id, nil
	case string:
		parsed, err := uuid.Parse(id)
		if err != nil {
			return uuid.Nil, NewMissingUserIDError(err)
		}
		return parsed, nil
	default:
		return uuid.Nil, NewMissingUserIDError(nil)
	}
}

func (h *HTTPHandler) getAuthenticatedUserID(c *gin.Context) (uuid.UUID, bool) {
	userID, err := h.getIDOfUser(c)
	if err != nil {
		h.writeMessage(c, http.StatusUnauthorized, messages.Unauthorized)
		return uuid.Nil, false
	}
	return userID, true
}

func (h *HTTPHandler) GetOrders(c *gin.Context) {
	userID, ok := h.getAuthenticatedUserID(c)
	if !ok {
		return
	}

	orders, err := h.orderUseCase.GetOrders(userID)
	if err != nil {
		h.writeMappedError(c, err)
		return
	}
	c.JSON(http.StatusOK, ordersToResponse(orders))
}

func (h *HTTPHandler) CreateOrder(c *gin.Context) {
	userID, ok := h.getAuthenticatedUserID(c)
	if !ok {
		return
	}

	var req shopApiGen.CreateOrderJSONRequestBody
	if !h.bindJSON(c, &req) {
		return
	}

	items := make([]domainOrder.ProductRequest, 0, len(req.Items))
	for _, item := range req.Items {
		if item.Quantity < 1 {
			h.writeMessage(c, http.StatusBadRequest, messages.InvalidArgument)
			return
		}
		items = append(items, domainOrder.ProductRequest{
			ID:       item.ProductId,
			Price:    item.ExpectedPrice.Amount,
			Quantity: uint64(item.Quantity),
			Currency: item.ExpectedPrice.Currency,
		})
	}

	start := time.Now()
	o, err := h.orderUseCase.CreateOrder(userID, items)
	h.metrics.ObserveOrderProcessing(time.Since(start).Seconds())
	if err != nil {
		h.writeMappedError(c, err)
		return
	}

	h.metrics.IncOrder()
	c.JSON(http.StatusCreated, orderToResponse(o))
}

func (h *HTTPHandler) DeleteOrder(c *gin.Context, orderID shopApiGen.OrderId) {
	userID, ok := h.getAuthenticatedUserID(c)
	if !ok {
		return
	}

	deletedID, msg, status, err := h.orderUseCase.DeleteOrder(userID, orderID)
	if err != nil {
		h.writeMappedError(c, err)
		return
	}

	c.JSON(http.StatusOK, orderDeletedResponse(deletedID, msg, status))
}

func (h *HTTPHandler) GetOrderById(c *gin.Context, orderID shopApiGen.OrderId) {
	userID, ok := h.getAuthenticatedUserID(c)
	if !ok {
		return
	}

	o, err := h.orderUseCase.GetOrder(userID, orderID)
	if err != nil {
		h.writeMappedError(c, err)
		return
	}
	c.JSON(http.StatusOK, orderWithItemsToResponse(o))
}

func (h *HTTPHandler) GetProducts(c *gin.Context, params shopApiGen.GetProductsParams) {
	var (
		products []domainCatalog.Product
		err      error
	)
	if params.CategoryId != nil {
		products, err = h.catalogUseCase.GetProductsOfCategoryID(*params.CategoryId)
	} else {
		products, err = h.catalogUseCase.GetProducts()
	}
	if err != nil {
		h.writeMappedError(c, err)
		return
	}
	c.JSON(http.StatusOK, productsToResponse(products))
}

func (h *HTTPHandler) GetProductById(c *gin.Context, productID shopApiGen.ProductId, _ shopApiGen.GetProductByIdParams) {
	product, err := h.catalogUseCase.GetProductByID(productID)
	if err != nil {
		h.writeMappedError(c, err)
		return
	}
	c.JSON(http.StatusOK, productToResponse(product))
}

func (h *HTTPHandler) CreateUser(c *gin.Context) {
	var req shopApiGen.CreateUserJSONRequestBody
	if !h.bindJSON(c, &req) {
		return
	}

	userDTO := DTO.User{
		Name:     req.Name,
		Email:    string(req.Email),
		Password: req.Password,
		Phone:    req.Phone,
		Address:  req.Address,
	}
	name, tokens, msg, err := h.userUseCase.RegisterUser(userDTO)
	if err != nil {
		h.writeMappedError(c, err)
		return
	}

	h.setRefreshCookie(c, tokens.RefreshToken)
	h.metrics.IncRegistration()
	c.JSON(http.StatusCreated, userWithTokenResponse(name, tokens, msg))
}

func (h *HTTPHandler) DeleteUser(c *gin.Context) {
	userID, ok := h.getAuthenticatedUserID(c)
	if !ok {
		return
	}

	if _, err := h.userUseCase.DeleteUser(userID); err != nil {
		h.writeMappedError(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *HTTPHandler) UpdateUserEmail(c *gin.Context) {
	userID, ok := h.getAuthenticatedUserID(c)
	if !ok {
		return
	}

	var req shopApiGen.UpdateUserEmailJSONRequestBody
	if !h.bindJSON(c, &req) {
		return
	}

	msg, err := h.userUseCase.UpdateEmailOfUser(userID, string(req.Email))
	if err != nil {
		h.writeMappedError(c, err)
		return
	}
	h.writeMessage(c, http.StatusOK, msg)
}

func (h *HTTPHandler) UpdateUserPassword(c *gin.Context) {
	userID, ok := h.getAuthenticatedUserID(c)
	if !ok {
		return
	}

	var req shopApiGen.UpdateUserPasswordJSONRequestBody
	if !h.bindJSON(c, &req) {
		return
	}

	msg, err := h.userUseCase.UpdatePasswordOfUser(userID, req.Password)
	if err != nil {
		h.writeMappedError(c, err)
		return
	}
	h.writeMessage(c, http.StatusOK, msg)
}

func (h *HTTPHandler) UpdateUserPhone(c *gin.Context) {
	userID, ok := h.getAuthenticatedUserID(c)
	if !ok {
		return
	}

	var req shopApiGen.UpdateUserPhoneJSONRequestBody
	if !h.bindJSON(c, &req) {
		return
	}

	msg, err := h.userUseCase.UpdatePhoneOfUser(userID, req.Phone)
	if err != nil {
		h.writeMappedError(c, err)
		return
	}
	h.writeMessage(c, http.StatusOK, msg)
}

func (h *HTTPHandler) UpdateUserProfile(c *gin.Context) {
	userID, ok := h.getAuthenticatedUserID(c)
	if !ok {
		return
	}

	var req shopApiGen.UpdateUserProfileJSONRequestBody
	if !h.bindJSON(c, &req) {
		return
	}

	msg, name, err := h.userUseCase.UpdateProfileOfUser(userID, req.Name, req.Address)
	if err != nil {
		h.writeMappedError(c, err)
		return
	}
	c.JSON(http.StatusOK, userResponse(name, msg))
}

func (h *HTTPHandler) GetBasket(c *gin.Context, _ shopApiGen.GetBasketParams) {
	userID, ok := h.getAuthenticatedUserID(c)
	if !ok {
		return
	}

	basket, err := h.basketUseCase.GetBasket(userID)
	if err != nil {
		h.writeMappedError(c, err)
		return
	}

	c.JSON(http.StatusOK, basketToResponse(basket))
}

func (h *HTTPHandler) AddItemToBasket(c *gin.Context) {
	userID, ok := h.getAuthenticatedUserID(c)
	if !ok {
		return
	}

	var req shopApiGen.AddItemToBasketJSONRequestBody
	if !h.bindJSON(c, &req) {
		return
	}
	if req.Quantity < 1 {
		h.writeMessage(c, http.StatusBadRequest, messages.InvalidArgument)
		return
	}

	item, err := h.basketUseCase.AddItemToBasket(userID, req.ProductId, uint64(req.Quantity))
	if err != nil {
		h.writeMappedError(c, err)
		return
	}
	c.JSON(http.StatusCreated, basketItemResponse(item))
}

func (h *HTTPHandler) ClearBasket(c *gin.Context) {
	userID, ok := h.getAuthenticatedUserID(c)
	if !ok {
		return
	}

	if err := h.basketUseCase.DeleteBasket(userID); err != nil {
		h.writeMappedError(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *HTTPHandler) RemoveItemFromBasket(c *gin.Context, productID shopApiGen.ProductId) {
	userID, ok := h.getAuthenticatedUserID(c)
	if !ok {
		return
	}

	msg, err := h.basketUseCase.RemoveItemFromBasket(userID, productID)
	if err != nil {
		h.writeMappedError(c, err)
		return
	}
	h.writeMessage(c, http.StatusOK, msg)
}
