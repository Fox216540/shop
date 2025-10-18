package api

import (
	"context"
	"github.com/Fox216540/shop/catalog-service/app/catalog"
	productDomain "github.com/Fox216540/shop/catalog-service/domain/product"
	pb "github.com/Fox216540/shop/proto/catalog-service/gen"
	types "github.com/Fox216540/shop/proto/common/gen"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCHandler struct {
	catalogUC catalog.UseCase
	pb.UnimplementedInternalCatalogServiceServer
	pb.UnimplementedApiCatalogServiceServer
}

func (h *GRPCHandler) GetCategories(ctx context.Context, req *emptypb.Empty) (*pb.GetCategoriesResponse, error) {
	categories, err := h.catalogUC.GetCategories()
	if err != nil {
		return nil, err
	}

	categoriesResp := make([]*pb.Category, 0, len(categories))
	for _, category := range categories {
		categoriesResp = append(categoriesResp, &pb.Category{
			Id:   category.ID.String(),
			Name: category.Name,
		})
	}

	return &pb.GetCategoriesResponse{
		Categories: categoriesResp,
	}, nil
}

func (h *GRPCHandler) returnProducts(products []productDomain.Product) *types.GetProductsResponse {
	productsResp := make([]*types.Product, 0, len(products))
	for _, product := range products {
		productsResp = append(productsResp, &types.Product{
			Id:          product.ID.String(),
			Name:        product.Name,
			Img:         product.Img,
			Price:       product.Price,
			CategoryId:  product.CategoryID.String(),
			Description: product.Description,
			Stock:       product.Stock,
		})
	}
	return &types.GetProductsResponse{
		Products: productsResp,
	}
}

func (h *GRPCHandler) GetAllProducts(ctx context.Context, req *emptypb.Empty) (*types.GetProductsResponse, error) {
	products, err := h.catalogUC.GetAllProducts()
	if err != nil {
		return nil, err
	}

	return h.returnProducts(products), nil
}

func (h *GRPCHandler) GetProductsOfCategoryId(ctx context.Context, req *pb.GetProductsOfCategoryIdRequest) (*types.GetProductsResponse, error) {
	categoryID, err := uuid.Parse(req.CategoryId)
	if err != nil {
		return nil, err
	}

	products, err := h.catalogUC.GetProductsOfCategoryID(categoryID)
	if err != nil {
		return nil, err
	}
	return h.returnProducts(products), nil
}

func (h *GRPCHandler) GetProductById(ctx context.Context, req *pb.GetProductByIdRequest) (*types.Product, error) {
	productID, err := uuid.Parse(req.ProductId)
	if err != nil {
		return nil, err
	}

	product, err := h.catalogUC.GetProductByID(productID)
	if err != nil {
		return nil, err
	}
	return &types.Product{
		Id:          product.ID.String(),
		Name:        product.Name,
		Img:         product.Img,
		Price:       product.Price,
		CategoryId:  product.CategoryID.String(),
		Description: product.Description,
		Stock:       product.Stock,
	}, nil
}

func (h *GRPCHandler) GetProductsByIds(ctx context.Context, req *pb.GetProductsByIdsRequest) (*types.GetProductsResponse, error) {
	productIDs := make([]uuid.UUID, 0, len(req.ProductIds))
	for _, productID := range req.ProductIds {
		productID, err := uuid.Parse(productID)
		if err != nil {
			return nil, err
		}
		productIDs = append(productIDs, productID)
	}

	products, err := h.catalogUC.GetProductsByIDs(productIDs)
	if err != nil {
		return nil, err
	}
	return h.returnProducts(products), nil
}

func NewGRPCHandler(catalogUC catalog.UseCase) *GRPCHandler {
	return &GRPCHandler{
		catalogUC: catalogUC,
	}
}
