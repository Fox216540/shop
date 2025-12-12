package catalog

import (
	"github.com/Fox216540/shop/order-service/domain/catalog"
	"github.com/Fox216540/shop/order-service/infra/client"
	pb "github.com/Fox216540/shop/proto/catalog-service/gen/interservice"
	"github.com/google/uuid"
	"google.golang.org/grpc/status"
)

type GRPCClient struct {
	client *client.GRPCClient
	pb     pb.InterserviceServiceClient
}

func (c *GRPCClient) GetProductsByIDs(ids []uuid.UUID) ([]catalog.Product, error) {
	idsStrings := make([]string, 0, len(ids))
	for _, id := range ids {
		idsStrings = append(idsStrings, id.String())
	}

	req := &pb.GetProductsByIdsRequest{
		ProductIds: idsStrings,
	}

	ctx := c.client.Context()

	resp, err := c.pb.GetProductsByIds(ctx, req)
	if err != nil {
		if _, ok := status.FromError(err); ok {
			// это gRPC-ошибка — вернуть как есть
			return nil, err
		}

		// не gRPC — завернуть
		return nil, NewGRPCError(err)
	}

	// Конвертация protobuf сообщений в доменные объекты
	products := make([]catalog.Product, 0, len(resp.Products))
	for _, pbProduct := range resp.Products {
		productID, err := uuid.Parse(pbProduct.Id)
		if err != nil {
			return nil, NewInvalidUUID(err)
		}

		domainProduct := catalog.Product{
			ID:    productID,
			Name:  pbProduct.Name,
			Img:   pbProduct.Img,
			Price: pbProduct.Price,
		}
		products = append(products, domainProduct)
	}

	return products, nil
}

func NewGRPCClient(client *client.GRPCClient) *GRPCClient {
	return &GRPCClient{
		client: client,
		pb:     pb.NewInterserviceServiceClient(client.Conn()), // Инициализация клиента
	}
}
