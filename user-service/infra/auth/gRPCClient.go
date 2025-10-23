package auth

import (
	pb "github.com/Fox216540/shop/proto/auth-service/gen/interservice"
	types "github.com/Fox216540/shop/proto/common/gen"
	"github.com/Fox216540/shop/user-service/domain/auth"
	"github.com/Fox216540/shop/user-service/infra/client"
	"github.com/google/uuid"
)

type GRPCClient struct {
	client *client.GRPCClient
	pb     pb.InterserviceServiceClient
}

func (c *GRPCClient) GenerateTokens(userID uuid.UUID) (auth.Tokens, error) {
	ctx := c.client.Context()
	req := &types.UserId{
		Id: userID.String(),
	}
	resp, err := c.pb.GenerateTokens(ctx, req)
	if err != nil {
		return auth.Tokens{}, err
	}
	return auth.Tokens{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
	}, nil
}

func (c *GRPCClient) DeleteAllRefreshTokens(userID uuid.UUID) error {
	ctx := c.client.Context()
	req := &types.UserId{
		Id: userID.String(),
	}
	_, err := c.pb.DeleteAllRefreshTokens(ctx, req)
	return err
}

func NewGRPCClient(client *client.GRPCClient) *GRPCClient {
	return &GRPCClient{
		client: client,
		pb:     pb.NewInterserviceServiceClient(client.Conn()),
	}
}
