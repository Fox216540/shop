package api

import (
	"context"

	types "github.com/Fox216540/shop/proto/common/gen"
	pbApi "github.com/Fox216540/shop/proto/user-service/gen/api"
	pbInterservice "github.com/Fox216540/shop/proto/user-service/gen/interservice"
	user "github.com/Fox216540/shop/user-service/app"
	"github.com/Fox216540/shop/user-service/app/dto"
	"github.com/Fox216540/shop/user-service/domain/auth"
	userDomain "github.com/Fox216540/shop/user-service/domain/user"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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
		return uuid.Nil, status.Error(codes.InvalidArgument, messages.MetadataNotFound)
	}

	values := md.Get("user_id")
	if len(values) == 0 {
		return uuid.Nil, status.Error(codes.InvalidArgument, messages.UserIDMetadataNotFound)
	}

	userID, err := uuid.Parse(values[0])
	if err != nil {
		return uuid.Nil, status.Error(codes.InvalidArgument, messages.InvalidArgument)
	}
	return userID, nil
}

func (h *GRPCHandler) RegisterUser(
	ctx context.Context,
	req *pbApi.RegisterUserRequest,
) (*types.UserWithTokensResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, messages.InvalidArgument)
	}
	uDomain := userDomain.User{
		Email:    req.Email,
		Name:     req.Name,
		Phone:    req.Phone,
		Password: req.Password,
		Address:  req.Address,
	}
	u, tokens, err := h.userUS.Register(uDomain)
	if err != nil {
		return nil, h.mapper.MapError(ctx, err)
	}
	return h.returnUserWithTokensResponse(
		u,
		tokens,
		messages.UserRegistered,
	), nil
}

func (h *GRPCHandler) DeleteUser(
	ctx context.Context,
	req *emptypb.Empty,
) (*types.MessageResponse, error) {
	userID, err := h.getUserIDFromMetadata(ctx)
	if err != nil {
		return nil, h.mapper.MapError(ctx, err)
	}
	if err = h.userUS.DeleteUser(userID); err != nil {
		return nil, h.mapper.MapError(ctx, err)
	}
	return &types.MessageResponse{
		Message: messages.UserDeleted,
	}, nil
}

func (h *GRPCHandler) VerifyCredentials(
	ctx context.Context,
	req *types.CredentialsRequest,
) (*pbInterservice.VerifyCredentialsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, messages.InvalidArgument)
	}
	name, id, err := h.userUS.VerifyCredentials(req.PhoneOrEmail, req.Password)
	if err != nil {
		return nil, h.mapper.MapError(ctx, err)
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
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, messages.InvalidArgument)
	}
	userID, err := h.getUserIDFromMetadata(ctx)
	if err != nil {
		return nil, h.mapper.MapError(ctx, err)
	}
	u, err := h.userUS.UpdateEmail(userID, req.Email)
	if err != nil {
		return nil, h.mapper.MapError(ctx, err)
	}
	return h.returnUserWithMessageResponse(
		u,
		messages.EmailUpdated,
	), nil
}

func (h *GRPCHandler) UpdatePassword(
	ctx context.Context,
	req *pbApi.UpdatePasswordRequest,
) (*pbApi.UserWithMessageResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, messages.InvalidArgument)
	}
	userID, err := h.getUserIDFromMetadata(ctx)
	if err != nil {
		return nil, h.mapper.MapError(ctx, err)
	}
	u, err := h.userUS.UpdatePassword(userID, req.Password)
	if err != nil {
		return nil, h.mapper.MapError(ctx, err)
	}
	return h.returnUserWithMessageResponse(
		u,
		messages.PasswordUpdated,
	), nil
}

func (h *GRPCHandler) UpdatePhone(
	ctx context.Context,
	req *pbApi.UpdatePhoneRequest,
) (*pbApi.UserWithMessageResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, messages.InvalidArgument)
	}
	userID, err := h.getUserIDFromMetadata(ctx)
	if err != nil {
		return nil, h.mapper.MapError(ctx, err)
	}
	u, err := h.userUS.UpdatePhone(userID, req.Phone)
	if err != nil {
		return nil, h.mapper.MapError(ctx, err)
	}
	return h.returnUserWithMessageResponse(
		u,
		messages.PhoneUpdated,
	), nil
}

func (h *GRPCHandler) UpdateProfile(
	ctx context.Context,
	req *pbApi.UpdateProfileRequest,
) (*pbApi.UserWithMessageResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, messages.InvalidArgument)
	}
	userID, err := h.getUserIDFromMetadata(ctx)
	if err != nil {
		return nil, h.mapper.MapError(ctx, err)
	}
	dtoUS := dto.ProfileUpdate{
		Name:    req.Name,
		Address: req.Address,
	}
	u, err := h.userUS.UpdateProfile(userID, dtoUS)
	if err != nil {
		return nil, h.mapper.MapError(ctx, err)
	}
	return h.returnUserWithMessageResponse(
		u,
		messages.ProfileUpdated,
	), nil
}

func NewGRPCHandler(userUS user.UseCase, mapper *ErrorMapper) *GRPCHandler {
	return &GRPCHandler{
		userUS: userUS,
		mapper: mapper,
	}
}
