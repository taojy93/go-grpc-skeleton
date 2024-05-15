package handler

import (
	"context"

	"go-grpc-skeleton/internal/service"
	pb "go-grpc-skeleton/proto"
)

type ProductHandler struct {
	pb.UnimplementedProductServiceServer
	service service.IProductService
}

func NewProductHandler(service service.IProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	product, err := h.service.GetProduct(req.Id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, nil
	}
	return &pb.GetProductResponse{
		Product: &pb.Product{
			Id:    int64(product.ID),
			Name:  product.Name,
			Price: product.Price,
		},
	}, nil
}

func (h *ProductHandler) ListProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	products, err := h.service.ListProducts()
	if err != nil {
		return nil, err
	}
	response := &pb.ListProductsResponse{}
	for _, product := range products {
		response.Products = append(response.Products, &pb.Product{
			Id:    int64(product.ID),
			Name:  product.Name,
			Price: product.Price,
		})
	}
	return response, nil
}

func (h *ProductHandler) StreamProducts(req *pb.StreamProductsRequest, stream pb.ProductService_StreamProductsServer) error {
	products, err := h.service.ListProducts()
	if err != nil {
		return err
	}
	for _, product := range products {
		res := &pb.StreamProductsResponse{
			Product: &pb.Product{
				Id:    int64(product.ID),
				Name:  product.Name,
				Price: product.Price,
			},
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
	return nil
}
