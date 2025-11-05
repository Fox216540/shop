package api

import (
	shopApiGen "github.com/Fox216540/shop/api/gen"
	aUseCase "github.com/Fox216540/shop/apigateway-service/app/auth"
	cUseCase "github.com/Fox216540/shop/apigateway-service/app/catalog"
	uUseCase "github.com/Fox216540/shop/apigateway-service/app/user"
	"github.com/gin-gonic/gin"
	openapiTypes "github.com/oapi-codegen/runtime/types"
	"net/http"
)

type HTTPHandler struct {
	authUseCase    aUseCase.UseCase
	catalogUseCase cUseCase.UseCase
	userUseCase    uUseCase.UseCase
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
	resp := shopApiGen.UserWithTokenResponse{
		AccessToken: tokens.AccessToken,
		Message:     message,
		Name:        name,
	}
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
	//TODO implement me
	panic("implement me")
}

func (h *HTTPHandler) GetOrders(c *gin.Context) {
	//TODO implement me
	panic("implement me")
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

func (h *HTTPHandler) GetProducts(c *gin.Context, params shopApiGen.GetProductsParams) {
	//TODO implement me
	panic("implement me")
}

func (h *HTTPHandler) GetProductById(c *gin.Context, id openapiTypes.UUID) {
	//TODO implement me
	panic("implement me")
}

func (h *HTTPHandler) CreateUser(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h *HTTPHandler) DeleteUser(c *gin.Context) {
	//TODO implement me
	panic("implement me")
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
