package api

import (
	"context"
	"errors"
	pb "github.com/Fox216540/shop/auth-service/api/proto"
	"github.com/Fox216540/shop/auth-service/app/auth"
	authDomain "github.com/Fox216540/shop/auth-service/domain/auth"
	"github.com/google/uuid"
)

type GRPCHandler struct {
	auth auth.UseCase
	pb.UnimplementedAuthServiceServer
}

func (h *GRPCHandler) SignUp(ctx context.Context, req *pb.SignRequest) (*pb.SignResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}
	a := authDomain.Auth{
		UserID:   userID,
		Password: req.Password,
	}
	hash, tokens, err := h.auth.SignUp(a)

	if err != nil {
		return nil, err
	}

	return &pb.SignResponse{
		Hash:         hash,
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}

func (h *GRPCHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.TokensResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}
	a := authDomain.Auth{
		UserID:   userID,
		Password: req.Password,
	}
	tokens, err := h.auth.Login(a, req.Hash)
	if err != nil {
		return nil, err
	}
	return &pb.TokensResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}

func (h *GRPCHandler) DecodeToken(ctx context.Context, req *pb.DecodeTokenRequest) (*pb.DecodeTokenResponse, error) {
	token := req.Token
	tokenType := req.TokenType

	switch tokenType {
	case pb.TokenType_ACCESS:
		userJWT, err := h.auth.DecodeAccessToken(token)
		if err != nil {
			return nil, err
		}
		return &pb.DecodeTokenResponse{
			UserId: userJWT.UserID.String(),
			Jti:    userJWT.JTI.String(),
		}, nil
	case pb.TokenType_REFRESH:
		userJWT, err := h.auth.DecodeRefreshToken(token)
		if err != nil {
			return nil, err
		}
		return &pb.DecodeTokenResponse{
			UserId: userJWT.UserID.String(),
			Jti:    userJWT.JTI.String(),
		}, nil
	default:
		return nil, errors.New("invalid token type")
	}
}

func (h *GRPCHandler) DeleteRefreshToken(ctx context.Context, req *pb.DeleteRefreshTokenRequest) (*pb.DeleteSuccessResponse, error) {
	if err := h.auth.DeleteRefreshToken(req.Token); err != nil {
		return nil, err
	}
	return &pb.DeleteSuccessResponse{
		//TODO: Вынести в константы
		Message: "Refresh token deleted successfully",
	}, nil
}

func (h *GRPCHandler) DeleteAllTokens(ctx context.Context, req *pb.DeleteAllTokensRequest) (*pb.DeleteSuccessResponse, error) {
	if err := h.auth.DeleteAllTokens(req.Token); err != nil {
		return nil, err
	}
	return &pb.DeleteSuccessResponse{
		//TODO: Вынести в константы
		Message: "All tokens deleted successfully",
	}, nil
}

func (h *GRPCHandler) NewPassword(ctx context.Context, req *pb.NewPasswordRequest) (*pb.NewPasswordResponse, error) {
	hash, err := h.auth.NewPassword(req.OldPassword, req.Hash, req.NewPassword)
	if err != nil {
		return nil, err
	}
	return &pb.NewPasswordResponse{
		Hash: hash,
	}, nil
}

func (h *GRPCHandler) RefreshTokens(ctx context.Context, req *pb.NewTokensRequest) (*pb.TokensResponse, error) {
	tokens, err := h.auth.RefreshTokens(req.Token)
	if err != nil {
		return nil, err
	}
	return &pb.TokensResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}

func NewGRPCHandler(useCase auth.UseCase) *GRPCHandler {
	return &GRPCHandler{
		auth: useCase,
	}
}
