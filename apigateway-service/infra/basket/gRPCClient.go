package basket

import (
	"github.com/Fox216540/shop/apigateway-service/core/transport"
	"github.com/Fox216540/shop/apigateway-service/domain/basket"
	"github.com/Fox216540/shop/apigateway-service/infra/client"
	pb "github.com/Fox216540/shop/proto/basket-service/gen/api"
	common "github.com/Fox216540/shop/proto/common/gen"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCClient struct {
	conn     *client.GRPCClient
	pbClient pb.ApiServiceClient
}

func NewGRPCClient(client *client.GRPCClient) *GRPCClient {
	return &GRPCClient{
		conn:     client,
		pbClient: pb.NewApiServiceClient(client.Conn()),
	}
}

func (c *GRPCClient) contextWithUserID(userID uuid.UUID) metadata.MD {
	return metadata.Pairs(transport.UserIDKey, userID.String())
}

func mapMoney(price *common.Money) basket.Money {
	if price == nil {
		return basket.Money{}
	}
	return basket.Money{
		Amount:   price.GetAmount(),
		Currency: price.GetCurrency(),
	}
}

func mapItem(item *pb.BasketItem) (basket.Item, error) {
	if item == nil || item.Product == nil {
		return basket.Item{}, NewBasketItemNilError(nil)
	}
	productID, err := uuid.Parse(item.Product.Id)
	if err != nil {
		return basket.Item{}, err
	}

	price := item.Product.GetPrice()
	product := basket.Product{
		ID:   productID,
		Name: item.Product.Name,
		Img:  item.Product.Img,
	}
	if price != nil {
		product.Price = price.GetAmount()
		product.Currency = price.GetCurrency()
	}

	return basket.Item{
		Product:  product,
		Quantity: item.Quantity,
	}, nil
}

func (c *GRPCClient) GetBasket(userID uuid.UUID) (basket.Basket, error) {
	ctx := metadata.NewOutgoingContext(c.conn.Context(), c.contextWithUserID(userID))
	resp, err := c.pbClient.GetBasket(ctx, &emptypb.Empty{})
	if err != nil {
		return basket.Basket{}, err
	}
	if resp == nil {
		return basket.Basket{}, NewBasketResponseNilError(nil)
	}

	items := make([]basket.Item, 0, len(resp.Items))
	for _, item := range resp.Items {
		mapped, err := mapItem(item)
		if err != nil {
			return basket.Basket{}, err
		}
		items = append(items, mapped)
	}

	return basket.Basket{
		Items: items,
		Total: mapMoney(resp.GetTotal()),
	}, nil
}

func (c *GRPCClient) AddItemToBasket(userID, productID uuid.UUID, quantity uint64) (basket.Item, error) {
	ctx := metadata.NewOutgoingContext(c.conn.Context(), c.contextWithUserID(userID))
	resp, err := c.pbClient.AddItemToBasket(ctx, &pb.AddItemToBasketRequest{
		ProductId: productID.String(),
		Quantity:  quantity,
	})
	if err != nil {
		return basket.Item{}, err
	}
	return mapItem(resp.GetItem())
}

func (c *GRPCClient) DeleteBasket(userID uuid.UUID) error {
	ctx := metadata.NewOutgoingContext(c.conn.Context(), c.contextWithUserID(userID))
	_, err := c.pbClient.DeleteBasket(ctx, &emptypb.Empty{})
	return err
}

func (c *GRPCClient) RemoveItemFromBasket(userID, productID uuid.UUID) (string, error) {
	ctx := metadata.NewOutgoingContext(c.conn.Context(), c.contextWithUserID(userID))
	resp, err := c.pbClient.RemoveItemFromBasket(ctx, &pb.RemoveItemFromBasketRequest{
		ProductId: productID.String(),
	})
	if err != nil {
		return "", err
	}
	if resp == nil {
		return "", NewRemoveBasketItemResponseNilError(nil)
	}
	return resp.Message, nil
}
