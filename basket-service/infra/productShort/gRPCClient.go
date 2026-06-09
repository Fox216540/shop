package productShort

import (
	"github.com/Fox216540/shop/basket-service/domain/money"
	domainproduct "github.com/Fox216540/shop/basket-service/domain/productShort"
	"github.com/Fox216540/shop/basket-service/infra/client"
	pb "github.com/Fox216540/shop/proto/catalog-service/gen/interservice"
	types "github.com/Fox216540/shop/proto/common/gen"
	"github.com/google/uuid"
)

type GRPCClient struct {
	client *client.GRPCClient
	pb     pb.InterserviceServiceClient
}

func NewGRPCClient(client *client.GRPCClient) *GRPCClient {
	return &GRPCClient{
		client: client,
		pb:     pb.NewInterserviceServiceClient(client.Conn()),
	}
}

func (c *GRPCClient) GetProductByID(id uuid.UUID) (domainproduct.ProductShort, error) {
	ctx := c.client.Context()
	req := &pb.GetProductsByIdsRequest{
		ProductIds: []string{id.String()},
	}

	resp, err := c.pb.GetProductsByIds(ctx, req)
	if err != nil {
		return domainproduct.ProductShort{}, err
	}

	if len(resp.Products) == 0 {
		return domainproduct.ProductShort{}, domainproduct.NewNotFoundError(nil)
	}

	return fromProtoProduct(resp.Products[0])
}

func fromProtoProduct(p *types.Product) (domainproduct.ProductShort, error) {
	if p == nil || p.Price == nil {
		return domainproduct.ProductShort{}, domainproduct.NewNotFoundError(nil)
	}

	productID, err := uuid.Parse(p.Id)
	if err != nil {
		return domainproduct.ProductShort{}, domainproduct.NewInvalidUUIDError(err)
	}

	return domainproduct.ProductShort{
		ID:   productID,
		Name: p.Name,
		Img:  p.Img,
		Price: money.Money{
			Amount:   p.Price.GetAmount(),
			Currency: p.Price.GetCurrency(),
		},
	}, nil
}
