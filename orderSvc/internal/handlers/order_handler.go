package handlers

import (
	"context"
	"order-service/internal/entities"
	"order-service/internal/protobuff"
	"order-service/internal/services"
	"time"
)

type ProductHandler struct {
	protobuff.UnimplementedOrderServiceServer
	svc services.OrderService
}

func NewProductHandler(svc services.OrderService) *ProductHandler {
	return &ProductHandler{svc: svc}
}

func (h *ProductHandler) CreateOrder(ctx context.Context, req *protobuff.CreateOrderRequest) (*protobuff.CreateOrderResponse, error) {
	order := &entities.Order{
		ProductID: req.ProductId,
		Quantity:  int(req.Quantity),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createdOrder, err := h.svc.CreateOrder(order)
	if err != nil {
		return nil, err
	}

	return &protobuff.CreateOrderResponse{
		Message: "Order created successfully",
		Order: &protobuff.Order{
			Id:        createdOrder.ID,
			ProductId: createdOrder.ProductID,
			Quantity:  int32(createdOrder.Quantity),
			Total:     createdOrder.Total,
			CreatedAt: createdOrder.CreatedAt.String(),
			UpdatedAt: createdOrder.UpdatedAt.String(),
		},
	}, nil
}
