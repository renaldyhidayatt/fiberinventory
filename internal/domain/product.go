package domain

import "github.com/go-playground/validator/v10"

type ProductInput struct {
	ID         string `json:"id" validate:"uuid"`
	Name       string `json:"name" validate:"required,lowercase"`
	Image      string `json:"image" validate:"required"`
	Qty        string `json:"qty" validate:"required"`
	CategoryID string `json:"category_id" validate:"required"`
}

func (c *ProductInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}
