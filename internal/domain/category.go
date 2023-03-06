package domain

import "github.com/go-playground/validator/v10"

type CategoryInput struct {
	ID   string `json:"id" validate:"required,uuid"`
	Name string `json:"name" validate:"required,lowercase" schema:"name"`
}

func (c *CategoryInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}
