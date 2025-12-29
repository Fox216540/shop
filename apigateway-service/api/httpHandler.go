package api

import (
	"errors"
	shopApiGen "github.com/Fox216540/shop/api/gen"
	aUseCase "github.com/Fox216540/shop/apigateway-service/app/auth"
	cUseCase "github.com/Fox216540/shop/apigateway-service/app/catalog"
	DTO "github.com/Fox216540/shop/apigateway-service/app/dto"
	oUseCase "github.com/Fox216540/shop/apigateway-service/app/order"
	uUseCase "github.com/Fox216540/shop/apigateway-service/app/user"
	"github.com/Fox216540/shop/apigateway-service/core/metrics"
	domainAuth "github.com/Fox216540/shop/apigateway-service/domain/auth"
	domainCatalog "github.com/Fox216540/shop/apigateway-service/domain/catalog"
	domainOrder "github.com/Fox216540/shop/apigateway-service/domain/order"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	openapiTypes "github.com/oapi-codegen/runtime/types"
	"net/http"
	"time"
)

type HTTPHandler struct {
	authUseCase    aUseCase.UseCase
	catalogUseCase cUseCase.UseCase
	userUseCase    uUseCase.UseCase
	orderUseCase   oUseCase.UseCase
	m              HTTPMapper
	metrics        metrics.Metrics
}

func (h *HTTPHandler) userWithTokenResponse(name string, tokens domainAuth.Tokens, message string) shopApiGen.UserWithTokenResponse {
	return shopApiGen.UserWithTokenResponse{
		AccessToken: tokens.AccessToken,
		Message:     message,
		Name:        name,
	}
}

func (h *HTTPHandler) PostAuthLogin(c *gin.Context) {
	var req shopApiGen.PostAuthLoginJSONRequestBody
	if err := c.ShouldBindJSON(&req); err != nil {
		//TODO: Придумать ошибку
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	name, tokens, message, err := h.authUseCase.LogIn(req.PhoneOrEmail, req.Password)
	if err != nil {
		h.metrics.IncLoginFailure()
		c.JSON(h.m.MapError(err))
		return
	}
	resp := h.userWithTokenResponse(name, tokens, message)
	c.JSON(http.StatusOK, resp)
}

func (h *HTTPHandler) PostAuthLogout(c *gin.Context) {
	refresh, err := c.Cookie("refresh")
	if err != nil {
		//TODO: Придумать ошибку
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Decode refresh token"})
		return
	}
	msg, err := h.authUseCase.LogOut(refresh)
	if err != nil {
		c.JSON(h.m.MapError(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": msg})
}

func (h *HTTPHandler) PostAuthLogoutAll(c *gin.Context) {
	refresh, err := c.Cookie("refresh")
	if err != nil {
		//TODO: Придумать ошибку
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Decode refresh token"})
		return
	}
	msg, err := h.authUseCase.LogOutAll(refresh)
	if err != nil {
		c.JSON(h.m.MapError(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": msg})
}

func (h *HTTPHandler) PostAuthRefresh(c *gin.Context) {
	refresh, err := c.Cookie("refresh")
	if err != nil {
		//TODO: Придумать ошибку
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Decode refresh token"})
		return
	}
	msg, err := h.authUseCase.RefreshTokens(refresh)
	if err != nil {
		c.JSON(h.m.MapError(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": msg})
}

func (h *HTTPHandler) GetCategories(c *gin.Context) {
	categories, err := h.catalogUseCase.GetCategories()
	if err != nil {
		c.JSON(h.m.MapError(err))
		return
	}
	resp := make([]shopApiGen.CategoryResponse, 0, len(categories))
	for _, category := range categories {
		resp = append(resp, shopApiGen.CategoryResponse{
			Id:   category.ID,
			Name: category.Name,
		})
	}
	c.JSON(http.StatusOK, resp)

}

func (h *HTTPHandler) getIDOfUser(c *gin.Context) (uuid.UUID, error) {
	idValue, exists := c.Get("user_id")
	if !exists {
		//TODO: Придумать ошибку
		return uuid.Nil, errors.New("user id not found")
	}
	idString, ok := idValue.(string)
	if !ok {
		//TODO: Придумать ошибку
		return uuid.Nil, errors.New("user id not found")
	}
	id, err := uuid.Parse(idString)
	if err != nil {
		//TODO: Придумать ошибку
		return uuid.Nil, errors.New("user id not found")
	}
	return id, nil
}

func (h *HTTPHandler) orderWithItemsResponse(order domainOrder.OrderWithItems) shopApiGen.OrderWithItemsResponse {
	items := make([]shopApiGen.OrderItem, 0, len(order.Items))
	for _, item := range order.Items {
		items = append(items, shopApiGen.OrderItem{
			Product: shopApiGen.ProductShort{
				Currency: item.Product.Currency,
				Id:       item.Product.ID,
				Img:      item.Product.Img,
				Name:     item.Product.Name,
				Price:    item.Product.Price,
			},
			Quantity: item.Quantity,
		})

	}
	return shopApiGen.OrderWithItemsResponse{
		Currency:    order.Order.Currency,
		Id:          order.Order.ID,
		OrderNumber: order.Order.OrderNum,
		Status:      order.Order.Status,
		Total:       order.Order.Total,
		OrderItems:  items,
	}
}

func (h *HTTPHandler) ordersToResponse(orders []domainOrder.Order) []shopApiGen.OrderResponse {
	resp := make([]shopApiGen.OrderResponse, 0, len(orders))
	for _, order := range orders {
		resp = append(resp, shopApiGen.OrderResponse{
			Currency:    order.Currency,
			Id:          order.ID,
			OrderNumber: order.OrderNum,
			Status:      order.Status,
			Total:       order.Total,
		})
	}
	return resp
}

func (h *HTTPHandler) GetOrders(c *gin.Context) {
	userID, err := h.getIDOfUser(c)
	if err != nil {
		//TODO: Придумать ошибку
		return
	}
	orders, err := h.orderUseCase.GetOrders(userID)
	if err != nil {
		c.JSON(h.m.MapError(err))
		return
	}
	c.JSON(http.StatusOK, h.ordersToResponse(orders))
}

func (h *HTTPHandler) PostOrders(c *gin.Context) {
	userID, err := h.getIDOfUser(c)
	if err != nil {
		//TODO: Придумать ошибку
		return
	}
	var req shopApiGen.PostOrdersJSONRequestBody
	if err = c.ShouldBindJSON(&req); err != nil {
		//TODO: Придумать ошибку
		return
	}

	items := make([]domainOrder.ProductRequest, 0, len(req.ProductItems))
	for _, item := range req.ProductItems {
		items = append(items, domainOrder.ProductRequest{
			ID:       item.ProductId,
			Price:    item.Value,
			Quantity: item.Quantity,
			Currency: item.Currency,
		})
	}
	start := time.Now()
	o, err := h.orderUseCase.CreateOrder(userID, items)
	h.metrics.ObserveOrderProcessing(time.Since(start).Seconds())
	if err != nil {
		c.JSON(h.m.MapError(err))
		return
	}
	h.metrics.IncOrder()
	c.JSON(http.StatusCreated, shopApiGen.OrderResponse{
		Currency:    o.Currency,
		Id:          o.ID,
		OrderNumber: o.OrderNum,
		Status:      o.Status,
		Total:       o.Total,
	})
}

func (h *HTTPHandler) DeleteOrdersId(c *gin.Context, id openapiTypes.UUID) {
	userID, err := h.getIDOfUser(c)
	if err != nil {
		//TODO: Придумать ошибку
		return
	}
	deletedID, msg, status, err := h.orderUseCase.DeleteOrder(userID, id)
	if err != nil {
		c.JSON(h.m.MapError(err))
		return
	}
	c.JSON(http.StatusOK, shopApiGen.OrderDeletedResponse{
		Id:      deletedID,
		Message: msg,
		Status:  status,
	})
}

func (h *HTTPHandler) GetOrdersId(c *gin.Context, id openapiTypes.UUID) {
	userID, err := h.getIDOfUser(c)
	if err != nil {
		//TODO: Придумать ошибку
		return
	}
	o, err := h.orderUseCase.GetOrder(userID, id)
	if err != nil {
		c.JSON(h.m.MapError(err))
		return
	}
	c.JSON(http.StatusOK, h.orderWithItemsResponse(o))
}

func (h *HTTPHandler) productsToResponse(products []domainCatalog.Product) []shopApiGen.ProductResponse {
	resp := make([]shopApiGen.ProductResponse, 0, len(products))
	for _, product := range products {
		resp = append(resp, shopApiGen.ProductResponse{
			Id:          product.ID,
			Name:        product.Name,
			Img:         product.Img,
			Price:       product.Price,
			CategoryId:  product.CategoryID,
			Description: product.Description,
			Stock:       product.Stock,
			Currency:    product.Currency,
		})
	}
	return resp
}

func (h *HTTPHandler) GetProducts(c *gin.Context, params shopApiGen.GetProductsParams) {
	var (
		products []domainCatalog.Product
		err      error
	)
	if params.Category != nil {
		products, err = h.catalogUseCase.GetProductsOfCategoryID(*params.Category)
	} else {
		products, err = h.catalogUseCase.GetProducts()
	}
	if err != nil {
		c.JSON(h.m.MapError(err))
		return
	}
	c.JSON(http.StatusOK, h.productsToResponse(products))
}

func (h *HTTPHandler) GetProductById(c *gin.Context, id openapiTypes.UUID) {
	products, err := h.catalogUseCase.GetProductsOfCategoryID(id)
	if err != nil {
		c.JSON(h.m.MapError(err))
		return
	}
	c.JSON(http.StatusOK, h.productsToResponse(products))
}

func (h *HTTPHandler) CreateUser(c *gin.Context) {
	var req shopApiGen.CreateUserJSONRequestBody
	if err := c.ShouldBindJSON(&req); err != nil {
		//TODO: Придумать ошибку
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
		c.JSON(h.m.MapError(err))
		return
	}
	resp := h.userWithTokenResponse(name, tokens, msg)
	h.metrics.IncRegistration()
	c.JSON(http.StatusCreated, resp)

}

func (h *HTTPHandler) DeleteUser(c *gin.Context) {
	userID, err := h.getIDOfUser(c)
	if err != nil {
		//TODO: Придумать ошибку
		return
	}
	msg, err := h.userUseCase.DeleteUser(userID)
	if err != nil {
		c.JSON(h.m.MapError(err))
		return
	}
	c.JSON(http.StatusOK, shopApiGen.MessageResponse{
		Message: msg,
	})
}

func (h *HTTPHandler) PatchUsersMeEmail(c *gin.Context) {
	userID, err := h.getIDOfUser(c)
	if err != nil {
		//TODO: Придумать ошибку
		return
	}
	var req shopApiGen.PatchUsersMeEmailJSONRequestBody
	if err = c.ShouldBindJSON(&req); err != nil {
		//TODO: Придумать ошибку
		return
	}

	msg, err := h.userUseCase.UpdateEmailOfUser(userID, string(req.Email))
	if err != nil {
		c.JSON(h.m.MapError(err))
		return
	}
	c.JSON(http.StatusOK, shopApiGen.MessageResponse{
		Message: msg,
	})
}

func (h *HTTPHandler) PatchUsersMePassword(c *gin.Context) {
	userID, err := h.getIDOfUser(c)
	if err != nil {
		//TODO: Придумать ошибку
		return
	}
	var req shopApiGen.PatchUsersMePasswordJSONRequestBody
	if err = c.ShouldBindJSON(&req); err != nil {
		//TODO: Придумать ошибку
		return
	}

	msg, err := h.userUseCase.UpdatePasswordOfUser(userID, req.Password)
	if err != nil {
		c.JSON(h.m.MapError(err))
		return
	}
	c.JSON(http.StatusOK, shopApiGen.MessageResponse{
		Message: msg,
	})
}

func (h *HTTPHandler) PatchUsersMePhone(c *gin.Context) {
	userID, err := h.getIDOfUser(c)
	if err != nil {
		//TODO: Придумать ошибку
		return
	}
	var req shopApiGen.PatchUsersMePhoneJSONRequestBody
	if err = c.ShouldBindJSON(&req); err != nil {
		//TODO: Придумать ошибку
		return
	}
	msg, err := h.userUseCase.UpdatePhoneOfUser(userID, req.Phone)
	if err != nil {
		c.JSON(h.m.MapError(err))
		return
	}
	c.JSON(http.StatusOK, shopApiGen.MessageResponse{
		Message: msg,
	})
}

func (h *HTTPHandler) PatchUsersMeProfile(c *gin.Context) {
	userID, err := h.getIDOfUser(c)
	if err != nil {
		//TODO: Придумать ошибку
		return
	}
	var req shopApiGen.PatchUsersMeProfileJSONRequestBody
	if err = c.ShouldBindJSON(&req); err != nil {
		//TODO: Придумать ошибку
		return
	}
	msg, name, err := h.userUseCase.UpdateProfileOfUser(userID, req.Name, req.Address)
	if err != nil {
		c.JSON(h.m.MapError(err))
		return
	}
	c.JSON(http.StatusOK, shopApiGen.UserResponse{
		Name:    name,
		Message: msg,
	})
}
