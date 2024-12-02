package services

import (
	"product-service/internal/entities"
	"product-service/internal/repositories"
)

//go:generate mockgen -source=product_service.go -destination=mock_auth_service.go -package=mocks
type ProductService interface {
	CreateProduct(product *entities.Product) (*entities.Product, error)
	GetProduct(id string) (*entities.Product, error)
	ListProducts(page, limit int) ([]*entities.Product, error)
	UpdateProduct(product *entities.Product) (*entities.Product, error)
	DeleteProduct(id string) error
	GetProductStock(id string) (int32, error)
}

type productService struct {
	productRepo repositories.ProductRepository
}

func NewProductService(productRepo repositories.ProductRepository) *productService {
	return &productService{productRepo: productRepo}
}

func (s *productService) CreateProduct(product *entities.Product) (*entities.Product, error) {
	return s.productRepo.CreateProduct(product)
}

func (s *productService) GetProduct(id string) (*entities.Product, error) {
	return s.productRepo.GetProduct(id)
}

func (s *productService) ListProducts(page, limit int) ([]*entities.Product, error) {
	return s.productRepo.ListProducts(page, limit)
}

func (s *productService) UpdateProduct(product *entities.Product) (*entities.Product, error) {
	return s.productRepo.UpdateProduct(product)
}

func (s *productService) DeleteProduct(id string) error {
	return s.productRepo.DeleteProduct(id)
}

func (s *productService) GetProductStock(id string) (int32, error) {
	return s.productRepo.GetProductStock(id)
}
