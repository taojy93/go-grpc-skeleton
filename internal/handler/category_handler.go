package handler

import (
	"context"
	"go-grpc-skeleton/internal/service"
	pb "go-grpc-skeleton/proto"
)

type CategoryHandler struct {
	pb.UnimplementedCategoryServiceServer
	service service.ICategoryService
}

func NewCategoryHandler(service service.ICategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (h *CategoryHandler) GetCategory(ctx context.Context, req *pb.GetCategoryRequest) (*pb.GetCategoryResponse, error) {
	category, err := h.service.GetCategory(req.Id)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, nil
	}
	return &pb.GetCategoryResponse{
		Category: &pb.Category{
			Id:   int64(category.ID),
			Name: category.Name,
		},
	}, nil
}

func (h *CategoryHandler) ListCategorys(ctx context.Context, req *pb.ListCategoryRequest) (*pb.ListCategoryResponse, error) {
	categorys, err := h.service.ListCategorys()
	if err != nil {
		return nil, err
	}
	response := &pb.ListCategoryResponse{}
	for _, category := range categorys {
		response.Cates = append(response.Cates, &pb.Category{
			Id:   int64(category.ID),
			Name: category.Name,
		})
	}
	return response, nil
}
