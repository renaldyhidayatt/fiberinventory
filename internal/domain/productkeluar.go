package domain

import "github.com/go-playground/validator/v10"

type ProductKeluarInput struct {
	ID         string `json:"id" validate:"uuid"`
	Qty        string `json:"qty" validate:"required"`
	ProductID  string `json:"product_id" validate:"required"`
	CategoryID string `json:"category_id" validate:"required"`
}

func (c *ProductKeluarInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}
