package domain

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type UpdateUserRequest struct {
	ID        string `json:"id" validate:"uuid"`
	FirstName string `json:"first_name" validate:"required,lowercase"`
	LastName  string `json:"last_name" validate:"required,lowercase"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,gte=8"`
	Role      string `json:"role" validate:"required,lowercase"`
}

func (c *UpdateUserRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}

type UsersResponse struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
