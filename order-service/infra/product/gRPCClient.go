package product

import (
	"github.com/Fox216540/shop/order-service/domain/product"
	"github.com/Fox216540/shop/order-service/infra/client"
	pb "github.com/Fox216540/shop/proto/catalog-service/gen"
	"github.com/google/uuid"
)

type GRPCClient struct {
	client *client.GRPCClient
	pb     pb.InternalCatalogServiceClient
}

func NewGRPCClient(client *client.GRPCClient) *GRPCClient {
	return &GRPCClient{
		client: client,
		pb:     pb.NewInternalCatalogServiceClient(client.Conn()), // Инициализация клиента
	}
}

func (r *GRPCClient) GetProductsByIDs(ids []uuid.UUID) ([]product.Product, error) {
	idsStrings := make([]string, 0, len(ids))
	for _, id := range ids {
		idsStrings = append(idsStrings, id.String())
	}

	req := &pb.GetProductsByIdsRequest{
		ProductIds: idsStrings,
	}

	// Используем контекст из клиента или создаем новый
	ctx := r.client.Context() // если есть такой метод

	resp, err := r.pb.GetProductsByIds(ctx, req)
	if err != nil {
		return nil, err
	}

	// Конвертация protobuf сообщений в доменные объекты
	products := make([]product.Product, 0, len(resp.Products))
	for _, pbProduct := range resp.Products {
		productID, err := uuid.Parse(pbProduct.Id)
		if err != nil {
			return nil, err
		}

		domainProduct := product.Product{
			ID:    productID,
			Name:  pbProduct.Name,
			Img:   pbProduct.Img,
			Price: pbProduct.Price,
		}
		products = append(products, domainProduct)
	}

	return products, nil
}
