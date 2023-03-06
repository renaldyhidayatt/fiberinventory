package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ModelUser struct {
	ID        string    `json:"id" gorm:"primary_key"`
	FirstName string    `json:"first_name" gorm:"type:varchar;  not null"`
	LastName  string    `json:"last_name" gorm:"type:varchar; not null"`
	Email     string    `json:"email" gorm:"type:varchar; unique; not null"`
	Password  string    `json:"password" gorm:"type:varchar; not null"`
	Role      string    `json:"role" gorm:"type:varchar; not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *ModelUser) BeforeCreate(db *gorm.DB) error {

	m.ID = uuid.NewString()
	m.CreatedAt = time.Now()
	return nil
}

func (m *ModelUser) BeforeUpdate(db *gorm.DB) error {

	m.ID = uuid.NewString()
	m.UpdatedAt = time.Now()
	return nil
}
