package services

import (
	"backend/grpc"
	pb "backend/proto"
	"context"
)

type CartService struct {
}

func NewCartService() *CartService {
	return &CartService{}
}

func (c *CartService) AddItem(ctx context.Context, sessionID string, item *pb.CartItem) error {
	_, err := pb.NewCartServiceClient(grpc.ApiServerInstance.CartServiceConn).AddItem(ctx, &pb.AddItemRequest{
		SessionId: sessionID,
		Item:      item,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *CartService) GetCart(ctx context.Context, sessionID string) (*pb.Cart, error) {
	resp, err := pb.NewCartServiceClient(grpc.ApiServerInstance.CartServiceConn).GetCart(ctx, &pb.GetCartRequest{
		SessionId: sessionID,
	})
	return resp, err
}

func (c *CartService) EmptyCart(ctx context.Context, sessionID string) error {
	_, err := pb.NewCartServiceClient(grpc.ApiServerInstance.CartServiceConn).EmptyCart(ctx, &pb.EmptyCartRequest{
		SessionId: sessionID,
	})
	return err
}

func (c *CartService) RemoveItem(ctx context.Context, sessionID string, productId uint64) error {
	_, err := pb.NewCartServiceClient(grpc.ApiServerInstance.CartServiceConn).RemoveItem(ctx, &pb.RemoveItemRequest{
		SessionId: sessionID,
		Id:        productId,
	})
	return err
}

func (c *CartService) UpdateItemQuantity(ctx context.Context, sessionID string, productID uint64, quantity uint64) error {
	_, err := pb.NewCartServiceClient(grpc.ApiServerInstance.CartServiceConn).UpdateItemQuantity(ctx, &pb.UpdateItemQuantityRequest{
		SessionId: sessionID,
		Id:        productID,
		Quantity:  quantity,
	})
	return err
}
