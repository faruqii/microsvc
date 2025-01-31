package repositories

import (
	"product-service/internal/entities"

	"gorm.io/gorm"
)

//go:generate mockgen -source=product_repository.go -destination=mock_product_repository.go -package=mocks
type ProductRepository interface {
	Repositories[entities.Product]
	GetProductStock(id string) (int32, error)
}

type productRepositoryImpl struct {
	Repositories[entities.Product]
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepositoryImpl{
		Repositories: NewRepository[entities.Product](db),
		db: db,
	}
}

func (r *productRepositoryImpl) GetProductStock(id string) (int32, error) {
	var product entities.Product
	if err := r.db.Where("id = ?", id).First(&product).Error; err != nil {
		return 0, err
	}
	return product.Stock, nil
}
