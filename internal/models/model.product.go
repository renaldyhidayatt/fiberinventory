package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ModelProduct struct {
	ID         string `json:"id" gorm:"primary_key"`
	Name       string `json:"name" gorm:"type:varchar; not null"`
	Image      string `json:"image" gorm:"type:varchar; not null"`
	Qty        string `json:"qty" gorm:"type:varchar; not null"`
	Category   ModelCategory
	CategoryID string    `json:"category_id" gorm:"index"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (m *ModelProduct) BeforeCreate(db *gorm.DB) error {
	m.ID = uuid.NewString()
	m.CreatedAt = time.Now().Local()
	return nil
}

func (m *ModelProduct) BeforeUpdate(db *gorm.DB) error {
	m.ID = uuid.NewString()
	m.UpdatedAt = time.Now().Local()

	return nil
}
