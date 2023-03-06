package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ModelProductKeluar struct {
	ID         string `json:"id" gorm:"primary_key"`
	Qty        string `json:"qty" gorm:"type:varchar; not null"`
	Product    ModelProduct
	ProductID  string `json:"product_id" gorm:"index"`
	Category   ModelCategory
	CategoryID string    `json:"category_id" gorm:"index"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (m *ModelProductKeluar) BeforeCreate(db *gorm.DB) error {
	m.ID = uuid.NewString()
	m.CreatedAt = time.Now().Local()
	return nil
}

func (m *ModelProductKeluar) BeforeUpdate(db *gorm.DB) error {
	m.ID = uuid.NewString()
	m.UpdatedAt = time.Now().Local()
	return nil
}
