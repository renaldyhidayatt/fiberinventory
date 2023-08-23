package domain

import "github.com/go-playground/validator/v10"

type CreateProductMasukRequest struct {
	Name       string `json:"name" validate:"required"`
	Qty        string `json:"qty" validate:"required"`
	ProductID  string `json:"product_id" validate:"required"`
	SupplierID string `json:"supplier_id" validate:"required"`
}

func (c *CreateProductMasukRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}

type UpdateProductMasukRequest struct {
	ID         string `json:"id" validate:"uuid"`
	Name       string `json:"name" validate:"required"`
	Qty        string `json:"qty" validate:"required"`
	ProductID  string `json:"product_id" validate:"required"`
	SupplierID string `json:"supplier_id" validate:"required"`
}

func (c *UpdateProductMasukRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}
