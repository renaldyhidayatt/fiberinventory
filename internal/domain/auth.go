package domain

import "github.com/go-playground/validator/v10"

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8"`
}

func (l *LoginInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(l)

	if err != nil {
		return err
	}

	return nil
}

type RegisterInput struct {
	FirstName string `json:"first_name" validate:"required,lowercase"`
	LastName  string `json:"last_name" validate:"required,lowercase"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,gte=8"`
	Role      string `json:"role" validate:"required,lowercase"`
}

func (c *RegisterInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}
