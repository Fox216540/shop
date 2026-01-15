package user

import (
	"github.com/Fox216540/shop/auth-service/infra/client"
	types "github.com/Fox216540/shop/proto/common/gen"
	pb "github.com/Fox216540/shop/proto/user-service/gen/interservice"
	"github.com/google/uuid"
	"google.golang.org/grpc/status"
)

type GRPCClient struct {
	client *client.GRPCClient
	pb     pb.InterserviceServiceClient
}

func (c *GRPCClient) VerifyCredentialsOfUser(phoneOrEmail, password string) (name string, id uuid.UUID, error error) {
	ctx := c.client.Context()
	req := &types.CredentialsRequest{
		PhoneOrEmail: phoneOrEmail,
		Password:     password,
	}
	resp, err := c.pb.VerifyCredentials(ctx, req)
	if err != nil {
		if _, ok := status.FromError(err); ok {
			return "", uuid.Nil, err
		}

		return "", uuid.Nil, NewGRPCError(err)
	}
	userID, err := uuid.Parse(resp.Id)
	if err != nil {
		return "", uuid.Nil, NewInvalidUUID(err)
	}
	return resp.Name, userID, nil
}
