package api

import (
	shopApiGen "github.com/Fox216540/shop/api/gen"
	aUseCase "github.com/Fox216540/shop/apigateway-service/app/auth"
	cUseCase "github.com/Fox216540/shop/apigateway-service/app/catalog"
	DTO "github.com/Fox216540/shop/apigateway-service/app/dto"
	uUseCase "github.com/Fox216540/shop/apigateway-service/app/user"
	domainAuth "github.com/Fox216540/shop/apigateway-service/domain/auth"
	domainCatalog "github.com/Fox216540/shop/apigateway-service/domain/catalog"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	openapiTypes "github.com/oapi-codegen/runtime/types"
	"net/http"
)

type HTTPHandler struct {
	authUseCase    aUseCase.UseCase
	catalogUseCase cUseCase.UseCase
	userUseCase    uUseCase.UseCase
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
		//TODO: Придумать ошибку
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
	}
	msg, err := h.authUseCase.LogOut(refresh)
	if err != nil {
		//TODO: Придумать ошибку
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": msg})
}

func (h *HTTPHandler) PostAuthLogoutAll(c *gin.Context) {
	refresh, err := c.Cookie("refresh")
	if err != nil {
		//TODO: Придумать ошибку
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Decode refresh token"})
	}
	msg, err := h.authUseCase.LogOutAll(refresh)
	if err != nil {
		//TODO: Придумать ошибку
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": msg})
}

func (h *HTTPHandler) PostAuthRefresh(c *gin.Context) {
	refresh, err := c.Cookie("refresh")
	if err != nil {
		//TODO: Придумать ошибку
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Decode refresh token"})
	}
	msg, err := h.authUseCase.RefreshTokens(refresh)
	if err != nil {
		//TODO: Придумать ошибку
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": msg})
}

func (h *HTTPHandler) GetCategories(c *gin.Context) {
	categories, err := h.catalogUseCase.GetCategories()
	if err != nil {
		//TODO: Придумать ошибку
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

func (h *HTTPHandler) GetOrders(c *gin.Context) {
	id, exist := c.Get("user_id")
	if !exist {
		//TODO: Придумать ошибку
		return
	}
	orders, err := h.catalogUseCase
}

func (h *HTTPHandler) PostOrders(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *HTTPHandler) DeleteOrdersId(c *gin.Context, id openapiTypes.UUID) {
	//TODO implement me
	panic("implement me")
}

func (h *HTTPHandler) GetOrdersId(c *gin.Context, id openapiTypes.UUID) {
	//TODO implement me
	panic("implement me")
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
		})
	}
	return resp
}

func (h *HTTPHandler) GetProducts(c *gin.Context, params shopApiGen.GetProductsParams) {
	products, err := h.catalogUseCase.GetProducts()
	if err != nil {
		//TODO:
		return
	}
	c.JSON(http.StatusOK, h.productsToResponse(products))
}

func (h *HTTPHandler) GetProductById(c *gin.Context, id openapiTypes.UUID) {
	idString := c.Param("id")
	id, err := uuid.Parse(idString)
	if err != nil {
		//TODO: Придумать ошибку
		return
	}
	products, err := h.catalogUseCase.GetProductsOfCategoryID(id)
	if err != nil {
		//TODO: Придумать ошибку
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
		//TODO: Придумать ошибку
		return
	}
	resp := h.userWithTokenResponse(name, tokens, msg)

	c.JSON(http.StatusCreated, resp)

}

func (h *HTTPHandler) DeleteUser(c *gin.Context) {
	idValue, exists := c.Get("user_id")
	if !exists {
		//TODO: Придумать ошибку
		return
	}
	idString, ok := idValue.(string)
	if !ok {
		//TODO: Придумать ошибку
		return
	}
	id, err := uuid.Parse(idString)
	if err != nil {
		//TODO: Придумать ошибку
		return
	}
	msg, err := h.userUseCase.DeleteUser(id)
	if err != nil {
		//TODO: Придумать ошибку
		return
	}
	c.JSON(http.StatusOK, shopApiGen.MessageResponse{
		Message: msg,
	})
}

func (h *HTTPHandler) PatchUsersMeEmail(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *HTTPHandler) PatchUsersMePassword(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *HTTPHandler) PatchUsersMePhone(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *HTTPHandler) PatchUsersMeProfile(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
