package service

import (
	"go-grpc-skeleton/internal/models"
	repo "go-grpc-skeleton/internal/repository"
)

type IProductService interface {
	GetProduct(id int64) (*models.Product, error)
	ListProducts() ([]*models.Product, error)
}

type ProductService struct {
	repo repo.IProductRepository
}

func NewProductService(repo repo.IProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetProduct(id int64) (*models.Product, error) {
	return s.repo.GetProduct(id)
}

func (s *ProductService) ListProducts() ([]*models.Product, error) {
	return s.repo.ListProducts()
}
