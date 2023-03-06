package domain

import "github.com/go-playground/validator/v10"

type SupplierInput struct {
	ID      string `json:"id" validate:"required,uuid"`
	Name    string `json:"name" validate:"required,lowercase" schema:"name"`
	Alamat  string `json:"alamat" validate:"required,max=1000"`
	Email   string `json:"email" validate:"required,email"`
	Telepon string `json:"phone" validate:"required,gte=12"`
}

func (c *SupplierInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}