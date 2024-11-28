package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	ID        string    `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	ProductID string    `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Total     float32   `json:"total"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	o.ID = uuid.NewString()
	return
}
