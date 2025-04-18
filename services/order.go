package services

import (
	"context"

	"backend/grpc"
	pb "backend/proto"
)

type OrderService struct{}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (o *OrderService) GetOrder(ctx context.Context, orderId uint64) (*pb.Order, error) {
	resp, err := pb.NewOrderServiceClient(grpc.ApiServerInstance.OrderServiceConn).GetOrder(ctx, &pb.GetOrderRequest{
		Id: orderId,
	})
	return resp, err
}

func (o *OrderService) GetOrdersByUser(ctx context.Context, userId uint64) (*pb.GetOrdersResponse, error) {
	resp, err := pb.NewOrderServiceClient(grpc.ApiServerInstance.OrderServiceConn).GetOrdersByUser(ctx, &pb.GetOrdersByUserRequest{
		UserId: userId,
	})
	return resp, err
}

func (o *OrderService) GetOrdersByMerchant(ctx context.Context, merchantId uint64) (*pb.GetOrdersResponse, error) {
	resp, err := pb.NewOrderServiceClient(grpc.ApiServerInstance.OrderServiceConn).GetOrdersByMerchant(ctx, &pb.GetOrdersByMerchantRequest{
		MerchantId: merchantId,
	})
	return resp, err
}

func (o *OrderService) UpdateOrderStatus(ctx context.Context, orderId uint64, status string) (*pb.Order, error) {
	resp, err := pb.NewOrderServiceClient(grpc.ApiServerInstance.OrderServiceConn).UpdateOrderStatus(ctx, &pb.UpdateOrderStatusRequest{
		Id:     orderId,
		Status: status,
	})
	return resp, err
}

func (o *OrderService) CancelOrder(ctx context.Context, orderId uint64) (*pb.Order, error) {
	resp, err := pb.NewOrderServiceClient(grpc.ApiServerInstance.OrderServiceConn).CancelOrder(ctx, &pb.CancelOrderRequest{
		Id: orderId,
	})
	return resp, err
}

func (o *OrderService) UpdatePaymentStatus(ctx context.Context, event string, orderId uint64) error {
	_, err := pb.NewOrderServiceClient(grpc.ApiServerInstance.OrderServiceConn).UpdatePaymentStatus(ctx, &pb.UpdatePaymentStatusRequest{
		Event:   event,
		OrderId: orderId,
	})

	return err
}
