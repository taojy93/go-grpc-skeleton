package service

import (
	"go-grpc-skeleton/internal/models"
	repo "go-grpc-skeleton/internal/repository"
)

type ICategoryService interface {
	GetCategory(id int64) (*models.Category, error)
	ListCategorys() ([]*models.Category, error)
}

type CategoryService struct {
	repo repo.ICategoryRepository
}

func NewCategoryService(repo repo.ICategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetCategory(id int64) (*models.Category, error) {
	return s.repo.GetCategory(id)
}

func (s *CategoryService) ListCategorys() ([]*models.Category, error) {
	return s.repo.ListCategory()
}
