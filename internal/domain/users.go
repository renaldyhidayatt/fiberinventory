package domain

import "github.com/go-playground/validator/v10"

type UserInput struct {
	ID        string `json:"id" validate:"uuid"`
	FirstName string `json:"first_name" validate:"required,lowercase"`
	LastName  string `json:"last_name" validate:"required,lowercase"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,gte=8"`
	Role      string `json:"role" validate:"required,lowercase"`
}

func (c *UserInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}
