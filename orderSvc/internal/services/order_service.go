package services

import (
	"context"
	"fmt"
	"log"
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
	// check if product exists
	ctx := context.Background()
	productRequest := &externals.GetProductRequest{
		Id: order.ProductID,
	}

	productResponse, err := s.productClient.GetProduct(ctx, productRequest)
	if err != nil {
		if grpc.Code(err) == codes.NotFound {
			return nil, fmt.Errorf("product with ID %s not found", order.ProductID)
		}
		return nil, fmt.Errorf("failed to check product existence: %v", err)
	}

	log.Printf("productResponse: %v", productResponse)

	if productResponse == nil || productResponse.Product == nil {
		return nil, fmt.Errorf("product with ID %s does not exist", order.ProductID)
	}

	// check if product stock is enough
	if productResponse.Product.Stock < int32(order.Quantity) {
		return nil, fmt.Errorf("product stock is not enough")
	}

	// update product stock
	productUpdateRequest, err := s.productClient.UpdateProduct(ctx, &externals.UpdateProductRequest{
		Id:    order.ProductID,
		Stock: productResponse.Product.Stock - int32(order.Quantity),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update product stock: %v", err)
	}

	if productUpdateRequest == nil || productUpdateRequest.Product == nil {
		return nil, fmt.Errorf("failed to update product stock")
	}

	// calculate total price
	order.Total = productResponse.Product.Price * float32(order.Quantity)

	// Continue with order creation logic
	createdOrder, err := s.orderRepo.CreateOrder(order)
	if err != nil {
		return nil, fmt.Errorf("failed to create order: %v", err)
	}

	return createdOrder, nil
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
