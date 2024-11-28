package repositories

import (
	"order-service/internal/entities"

	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order *entities.Order) (*entities.Order, error)
	GetOrder(id string) (*entities.Order, error)
	ListOrders(page, limit int) ([]*entities.Order, error)
	UpdateOrder(order *entities.Order) (*entities.Order, error)
	DeleteOrder(id string) error
}

type orderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepositoryImpl {
	return &orderRepositoryImpl{db: db}
}

func (r *orderRepositoryImpl) CreateOrder(order *entities.Order) (*entities.Order, error) {
	if err := r.db.Create(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (r *orderRepositoryImpl) GetOrder(id string) (*entities.Order, error) {
	var order entities.Order
	if err := r.db.Where("id = ?", id).First(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *orderRepositoryImpl) ListOrders(page, limit int) ([]*entities.Order, error) {
	var orders []*entities.Order
	if err := r.db.Offset(page).Limit(limit).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *orderRepositoryImpl) UpdateOrder(order *entities.Order) (*entities.Order, error) {
	if err := r.db.Save(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (r *orderRepositoryImpl) DeleteOrder(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&entities.Order{}).Error; err != nil {
		return err
	}
	return nil
}
