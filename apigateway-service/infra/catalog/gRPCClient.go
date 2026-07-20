package catalog

import (
	"github.com/Fox216540/shop/apigateway-service/domain/catalog"
	"github.com/Fox216540/shop/apigateway-service/infra/client"
	pb "github.com/Fox216540/shop/proto/catalog-service/gen/api"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCClient struct {
	conn *client.GRPCClient
	pb   pb.ApiServiceClient
}

func NewGRPCClient(client *client.GRPCClient) *GRPCClient {
	return &GRPCClient{
		conn: client,
		pb:   pb.NewApiServiceClient(client.Conn()), // Инициализация клиента
	}
}

func (c *GRPCClient) GetAllCategories() ([]catalog.Category, error) {
	ctx := c.conn.Context()

	resp, err := c.pb.GetCategories(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, nil
	}

	var categories []catalog.Category
	for _, category := range resp.Categories {
		categoryID, err := uuid.Parse(category.Id)
		if err != nil {
			return nil, err
		}
		categories = append(categories, catalog.Category{
			ID:   categoryID,
			Name: category.Name,
		})
	}
	return categories, nil
}

func (c *GRPCClient) GetAllProducts() ([]catalog.Product, error) {
	ctx := c.conn.Context()

	resp, err := c.pb.GetAllProducts(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, nil
	}

	var products []catalog.Product
	for _, product := range resp.Products {
		if product == nil || product.Product == nil {
			continue
		}
		productID, err := uuid.Parse(product.Product.Id)
		if err != nil {
			return nil, err
		}
		categoryID, err := uuid.Parse(product.CategoryId)
		if err != nil {
			return nil, err
		}
		price := ""
		currency := ""
		if product.Product.GetPrice() != nil {
			price = product.Product.GetPrice().GetAmount()
			currency = product.Product.GetPrice().GetCurrency()
		}
		products = append(products, catalog.Product{
			ID:          productID,
			Name:        product.Product.Name,
			Img:         product.Product.Img,
			Price:       price,
			Currency:    currency,
			CategoryID:  categoryID,
			Description: product.Description,
			Stock:       product.Stock,
		})
	}
	return products, nil
}

func (c *GRPCClient) GetProductsOfCategoryID(ID uuid.UUID) ([]catalog.Product, error) {
	ctx := c.conn.Context()

	req := &pb.GetProductsOfCategoryIdRequest{
		CategoryId: ID.String(),
	}

	resp, err := c.pb.GetProductsOfCategoryId(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, nil
	}

	var products []catalog.Product
	for _, product := range resp.Products {
		if product == nil || product.Product == nil {
			continue
		}
		productID, err := uuid.Parse(product.Product.Id)
		if err != nil {
			return nil, err
		}
		categoryID, err := uuid.Parse(product.CategoryId)
		if err != nil {
			return nil, err
		}
		price := ""
		currency := ""
		if product.Product.GetPrice() != nil {
			price = product.Product.GetPrice().GetAmount()
			currency = product.Product.GetPrice().GetCurrency()
		}
		products = append(products, catalog.Product{
			ID:          productID,
			Name:        product.Product.Name,
			Img:         product.Product.Img,
			Price:       price,
			Currency:    currency,
			CategoryID:  categoryID,
			Description: product.Description,
			Stock:       product.Stock,
		})
	}
	return products, nil
}

func (c *GRPCClient) GetProductByID(ID uuid.UUID) (catalog.Product, error) {
	ctx := c.conn.Context()

	req := &pb.GetProductByIdRequest{
		ProductId: ID.String(),
	}

	resp, err := c.pb.GetProductById(ctx, req)
	if err != nil {
		return catalog.Product{}, err
	}
	if resp == nil || resp.Product == nil {
		return catalog.Product{}, nil
	}

	productID, err := uuid.Parse(resp.Product.Id)
	if err != nil {
		return catalog.Product{}, err
	}
	categoryID, err := uuid.Parse(resp.CategoryId)
	if err != nil {
		return catalog.Product{}, err
	}
	price := ""
	currency := ""
	if resp.Product.GetPrice() != nil {
		price = resp.Product.GetPrice().GetAmount()
		currency = resp.Product.GetPrice().GetCurrency()
	}
	return catalog.Product{
		ID:          productID,
		Name:        resp.Product.Name,
		Img:         resp.Product.Img,
		Price:       price,
		Currency:    currency,
		CategoryID:  categoryID,
		Description: resp.Description,
		Stock:       resp.Stock,
	}, nil
}
