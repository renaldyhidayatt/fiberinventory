package domain

import "github.com/go-playground/validator/v10"

type ProductMasukInput struct {
	ID         string `json:"id" validate:"uuid"`
	Name       string `json:"name" validate:"required"`
	Qty        string `json:"qty" validate:"required"`
	ProductID  string `json:"product_id" validate:"required"`
	SupplierID string `json:"supplier_id" validate:"required"`
}

func (c *ProductMasukInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}
