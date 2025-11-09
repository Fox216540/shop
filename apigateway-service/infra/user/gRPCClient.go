package user

import (
	"github.com/Fox216540/shop/apigateway-service/domain/auth"
	"github.com/Fox216540/shop/apigateway-service/domain/user"
	"github.com/Fox216540/shop/apigateway-service/infra/client"
	pb "github.com/Fox216540/shop/proto/user-service/gen/api"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
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

func (c *GRPCClient) RegisterUser(u user.User) (name string, tokens auth.Tokens, message string, err error) {
	ctx := c.conn.Context()

	req := &pb.RegisterUserRequest{
		Name:     u.Name,
		Email:    u.Email,
		Phone:    u.Phone,
		Password: u.Password,
		Address:  u.Address,
	}

	resp, err := c.pb.RegisterUser(ctx, req)
	if err != nil {
		return "", auth.Tokens{}, "", err
	}
	return resp.Name.Name, auth.Tokens{
		AccessToken:  resp.Tokens.AccessToken,
		RefreshToken: resp.Tokens.RefreshToken,
	}, resp.Message.Message, nil
}

func (c *GRPCClient) DeleteUser(id uuid.UUID) (message string, err error) {
	ctx := c.conn.Context()
	ctx = metadata.AppendToOutgoingContext(ctx, "user_id", id.String())
	resp, err := c.pb.DeleteUser(ctx, &emptypb.Empty{})
	if err != nil {
		return "", err
	}
	return resp.Message, nil
}
