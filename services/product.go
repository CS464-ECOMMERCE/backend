package services

import (
	"backend/grpc"
	pb "backend/proto"
	"context"
	"errors"
	"io"
	"mime/multipart"
)

type ProductService struct {
}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (p *ProductService) GetProduct(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	resp, err := pb.NewProductServiceClient(grpc.ApiServerInstance.ProductServiceConn).ListProducts(ctx, req)
	return resp, err
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

func (p *ProductService) UpdateProductImages(ctx context.Context, files multipart.Form, id uint64) (*pb.UpdateProductImagesResponse, error) {

	stream, err := pb.NewProductServiceClient(grpc.ApiServerInstance.ProductServiceConn).UpdateProductImages(ctx)
	if err != nil {
		return &pb.UpdateProductImagesResponse{}, errors.New("failed create stream")
	}
	for _, fileHeader := range files.File["images"] {
		file, err := fileHeader.Open()
		if err != nil {
			return &pb.UpdateProductImagesResponse{}, errors.New("failed to open file")
		}

		buffer := make([]byte, 1024)
		for {
			n, err := file.Read(buffer)
			if err == io.EOF {
				break
			}
			if err != nil {
				return &pb.UpdateProductImagesResponse{}, errors.New("failed to read file")
			}

			err = stream.Send(&pb.UpdateProductImagesRequest{
				ImageData: buffer[:n],
				Filename:  fileHeader.Filename,
				Id:        id,
			})
			if err != nil {
				return &pb.UpdateProductImagesResponse{}, errors.New("failed to send file")
			}

		}
		file.Close()
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		return &pb.UpdateProductImagesResponse{}, errors.New("failed to close stream")
	}
	return resp, nil
}

func (p *ProductService) GetProductByMerchantId(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	resp, err := pb.NewProductServiceClient(grpc.ApiServerInstance.ProductServiceConn).ListProducts(ctx, req)
	return resp, err
}