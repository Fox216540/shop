package api

import (
	"context"
	"errors"
	types "github.com/Fox216540/shop/proto/common/gen"
	pbApi "github.com/Fox216540/shop/proto/user-service/gen/api"
	pbInterservice "github.com/Fox216540/shop/proto/user-service/gen/interservice"
	user "github.com/Fox216540/shop/user-service/app"
	"github.com/Fox216540/shop/user-service/app/dto"
	"github.com/Fox216540/shop/user-service/domain/auth"
	userDomain "github.com/Fox216540/shop/user-service/domain/user"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCHandler struct {
	userUS user.UseCase
	pbApi.UnimplementedApiServiceServer
	pbInterservice.UnimplementedInterserviceServiceServer
	mapper *ErrorMapper
}

func (h *GRPCHandler) returnUserWithTokensResponse(
	u userDomain.User, tokens auth.Tokens,
	msg string) *types.UserWithTokensResponse {
	return &types.UserWithTokensResponse{
		Name: &types.UserNameResponse{
			Name: u.Name,
		},
		Tokens: &types.TokensResponse{
			AccessToken:  tokens.AccessToken,
			RefreshToken: tokens.RefreshToken,
		},
		Message: &types.MessageResponse{
			Message: msg,
		},
	}
}

func (h *GRPCHandler) returnUserWithMessageResponse(
	u userDomain.User, msg string,
) *pbApi.UserWithMessageResponse {
	return &pbApi.UserWithMessageResponse{
		Name: &types.UserNameResponse{
			Name: u.Name,
		},
		Message: &types.MessageResponse{
			Message: msg,
		},
	}
}

func (h *GRPCHandler) getUserIDFromMetadata(
	ctx context.Context,
) (uuid.UUID, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return uuid.Nil, errors.New("metadata not found")
	}
	userID, err := uuid.Parse(md.Get("user_id")[0])
	if err != nil {
		return uuid.Nil, err
	}
	return userID, nil
}

func (h *GRPCHandler) RegisterUser(
	ctx context.Context,
	req *pbApi.RegisterUserRequest,
) (*types.UserWithTokensResponse, error) {
	uDomain := userDomain.User{
		Email:    req.Email,
		Name:     req.Name,
		Phone:    req.Phone,
		Password: req.Password,
		Address:  req.Address,
	}
	u, tokens, err := h.userUS.Register(uDomain)
	if err != nil {
		return nil, h.mapper.MapError(err)
	}
	return h.returnUserWithTokensResponse(
		u,
		tokens,
		"User registered successfully",
	), nil
}

func (h *GRPCHandler) DeleteUser(
	ctx context.Context,
	req *emptypb.Empty,
) (*types.MessageResponse, error) {
	userID, err := h.getUserIDFromMetadata(ctx)
	if err != nil {
		return nil, h.mapper.MapError(err)
	}
	if err = h.userUS.DeleteUser(userID); err != nil {
		return nil, h.mapper.MapError(err)
	}
	return &types.MessageResponse{
		Message: "User deleted successfully",
	}, nil
}

func (h *GRPCHandler) VerifyCredentials(
	ctx context.Context,
	req *types.CredentialsRequest,
) (*pbInterservice.VerifyCredentialsResponse, error) {
	name, id, err := h.userUS.VerifyCredentials(req.PhoneOrEmail, req.Password)
	if err != nil {
		return nil, h.mapper.MapError(err)
	}
	return &pbInterservice.VerifyCredentialsResponse{
		Id:   id.String(),
		Name: name,
	}, nil
}

func (h *GRPCHandler) UpdateEmail(
	ctx context.Context,
	req *pbApi.UpdateEmailRequest,
) (*pbApi.UserWithMessageResponse, error) {
	userID, err := h.getUserIDFromMetadata(ctx)
	if err != nil {
		return nil, h.mapper.MapError(err)
	}
	u, err := h.userUS.UpdateEmail(userID, req.Email)
	if err != nil {
		return nil, h.mapper.MapError(err)
	}
	return h.returnUserWithMessageResponse(
		u,
		"Email updated successfully",
	), nil
}

func (h *GRPCHandler) UpdatePassword(
	ctx context.Context,
	req *pbApi.UpdatePasswordRequest,
) (*pbApi.UserWithMessageResponse, error) {
	userID, err := h.getUserIDFromMetadata(ctx)
	if err != nil {
		return nil, h.mapper.MapError(err)
	}
	u, err := h.userUS.UpdatePassword(userID, req.Password)
	if err != nil {
		return nil, h.mapper.MapError(err)
	}
	return h.returnUserWithMessageResponse(
		u,
		"Password updated successfully",
	), nil
}

func (h *GRPCHandler) UpdatePhone(
	ctx context.Context,
	req *pbApi.UpdatePhoneRequest,
) (*pbApi.UserWithMessageResponse, error) {
	userID, err := h.getUserIDFromMetadata(ctx)
	if err != nil {
		return nil, h.mapper.MapError(err)
	}
	u, err := h.userUS.UpdatePhone(userID, req.Phone)
	if err != nil {
		return nil, h.mapper.MapError(err)
	}
	return h.returnUserWithMessageResponse(
		u,
		"Phone updated successfully",
	), nil
}

func (h *GRPCHandler) UpdateProfile(
	ctx context.Context,
	req *pbApi.UpdateProfileRequest,
) (*pbApi.UserWithMessageResponse, error) {
	userID, err := h.getUserIDFromMetadata(ctx)
	if err != nil {
		return nil, h.mapper.MapError(err)
	}
	dtoUS := dto.ProfileUpdate{
		Name:    req.Name,
		Address: req.Address,
	}
	u, err := h.userUS.UpdateProfile(userID, dtoUS)
	if err != nil {
		return nil, h.mapper.MapError(err)
	}
	return h.returnUserWithMessageResponse(
		u,
		"Profile updated successfully",
	), nil
}

func NewGRPCHandler(userUS user.UseCase) *GRPCHandler {
	return &GRPCHandler{
		userUS: userUS,
	}
}
