package api

import (
	"context"
	"github.com/Fox216540/shop/auth-service/app/auth"
	"github.com/Fox216540/shop/auth-service/app/user"
	pbApi "github.com/Fox216540/shop/proto/auth-service/gen/api"
	pbInterservice "github.com/Fox216540/shop/proto/auth-service/gen/interservice"
	types "github.com/Fox216540/shop/proto/common/gen"
	"github.com/google/uuid"
)

type GRPCHandler struct {
	auth auth.UseCase
	user user.UseCase
	pbApi.UnimplementedApiServiceServer
	pbInterservice.UnimplementedInterserviceServiceServer
}

func (h *GRPCHandler) LogIn(
	ctx context.Context,
	req *types.CredentialsRequest,
) (*types.UserWithTokensResponse, error) {
	name, id, err := h.user.VerifyUser(req.PhoneOrEmail, req.Password)
	if err != nil {
		return nil, err
	}
	tokens, err := h.auth.GenerateTokens(id)
	if err != nil {
		return nil, err
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
			Message: "User logged in successfully",
		},
	}, nil
}

func (h *GRPCHandler) LogOut(
	ctx context.Context,
	req *types.DecodeTokenRequest,
) (*types.MessageResponse, error) {
	if err := h.auth.DeleteRefreshToken(req.Token); err != nil {
		return nil, err
	}
	return &types.MessageResponse{
		//TODO: Вынести в константы
		Message: "Token deleted successfully",
	}, nil
}

func (h *GRPCHandler) LogOutAll(
	ctx context.Context,
	req *types.DecodeTokenRequest,
) (*types.MessageResponse, error) {
	if err := h.auth.DeleteAllTokensByToken(req.Token); err != nil {
		return nil, err
	}
	return &types.MessageResponse{
		//TODO: Вынести в константы
		Message: "All tokens deleted successfully",
	}, nil
}

func (h *GRPCHandler) RefreshTokens(
	ctx context.Context,
	req *types.DecodeTokenRequest,
) (*types.TokensResponse, error) {
	tokens, err := h.auth.RefreshTokens(req.Token)
	if err != nil {
		return nil, err
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
	userJWT, err := h.auth.DecodeAccessToken(req.Token)
	if err != nil {
		return nil, err
	}
	return &types.UserId{
		Id: userJWT.UserID.String(),
	}, nil
}

func (h *GRPCHandler) GenerateTokens(
	ctx context.Context,
	req *types.UserId,
) (*types.TokensResponse, error) {
	userID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	tokens, err := h.auth.GenerateTokens(userID)
	if err != nil {
		return nil, err
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
	userID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	if err := h.auth.DeleteAllTokensByUserID(userID); err != nil {
		return nil, err
	}
	return &types.MessageResponse{
		//TODO: Вынести в константы
		Message: "All tokens deleted successfully",
	}, nil
}

func NewGRPCHandler(useCase auth.UseCase) *GRPCHandler {
	return &GRPCHandler{
		auth: useCase,
	}
}
