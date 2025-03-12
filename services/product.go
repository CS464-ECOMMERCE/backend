package services

import (
	"backend/grpc"
	pb "backend/proto"
	"context"
)

type ProductService struct {
}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (p *ProductService) GetProduct(ctx context.Context) ([]*pb.Product, error) {
	resp, err := pb.NewProductServiceClient(grpc.ApiServerInstance.ProductServiceConn).ListProducts(ctx, &pb.Empty{})
	return resp.GetProducts(), err
}

func (p *ProductService) CreateProduct(ctx context.Context, product *pb.CreateProductRequest) (*pb.Product, error) {
	resp, err := pb.NewProductServiceClient(grpc.ApiServerInstance.ProductServiceConn).CreateProduct(ctx, product)
	return resp, err
}

func (p *ProductService) UpdateProduct(ctx context.Context, product *pb.UpdateProductRequest) (*pb.Product, error) {
	resp, err := pb.NewProductServiceClient(grpc.ApiServerInstance.ProductServiceConn).UpdateProduct(ctx, product)
	return resp, err
}

func (p *ProductService) DeleteProduct(ctx context.Context, product *pb.DeleteProductRequest) (*pb.Empty, error) {
	resp, err := pb.NewProductServiceClient(grpc.ApiServerInstance.ProductServiceConn).DeleteProduct(ctx, product)
	return resp, err
}

func (p *ProductService) GetProductById(ctx context.Context, product *pb.GetProductRequest) (*pb.Product, error) {
	resp, err := pb.NewProductServiceClient(grpc.ApiServerInstance.ProductServiceConn).GetProduct(ctx, product)
	return resp, err
}
