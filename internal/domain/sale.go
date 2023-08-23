package domain

import "github.com/go-playground/validator/v10"

type CreateSaleRequest struct {
	Name    string `json:"name" validate:"required,lowercase" schema:"name"`
	Alamat  string `json:"alamat" validate:"required,max=1000"`
	Email   string `json:"email" validate:"required,email"`
	Telepon string `json:"phone" validate:"required,gte=12"`
}

func (c *CreateSaleRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}

type UpdateSaleRequest struct {
	ID      string `json:"id" validate:"required,uuid"`
	Name    string `json:"name" validate:"required,lowercase" schema:"name"`
	Alamat  string `json:"alamat" validate:"required,max=1000"`
	Email   string `json:"email" validate:"required,email"`
	Telepon string `json:"phone" validate:"required,gte=12"`
}

func (c *UpdateSaleRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}
