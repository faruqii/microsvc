package services

import (
	"context"
	"fmt"
	"order-service/externals"
	"order-service/internal/entities"
	"order-service/internal/repositories"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type OrderService interface {
	CreateOrder(order *entities.Order) (*entities.Order, error)
	GetOrder(id string) (*entities.Order, error)
	ListOrders(page, limit int) ([]*entities.Order, error)
	UpdateOrder(order *entities.Order) (*entities.Order, error)
	DeleteOrder(id string) error
}

type orderService struct {
	orderRepo     repositories.OrderRepository
	productClient externals.ProductServiceClient
}

func NewOrderService(
	orderRepo repositories.OrderRepository,
	productClient externals.ProductServiceClient,
) *orderService {
	return &orderService{
		orderRepo:     orderRepo,
		productClient: productClient,
	}
}

func (s *orderService) CreateOrder(order *entities.Order) (*entities.Order, error) {
	ctx := context.Background()

	// Validate product existence and stock
	product, err := s.checkProductAvailability(ctx, order.ProductID, order.Quantity)
	if err != nil {
		return nil, err
	}

	// Update product stock
	if err := s.updateProductStock(ctx, product, order.Quantity); err != nil {
		return nil, err
	}

	// Calculate total price
	order.Total = product.Price * float32(order.Quantity)

	// Step 4: Create the order
	createdOrder, err := s.orderRepo.CreateOrder(order)
	if err != nil {
		return nil, fmt.Errorf("failed to create order: %v", err)
	}

	return createdOrder, nil
}

// Helper to check product availability
func (s *orderService) checkProductAvailability(ctx context.Context, productID string, quantity int) (*externals.Product, error) {
	productRequest := &externals.GetProductRequest{Id: productID}
	productResponse, err := s.productClient.GetProduct(ctx, productRequest)
	if err != nil {
		if grpc.Code(err) == codes.NotFound {
			return nil, fmt.Errorf("product with ID %s not found", productID)
		}
		return nil, fmt.Errorf("failed to check product existence: %v", err)
	}

	if productResponse == nil || productResponse.Product == nil {
		return nil, fmt.Errorf("product with ID %s does not exist", productID)
	}

	if productResponse.Product.Stock < int32(quantity) {
		return nil, fmt.Errorf("product stock is insufficient")
	}

	return productResponse.Product, nil
}

// Helper to update product stock
func (s *orderService) updateProductStock(ctx context.Context, product *externals.Product, quantity int) error {
	_, err := s.productClient.UpdateProduct(ctx, &externals.UpdateProductRequest{
		Id:    product.Id,
		Stock: product.Stock - int32(quantity),
	})
	if err != nil {
		return fmt.Errorf("failed to update product stock: %v", err)
	}

	return nil
}

func (s *orderService) GetOrder(id string) (*entities.Order, error) {
	return s.orderRepo.GetOrder(id)
}

func (s *orderService) ListOrders(page, limit int) ([]*entities.Order, error) {
	return s.orderRepo.ListOrders(page, limit)
}

func (s *orderService) UpdateOrder(order *entities.Order) (*entities.Order, error) {
	return s.orderRepo.UpdateOrder(order)
}

func (s *orderService) DeleteOrder(id string) error {
	return s.orderRepo.DeleteOrder(id)
}
