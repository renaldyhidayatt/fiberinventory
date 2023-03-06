package service

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/models"
	"fiberinventory/internal/repository"
)

type serviceCategory struct {
	Repository repository.CategoryRepository
}

func NewServiceCategory(category repository.CategoryRepository) *serviceCategory {
	return &serviceCategory{Repository: category}
}

func (s *serviceCategory) Create(input *domain.CategoryInput) (*models.ModelCategory, error) {
	var category domain.CategoryInput
	category.Name = input.Name

	res, err := s.Repository.Create(&category)
	return res, err
}

func (s *serviceCategory) Results() (*[]models.ModelCategory, error) {
	res, err := s.Repository.Results()

	return res, err
}

func (s *serviceCategory) Result(input *domain.CategoryInput) (*models.ModelCategory, error) {
	var category domain.CategoryInput
	category.ID = input.ID

	res, err := s.Repository.Result(&category)

	return res, err
}

func (s *serviceCategory) Delete(input *domain.CategoryInput) (*models.ModelCategory, error) {
	var category domain.CategoryInput
	category.ID = input.ID

	res, err := s.Repository.Delete(&category)

	return res, err
}

func (s *serviceCategory) Update(input *domain.CategoryInput) (*models.ModelCategory, error) {
	var category domain.CategoryInput
	category.Name = input.Name

	res, err := s.Repository.Update(&category)

	return res, err
}
