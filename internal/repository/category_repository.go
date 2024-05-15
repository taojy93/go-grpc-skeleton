package repo

import (
	"go-grpc-skeleton/internal/models"

	"gorm.io/gorm"
)

type ICategoryRepository interface {
	GetCategory(id int64) (*models.Category, error)
	ListCategory() ([]*models.Category, error)
}

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (r *CategoryRepository) GetCategory(id int64) (*models.Category, error) {
	var cate models.Category
	if err := r.db.First(&cate, id).Error; err != nil {
		return nil, err
	}
	return &cate, nil
}

func (r *CategoryRepository) ListCategory() ([]*models.Category, error) {
	cates := make([]*models.Category, 0)
	if err := r.db.Find(&cates).Error; err != nil {
		return nil, err
	}
	return cates, nil
}
