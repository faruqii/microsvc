package repositories

import (
	"errors"
	"reflect"

	"gorm.io/gorm"
)

type Repositories[T any] interface {
	Create(entities *T) (*T, error)
	Get(id string) (*T, error)
	List(page, limit int) ([]*T, error)
	Update(entities *T) (*T, error)
	Delete(id string) error
}

type RepositoryImpl[T any] struct {
	db *gorm.DB
}

func NewRepository[T any](db *gorm.DB) *RepositoryImpl[T] {
	return &RepositoryImpl[T]{db: db}
}

func (r *RepositoryImpl[T]) Create(entities *T) (*T, error) {
	if err := r.db.Create(entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *RepositoryImpl[T]) Get(id string) (*T, error) {
	var entity T
	if err := r.db.Where("id = ?", id).First(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *RepositoryImpl[T]) List(page, limit int) ([]*T, error) {
	var entities []*T
	if err := r.db.Offset(page).Limit(limit).Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *RepositoryImpl[T]) Update(entities *T) (*T, error) {
	// using reflection to get the ID field as we use generic type
	value := reflect.ValueOf(entities).Elem()
	idField := value.FieldByName("ID")
	if !idField.IsValid() {
		return nil, errors.New("ID field not found")
	}
	// update the entity
	if err := r.db.Where("id = ?", idField.Interface()).Updates(entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *RepositoryImpl[T]) Delete(id string) error {
	var entity T
	if err := r.db.Where("id = ?", id).Delete(&entity).Error; err != nil {
		return err
	}
	return nil
}
