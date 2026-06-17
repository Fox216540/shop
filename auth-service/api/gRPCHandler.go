package api

import (
	"context"
	"github.com/Fox216540/shop/auth-service/app"
	pbApi "github.com/Fox216540/shop/proto/auth-service/gen/api"
	pbInterservice "github.com/Fox216540/shop/proto/auth-service/gen/interservice"
	types "github.com/Fox216540/shop/proto/common/gen"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCHandler struct {
	auth app.UseCase
	pbApi.UnimplementedApiServiceServer
	pbInterservice.UnimplementedInterserviceServiceServer
	mapper *ErrorMapper
}

func (h *GRPCHandler) LogIn(
	ctx context.Context,
	req *types.CredentialsRequest,
) (*types.UserWithTokensResponse, error) {
	if req == nil {
		return nil, h.mapper.MapError(ctx, status.Error(codes.InvalidArgument, NewMessages().InvalidArgument))
	}
	name, tokens, err := h.auth.LogIn(req.PhoneOrEmail, req.Password)
	if err != nil {
		return nil, h.mapper.MapError(ctx, err)
	}

	return &types.UserWithTokensResponse{
		Tokens: &types.TokensResponse{
			AccessToken:  tokens.AccessToken,
			RefreshToken: tokens.RefreshToken,
		},
		Name: &types.UserNameResponse{
			Name: name,
		},
		Message: &types.MessageResponse{
			Message: NewMessages().LogInSuccess,
		},
	}, nil
}

func (h *GRPCHandler) LogOut(
	ctx context.Context,
	req *types.DecodeTokenRequest,
) (*types.MessageResponse, error) {
	if req == nil {
		return nil, h.mapper.MapError(ctx, status.Error(codes.InvalidArgument, NewMessages().InvalidArgument))
	}
	if err := h.auth.DeleteRefreshToken(req.Token); err != nil {
		return nil, h.mapper.MapError(ctx, err)
	}
	return &types.MessageResponse{
		Message: NewMessages().LogOutSuccess,
	}, nil
}

func (h *GRPCHandler) LogOutAll(
	ctx context.Context,
	req *types.DecodeTokenRequest,
) (*types.MessageResponse, error) {
	if req == nil {
		return nil, h.mapper.MapError(ctx, status.Error(codes.InvalidArgument, NewMessages().InvalidArgument))
	}
	if err := h.auth.DeleteAllTokensByToken(req.Token); err != nil {
		return nil, h.mapper.MapError(ctx, err)
	}
	return &types.MessageResponse{
		Message: NewMessages().LogOutAllSuccess,
	}, nil
}

func (h *GRPCHandler) RefreshTokens(
	ctx context.Context,
	req *types.DecodeTokenRequest,
) (*types.TokensResponse, error) {
	if req == nil {
		return nil, h.mapper.MapError(ctx, status.Error(codes.InvalidArgument, NewMessages().InvalidArgument))
	}
	tokens, err := h.auth.RefreshTokens(req.Token)
	if err != nil {
		return nil, h.mapper.MapError(ctx, err)
	}
	return &types.TokensResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}

func (h *GRPCHandler) DecodeAccessToken(
	ctx context.Context,
	req *types.DecodeTokenRequest,
) (*types.UserId, error) {
	if req == nil {
		return nil, h.mapper.MapError(ctx, status.Error(codes.InvalidArgument, NewMessages().InvalidArgument))
	}
	userJWT, err := h.auth.DecodeAccessToken(req.Token)
	if err != nil {
		return nil, h.mapper.MapError(ctx, err)
	}
	return &types.UserId{
		Id: userJWT.UserID.String(),
	}, nil
}

func (h *GRPCHandler) GenerateTokens(
	ctx context.Context,
	req *types.UserId,
) (*types.TokensResponse, error) {
	if req == nil {
		return nil, h.mapper.MapError(ctx, status.Error(codes.InvalidArgument, NewMessages().InvalidArgument))
	}
	userID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, h.mapper.MapError(ctx, err)
	}
	tokens, err := h.auth.GenerateTokens(userID)
	if err != nil {
		return nil, h.mapper.MapError(ctx, err)
	}
	return &types.TokensResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}

func (h *GRPCHandler) DeleteAllRefreshTokens(
	ctx context.Context,
	req *types.UserId,
) (*types.MessageResponse, error) {
	if req == nil {
		return nil, h.mapper.MapError(ctx, status.Error(codes.InvalidArgument, NewMessages().InvalidArgument))
	}
	userID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, h.mapper.MapError(ctx, err)
	}

	if err = h.auth.DeleteAllTokensByUserID(userID); err != nil {
		return nil, h.mapper.MapError(ctx, err)
	}
	return &types.MessageResponse{
		Message: NewMessages().DeleteAllRefreshTokensSuccess,
	}, nil
}

func NewGRPCHandler(authUC app.UseCase, mapper *ErrorMapper) *GRPCHandler {
	return &GRPCHandler{
		auth:   authUC,
		mapper: mapper,
	}
}
