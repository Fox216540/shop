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
	conn     *client.GRPCClient
	pbClient pb.ApiServiceClient
}

func NewGRPCClient(client *client.GRPCClient) *GRPCClient {
	return &GRPCClient{
		conn:     client,
		pbClient: pb.NewApiServiceClient(client.Conn()), // Инициализация клиента
	}
}

func (c *GRPCClient) Register(u user.User) (name string, tokens auth.Tokens, message string, err error) {
	ctx := c.conn.Context()

	req := &pb.RegisterUserRequest{
		Name:     u.Name,
		Email:    u.Email,
		Phone:    u.Phone,
		Password: u.Password,
		Address:  u.Address,
	}

	resp, err := c.pbClient.RegisterUser(ctx, req)
	if err != nil {
		return "", auth.Tokens{}, "", err
	}
	return resp.Name.Name, auth.Tokens{
		AccessToken:  resp.Tokens.AccessToken,
		RefreshToken: resp.Tokens.RefreshToken,
	}, resp.Message.Message, nil
}

func (c *GRPCClient) Delete(id uuid.UUID) (message string, err error) {
	ctx := c.conn.Context()
	ctx = metadata.AppendToOutgoingContext(ctx, "user_id", id.String())
	resp, err := c.pbClient.DeleteUser(ctx, &emptypb.Empty{})
	if err != nil {
		return "", err
	}
	return resp.Message, nil
}

func (c *GRPCClient) UpdateEmail(id uuid.UUID, email string) (message string, err error) {
	ctx := c.conn.Context()
	ctx = metadata.AppendToOutgoingContext(ctx, "user_id", id.String())
	req := &pb.UpdateEmailRequest{
		Email: email,
	}
	resp, err := c.pbClient.UpdateEmail(ctx, req)
	if err != nil {
		return "", err
	}
	return resp.Message.Message, nil
}

func (c *GRPCClient) UpdatePassword(id uuid.UUID, password string) (message string, err error) {
	ctx := c.conn.Context()
	ctx = metadata.AppendToOutgoingContext(ctx, "user_id", id.String())
	req := &pb.UpdatePasswordRequest{
		Password: password,
	}
	resp, err := c.pbClient.UpdatePassword(ctx, req)
	if err != nil {
		return "", err
	}
	return resp.Message.Message, nil
}

func (c *GRPCClient) UpdatePhone(id uuid.UUID, phone string) (message string, err error) {
	ctx := c.conn.Context()
	ctx = metadata.AppendToOutgoingContext(ctx, "user_id", id.String())
	req := &pb.UpdatePhoneRequest{
		Phone: phone,
	}
	resp, err := c.pbClient.UpdatePhone(ctx, req)
	if err != nil {
		return "", err
	}
	return resp.Message.Message, nil
}

func (c *GRPCClient) UpdateProfile(id uuid.UUID, name *string, address *string) (newName string, message string, err error) {
	ctx := c.conn.Context()
	ctx = metadata.AppendToOutgoingContext(ctx, "user_id", id.String())
	req := &pb.UpdateProfileRequest{
		Name:    name,
		Address: address,
	}
	resp, err := c.pbClient.UpdateProfile(ctx, req)
	if err != nil {
		return "", "", err
	}
	return resp.Name.Name, resp.Message.Message, nil
}
