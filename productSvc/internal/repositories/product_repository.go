package repositories

import (
	"product-service/internal/entities"

	"gorm.io/gorm"
)

//go:generate mockgen -source=product_repository.go -destination=mock_product_repository.go -package=mocks
type ProductRepository interface {
	CreateProduct(product *entities.Product) (*entities.Product, error)
	GetProduct(id string) (*entities.Product, error)
	ListProducts(page, limit int) ([]*entities.Product, error)
	UpdateProduct(product *entities.Product) (*entities.Product, error)
	DeleteProduct(id string) error
	GetProductStock(id string) (int32, error)
}

type productRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepositoryImpl {
	return &productRepositoryImpl{db: db}
}

func (r *productRepositoryImpl) CreateProduct(product *entities.Product) (*entities.Product, error) {
	if err := r.db.Create(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productRepositoryImpl) GetProduct(id string) (*entities.Product, error) {
	var product entities.Product
	if err := r.db.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepositoryImpl) ListProducts(page, limit int) ([]*entities.Product, error) {
	var products []*entities.Product
	if err := r.db.Offset(page).Limit(limit).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepositoryImpl) UpdateProduct(product *entities.Product) (*entities.Product, error) {
	if err := r.db.Where("id = ?", product.ID).Updates(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productRepositoryImpl) DeleteProduct(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&entities.Product{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *productRepositoryImpl) GetProductStock(id string) (int32, error) {
	var product entities.Product
	if err := r.db.Where("id = ?", id).First(&product).Error; err != nil {
		return 0, err
	}
	return product.Stock, nil
}
