package domain

import "github.com/go-playground/validator/v10"

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required,lowercase" schema:"name"`
}

func (c *CreateCategoryRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}

type UpdateCategoryRequest struct {
	ID   string `json:"id" validate:"required,uuid"`
	Name string `json:"name" validate:"required,lowercase" schema:"name"`
}

func (c *UpdateCategoryRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}
