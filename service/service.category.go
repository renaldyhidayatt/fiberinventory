package service

import (
	"fiberinventory/entity"
	"fiberinventory/models"
	"fiberinventory/schemas"
)

type serviceCategory struct {
	category entity.EntityCategory
}

func NewServiceCategory(category entity.EntityCategory) *serviceCategory {
	return &serviceCategory{category: category}
}

func (s *serviceCategory) EntityCreate(input *schemas.SchemaCategory) (*models.ModelCategory, schemas.SchemaDatabaseError) {
	var category schemas.SchemaCategory
	category.Name = input.Name

	res, err := s.category.EntityCreate(&category)
	return res, err
}

func (s *serviceCategory) EntityResults() (*[]models.ModelCategory, schemas.SchemaDatabaseError) {
	res, err := s.category.EntityResults()

	return res, err
}

func (s *serviceCategory) EntityResult(input *schemas.SchemaCategory) (*models.ModelCategory, schemas.SchemaDatabaseError) {
	var category schemas.SchemaCategory
	category.ID = input.ID

	res, err := s.category.EntityResult(&category)

	return res, err
}

func (s *serviceCategory) EntityDelete(input *schemas.SchemaCategory) (*models.ModelCategory, schemas.SchemaDatabaseError) {
	var category schemas.SchemaCategory
	category.ID = input.ID

	res, err := s.category.EntityDelete(&category)

	return res, err
}

func (s *serviceCategory) EntityUpdate(input *schemas.SchemaCategory) (*models.ModelCategory, schemas.SchemaDatabaseError) {
	var category schemas.SchemaCategory
	category.Name = input.Name

	res, err := s.category.EntityUpdate(&category)

	return res, err
}
