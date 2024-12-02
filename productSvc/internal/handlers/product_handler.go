package handlers

import (
	"context"
	"log"
	"product-service/internal/entities"
	"product-service/internal/protobuff"

	"product-service/internal/services"
	"time"
)

type ProductHandler struct {
	protobuff.UnimplementedProductServiceServer
	svc services.ProductService
}

func NewProductHandler(svc services.ProductService) *ProductHandler {
	return &ProductHandler{svc: svc}
}

func (h *ProductHandler) CreateProduct(ctx context.Context, req *protobuff.CreateProductRequest) (*protobuff.CreateProductResponse, error) {
	product := &entities.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	createdProduct, err := h.svc.CreateProduct(product)
	if err != nil {
		return nil, err
	}

	return &protobuff.CreateProductResponse{
		Message: "Product created successfully",
		Product: &protobuff.Product{
			Id:          createdProduct.ID,
			Name:        createdProduct.Name,
			Description: createdProduct.Description,
			Price:       createdProduct.Price,
			Stock:       createdProduct.Stock,
			CreatedAt:   createdProduct.CreatedAt.String(),
			UpdatedAt:   createdProduct.UpdatedAt.String(),
		},
	}, nil
}

func (h *ProductHandler) GetProduct(ctx context.Context, req *protobuff.GetProductRequest) (*protobuff.GetProductResponse, error) {
	product, err := h.svc.GetProduct(req.Id)
	if err != nil {
		return nil, err
	}

	return &protobuff.GetProductResponse{
		Product: &protobuff.Product{
			Id:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
			CreatedAt:   product.CreatedAt.String(),
			UpdatedAt:   product.UpdatedAt.String(),
		},
	}, nil
}

func (h *ProductHandler) ListProducts(ctx context.Context, req *protobuff.ListProductRequest) (*protobuff.ListProductResponse, error) {
	products, err := h.svc.ListProducts(int(req.Page), int(req.Limit))
	if err != nil {
		return nil, err
	}

	var productsResp []*protobuff.Product
	for _, product := range products {
		productsResp = append(productsResp, &protobuff.Product{
			Id:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			CreatedAt:   product.CreatedAt.String(),
			UpdatedAt:   product.UpdatedAt.String(),
		})
	}

	return &protobuff.ListProductResponse{
		Message:  "List of products",
		Products: productsResp,
	}, nil
}

func (h *ProductHandler) UpdateProduct(ctx context.Context, req *protobuff.UpdateProductRequest) (*protobuff.UpdateProductResponse, error) {
	// log to check if the request is coming through from the client
	log.Printf("UpdateProduct request: %v", req)

	product := &entities.Product{
		ID:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		Stock:       req.Stock,
		Price:       req.Price,
		UpdatedAt:   time.Now(),
	}

	updatedProduct, err := h.svc.UpdateProduct(product)
	if err != nil {
		return nil, err
	}

	return &protobuff.UpdateProductResponse{
		Message: "Product updated successfully",
		Product: &protobuff.Product{
			Id:          updatedProduct.ID,
			Name:        updatedProduct.Name,
			Description: updatedProduct.Description,
			Price:       updatedProduct.Price,
			CreatedAt:   updatedProduct.CreatedAt.String(),
			UpdatedAt:   updatedProduct.UpdatedAt.String(),
		},
	}, nil
}

func (h *ProductHandler) DeleteProduct(ctx context.Context, req *protobuff.DeleteProductRequest) (*protobuff.DeleteProductResponse, error) {
	err := h.svc.DeleteProduct(req.Id)
	if err != nil {
		return nil, err
	}

	return &protobuff.DeleteProductResponse{
		Message: "Product deleted successfully",
	}, nil
}

func (h *ProductHandler) GetProductStock(ctx context.Context, req *protobuff.GetProductStockRequest) (*protobuff.GetProductStockResponse, error) {
	stock, err := h.svc.GetProductStock(req.Id)
	if err != nil {
		return nil, err
	}

	return &protobuff.GetProductStockResponse{

		Stock: int32(stock),
	}, nil
}
