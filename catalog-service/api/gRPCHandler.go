package api

import (
	"context"
	pb "github.com/Fox216540/shop/catalog-service/api/proto"
	"github.com/Fox216540/shop/catalog-service/app/catalog"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

type gRPCHandler struct {
	catalogUC catalog.UseCase
	pb.UnimplementedCatalogServiceServer
}

func (h *gRPCHandler) GetCategories(ctx context.Context, req *emptypb.Empty) (*pb.GetCategoriesResponse, error) {
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

func (h *gRPCHandler) GetAllProducts(ctx context.Context, req *emptypb.Empty) (*pb.GetProductsResponse, error) {
	products, err := h.catalogUC.GetAllProducts()
	if err != nil {
		return nil, err
	}
	productsResp := make([]*pb.Product, 0, len(products))
	for _, product := range products {
		productsResp = append(productsResp, &pb.Product{
			Id:   product.ID.String(),
			Name: product.Name,
		})
	}
	return &pb.GetProductsResponse{
		Products: productsResp,
	}, nil
}

func (h *gRPCHandler) GetProductsOfCategoryId(ctx context.Context, req *pb.GetProductsOfCategoryIdRequest) (*pb.GetProductsResponse, error) {
	categoryID, err := uuid.Parse(req.CategoryId)
	if err != nil {
		return nil, err
	}

	products, err := h.catalogUC.GetProductsOfCategoryID(categoryID)
	if err != nil {
		return nil, err
	}
	productsResp := make([]*pb.Product, 0, len(products))
	for _, product := range products {
		productsResp = append(productsResp, &pb.Product{
			Id:   product.ID.String(),
			Name: product.Name,
		})
	}
	return &pb.GetProductsResponse{
		Products: productsResp,
	}, nil
}

func (h *gRPCHandler) GetProductById(ctx context.Context, req *pb.GetProductByIdRequest) (*pb.Product, error) {
	productID, err := uuid.Parse(req.ProductId)
	if err != nil {
		return nil, err
	}

	product, err := h.catalogUC.GetProductByID(productID)
	if err != nil {
		return nil, err
	}
	return &pb.Product{
		Id:          product.ID.String(),
		Name:        product.Name,
		Img:         product.Img,
		Price:       product.Price,
		CategoryId:  product.CategoryID.String(),
		Description: product.Description,
		Stock:       product.Stock,
	}, nil
}

func (h *gRPCHandler) GetProductsByIds(ctx context.Context, req *pb.GetProductsByIdsRequest) (*pb.GetProductsResponse, error) {
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
	productsResp := make([]*pb.Product, 0, len(products))
	for _, product := range products {
		productsResp = append(productsResp, &pb.Product{
			Id:   product.ID.String(),
			Name: product.Name,
		})
	}
	return &pb.GetProductsResponse{
		Products: productsResp,
	}, nil
}

func NewGRPCHandler(catalogUC catalog.UseCase) pb.CatalogServiceServer {
	return &gRPCHandler{
		catalogUC: catalogUC,
	}
}
