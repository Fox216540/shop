package auth

import (
	pb "github.com/Fox216540/shop/proto/auth-service/gen/interservice"
	types "github.com/Fox216540/shop/proto/common/gen"
	"github.com/Fox216540/shop/user-service/domain/auth"
	"github.com/Fox216540/shop/user-service/infra/client"
	"github.com/google/uuid"
	"google.golang.org/grpc/status"
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
		if _, ok := status.FromError(err); ok {
			// это gRPC-ошибка — вернуть как есть
			return auth.Tokens{}, err
		}

		// не gRPC — завернуть
		return auth.Tokens{}, NewGRPCError(err)
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
	if _, ok := status.FromError(err); ok {
		// это gRPC-ошибка — вернуть как есть
		return err
	}

	// не gRPC — завернуть
	return NewGRPCError(err)
}

func NewGRPCClient(client *client.GRPCClient) *GRPCClient {
	return &GRPCClient{
		client: client,
		pb:     pb.NewInterserviceServiceClient(client.Conn()),
	}
}
