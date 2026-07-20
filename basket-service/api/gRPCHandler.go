package api

import (
	"context"

	"github.com/Fox216540/shop/basket-service/app"
	"github.com/Fox216540/shop/basket-service/domain/money"
	"github.com/Fox216540/shop/basket-service/domain/productShort"
	pbApi "github.com/Fox216540/shop/proto/basket-service/gen/api"
	types "github.com/Fox216540/shop/proto/common/gen"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCHandler struct {
	basketUS app.UseCase
	pbApi.UnimplementedApiServiceServer
	mapper *ErrorMapper
}

func NewGRPCHandler(basketUS app.UseCase, mapper *ErrorMapper) *GRPCHandler {
	return &GRPCHandler{
		basketUS: basketUS,
		mapper:   mapper,
	}
}

func (h *GRPCHandler) getUserIDFromMetadata(ctx context.Context) (uuid.UUID, error) {
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
		return uuid.Nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return userID, nil
}

func toProtoMoney(price money.Money) *types.Money {
	return &types.Money{
		Amount:   price.Amount,
		Currency: price.Currency,
	}
}

func toProtoProduct(product productShort.ProductShort) *types.Product {
	return &types.Product{
		Id:   product.ID.String(),
		Name: product.Name,
		Img:  product.Img,
		Price: &types.Money{
			Amount:   product.Price.Amount,
			Currency: product.Price.Currency,
		},
	}
}

func (h *GRPCHandler) GetBasket(ctx context.Context, req *emptypb.Empty) (*pbApi.GetBasketResponse, error) {
	userID, err := h.getUserIDFromMetadata(ctx)
	if err != nil {
		return nil, h.mapper.MapError(ctx, err)
	}

	basketDomain, err := h.basketUS.GetBasket(userID)
	if err != nil {
		return nil, h.mapper.MapError(ctx, err)
	}

	items := make([]*pbApi.BasketItem, 0, len(basketDomain.Products))
	for _, item := range basketDomain.Products {
		items = append(items, &pbApi.BasketItem{
			Product:  toProtoProduct(item.Product),
			Quantity: item.Quantity,
		})
	}

	return &pbApi.GetBasketResponse{
		Items: items,
		Total: toProtoMoney(basketDomain.Total),
	}, nil
}

func (h *GRPCHandler) AddItemToBasket(ctx context.Context, req *pbApi.AddItemToBasketRequest) (*pbApi.AddItemToBasketResponse, error) {
	userID, err := h.getUserIDFromMetadata(ctx)
	if err != nil {
		return nil, h.mapper.MapError(ctx, err)
	}

	productID, err := uuid.Parse(req.GetProductId())
	if err != nil {
		return nil, h.mapper.MapError(ctx, status.Error(codes.InvalidArgument, err.Error()))
	}

	item, err := h.basketUS.AddItemToBasket(userID, productID, req.GetQuantity())
	if err != nil {
		return nil, h.mapper.MapError(ctx, err)
	}

	return &pbApi.AddItemToBasketResponse{
		Item: &pbApi.BasketItem{
			Product:  toProtoProduct(item.Product),
			Quantity: item.Quantity,
		},
	}, nil
}

func (h *GRPCHandler) DeleteBasket(ctx context.Context, req *emptypb.Empty) (*types.MessageResponse, error) {
	userID, err := h.getUserIDFromMetadata(ctx)
	if err != nil {
		return nil, h.mapper.MapError(ctx, err)
	}

	if err := h.basketUS.DeleteBasket(userID); err != nil {
		return nil, h.mapper.MapError(ctx, err)
	}

	return &types.MessageResponse{Message: messages.BasketDeletedSuccessfully}, nil
}

func (h *GRPCHandler) RemoveItemFromBasket(ctx context.Context, req *pbApi.RemoveItemFromBasketRequest) (*types.MessageResponse, error) {
	userID, err := h.getUserIDFromMetadata(ctx)
	if err != nil {
		return nil, h.mapper.MapError(ctx, err)
	}

	productID, err := uuid.Parse(req.GetProductId())
	if err != nil {
		return nil, h.mapper.MapError(ctx, status.Error(codes.InvalidArgument, err.Error()))
	}

	if err := h.basketUS.RemoveItemFromBasket(userID, productID); err != nil {
		return nil, h.mapper.MapError(ctx, err)
	}

	return &types.MessageResponse{Message: messages.ItemRemovedSuccessfully}, nil
}
