package api

import (
	shopApiGen "github.com/Fox216540/shop/api/gen"
	openapiTypes "github.com/oapi-codegen/runtime/types"
	"net/http"
)

type HTTPHandler struct {
}

func (h *HTTPHandler) PostAuthLogin(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h *HTTPHandler) PostAuthLogout(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h *HTTPHandler) PostAuthLogoutAll(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h *HTTPHandler) PostAuthRefresh(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h *HTTPHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h *HTTPHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h *HTTPHandler) PostOrders(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h *HTTPHandler) DeleteOrdersId(w http.ResponseWriter, r *http.Request, id openapiTypes.UUID) {
	//TODO implement me
	panic("implement me")
}

func (h *HTTPHandler) GetOrdersId(w http.ResponseWriter, r *http.Request, id openapiTypes.UUID) {
	//TODO implement me
	panic("implement me")
}

func (h *HTTPHandler) GetProducts(w http.ResponseWriter, r *http.Request, params shopApiGen.GetProductsParams) {
	//TODO implement me
	panic("implement me")
}

func (h *HTTPHandler) GetProductById(w http.ResponseWriter, r *http.Request, id openapiTypes.UUID) {
	//TODO implement me
	panic("implement me")
}

func (h *HTTPHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h *HTTPHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h *HTTPHandler) PatchUsersMeEmail(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h *HTTPHandler) PatchUsersMePassword(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h *HTTPHandler) PatchUsersMePhone(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h *HTTPHandler) PatchUsersMeProfile(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
