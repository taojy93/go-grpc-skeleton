package repo

import (
	"go-grpc-skeleton/internal/models"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type IProductRepository interface {
	GetProduct(id int64) (*models.Product, error)
	ListProducts() ([]*models.Product, error)
}

type ProductRepository struct {
	db  *gorm.DB
	rds *redis.Client
	es  *elasticsearch.Client
}

func NewProductRepository(db *gorm.DB, redis *redis.Client, es *elasticsearch.Client) *ProductRepository {
	return &ProductRepository{
		db:  db,
		es:  es,
		rds: redis,
	}
}

func (r *ProductRepository) GetProduct(id int64) (*models.Product, error) {
	var product models.Product
	if err := r.db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) ListProducts() ([]*models.Product, error) {
	products := make([]*models.Product, 0)
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
