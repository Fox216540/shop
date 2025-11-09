package auth

import (
	"github.com/Fox216540/shop/apigateway-service/domain/auth"
	"github.com/Fox216540/shop/apigateway-service/infra/client"
	pb "github.com/Fox216540/shop/proto/auth-service/gen/api"
	types "github.com/Fox216540/shop/proto/common/gen"
	"github.com/google/uuid"
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

func (c *GRPCClient) LogInUser(phoneOrEmail, password string) (name string, tokens auth.Tokens, message string, err error) {
	ctx := c.conn.Context()

	req := &types.CredentialsRequest{
		PhoneOrEmail: phoneOrEmail,
		Password:     password,
	}

	resp, err := c.pb.LogIn(ctx, req)
	if err != nil {
		return "", auth.Tokens{}, "", err
	}
	return resp.Name.Name, auth.Tokens{
		AccessToken:  resp.Tokens.AccessToken,
		RefreshToken: resp.Tokens.RefreshToken,
	}, resp.Message.Message, nil
}

func (c *GRPCClient) LogOutUser(token string) (msg string, error error) {
	ctx := c.conn.Context()

	req := &types.DecodeTokenRequest{
		Token: token,
	}

	resp, err := c.pb.LogOut(ctx, req)
	if err != nil {
		return "", err
	}
	return resp.Message, nil
}

func (c *GRPCClient) LogOutAllUser(token string) (msg string, error error) {
	ctx := c.conn.Context()

	req := &types.DecodeTokenRequest{
		Token: token,
	}

	resp, err := c.pb.LogOutAll(ctx, req)
	if err != nil {
		return "", err
	}
	return resp.Message, nil
}

func (c *GRPCClient) RefreshTokensOfUser(token string) (tokens auth.Tokens, error error) {
	ctx := c.conn.Context()

	req := &types.DecodeTokenRequest{
		Token: token,
	}

	resp, err := c.pb.RefreshTokens(ctx, req)
	if err != nil {
		return auth.Tokens{}, err
	}
	return auth.Tokens{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
	}, nil
}

func (c *GRPCClient) DecodeAccessTokenOfUser(token string) (userID uuid.UUID, error error) {
	ctx := c.conn.Context()

	req := &types.DecodeTokenRequest{
		Token: token,
	}

	resp, err := c.pb.DecodeAccessToken(ctx, req)
	if err != nil {
		return uuid.UUID{}, err
	}

	UUIDUserID, err := uuid.Parse(resp.Id)
	if err != nil {
		return uuid.Nil, err
	}

	return UUIDUserID, nil
}
