package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ModelCustomer struct {
	ID        string    `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" gorm:"type:varchar; not null"`
	Alamat    string    `json:"alamat" gorm:"type:varchar; not null"`
	Email     string    `json:"email" gorm:"type:varchar; not null"`
	Telepon   string    `json:"telepon" gorm:"type:bigint; unique; not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *ModelCustomer) BeforeCreate(db *gorm.DB) error {
	m.ID = uuid.NewString()
	m.CreatedAt = time.Now().Local()
	return nil
}

func (m *ModelCustomer) BeforeUpdate(db *gorm.DB) error {
	m.ID = uuid.NewString()
	m.UpdatedAt = time.Now().Local()
	return nil
}
