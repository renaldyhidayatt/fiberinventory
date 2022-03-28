package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ModelProductMasuk struct {
	ID         string `json:"id" gorm:"primary_key"`
	Name       string `json:"name" gorm:"type:varchar; not null"`
	Qty        string `json:"qty" gorm:"type:varchar; not null"`
	Product    ModelProduct
	ProductID  string `json:"product_id" gorm:"index"`
	Supplier   ModelSupplier
	SupplierID string    `json:"supplier_id" gorm:"index"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (m *ModelProductMasuk) BeforeCreate(db *gorm.DB) error {
	m.ID = uuid.NewString()
	m.CreatedAt = time.Now().Local()
	return nil
}

func (m *ModelProductMasuk) BeforeUpdate(db *gorm.DB) error {
	m.ID = uuid.NewString()
	m.UpdatedAt = time.Now().Local()
	return nil
}
