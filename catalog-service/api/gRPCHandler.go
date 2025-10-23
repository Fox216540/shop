package api

import (
	"context"
	"github.com/Fox216540/shop/catalog-service/app/catalog"
	productDomain "github.com/Fox216540/shop/catalog-service/domain/product"
	pbApi "github.com/Fox216540/shop/proto/catalog-service/gen/api"
	pbInterservice "github.com/Fox216540/shop/proto/catalog-service/gen/interservice"
	types "github.com/Fox216540/shop/proto/common/gen"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCHandler struct {
	catalogUC catalog.UseCase
	pbInterservice.UnimplementedInterserviceServiceServer
	pbApi.UnimplementedApiServiceServer
}

func productToPbWithSupplement(
	p productDomain.Product,
) *pbApi.ProductWithSupplement {
	return &pbApi.ProductWithSupplement{
		Product: &types.Product{
			Id:    p.ID.String(),
			Name:  p.Name,
			Img:   p.Img,
			Price: p.Price,
		},
		CategoryId:  p.CategoryID.String(),
		Description: p.Description,
		Stock:       p.Stock,
	}
}

func (h *GRPCHandler) productsToPbWithSupplement(
	products []productDomain.Product,
) *pbApi.GetProductsResponse {
	productsResp := make([]*pbApi.ProductWithSupplement, 0, len(products))
	for _, product := range products {
		productsResp = append(productsResp, productToPbWithSupplement(product))
	}
	return &pbApi.GetProductsResponse{
		Products: productsResp,
	}
}

func (h *GRPCHandler) productsToPb(
	products []productDomain.Product,
) []*types.Product {
	productsResp := make([]*types.Product, 0, len(products))
	for _, product := range products {
		productsResp = append(productsResp, &types.Product{
			Id:    product.ID.String(),
			Name:  product.Name,
			Img:   product.Img,
			Price: product.Price,
		})
	}
	return productsResp
}

func (h *GRPCHandler) GetCategories(
	ctx context.Context,
	req *emptypb.Empty,
) (*pbApi.GetCategoriesResponse, error) {
	categories, err := h.catalogUC.GetCategories()
	if err != nil {
		return nil, err
	}

	categoriesResp := make([]*pbApi.Category, 0, len(categories))
	for _, category := range categories {
		categoriesResp = append(categoriesResp, &pbApi.Category{
			Id:   category.ID.String(),
			Name: category.Name,
		})
	}

	return &pbApi.GetCategoriesResponse{
		Categories: categoriesResp,
	}, nil
}

func (h *GRPCHandler) GetAllProducts(
	ctx context.Context,
	req *emptypb.Empty,
) (*pbApi.GetProductsResponse, error) {
	products, err := h.catalogUC.GetAllProducts()
	if err != nil {
		return nil, err
	}

	return h.productsToPbWithSupplement(products), nil
}

func (h *GRPCHandler) GetProductsOfCategoryId(
	ctx context.Context,
	req *pbApi.GetProductsOfCategoryIdRequest,
) (*pbApi.GetProductsResponse, error) {
	categoryID, err := uuid.Parse(req.CategoryId)
	if err != nil {
		return nil, err
	}

	products, err := h.catalogUC.GetProductsOfCategoryID(categoryID)
	if err != nil {
		return nil, err
	}
	return h.productsToPbWithSupplement(products), nil
}

func (h *GRPCHandler) GetProductById(
	ctx context.Context,
	req *pbApi.GetProductByIdRequest,
) (*pbApi.ProductWithSupplement, error) {
	productID, err := uuid.Parse(req.ProductId)
	if err != nil {
		return nil, err
	}

	product, err := h.catalogUC.GetProductByID(productID)
	if err != nil {
		return nil, err
	}
	return productToPbWithSupplement(product), nil
}

func (h *GRPCHandler) GetProductsByIds(
	ctx context.Context,
	req *pbInterservice.GetProductsByIdsRequest,
) (*pbInterservice.GetProductsByIdsResponse, error) {
	productIDs := make([]uuid.UUID, 0, len(req.ProductIds))
	for _, productID := range req.ProductIds {
		id, err := uuid.Parse(productID)
		if err != nil {
			return nil, err
		}
		productIDs = append(productIDs, id)
	}

	products, err := h.catalogUC.GetProductsByIDs(productIDs)
	if err != nil {
		return nil, err
	}
	return &pbInterservice.GetProductsByIdsResponse{
		Products: h.productsToPb(products),
	}, nil
}

func NewGRPCHandler(catalogUC catalog.UseCase) *GRPCHandler {
	return &GRPCHandler{
		catalogUC: catalogUC,
	}
}
