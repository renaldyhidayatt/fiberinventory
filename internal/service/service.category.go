package service

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/models"
	"fiberinventory/internal/repository"
)

type ServiceCategory struct {
	Repository repository.CategoryRepository
}

func NewServiceCategory(category repository.CategoryRepository) *ServiceCategory {
	return &ServiceCategory{Repository: category}
}

func (s *ServiceCategory) Create(input *domain.CategoryInput) (*models.ModelCategory, error) {
	var category domain.CategoryInput
	category.Name = input.Name

	res, err := s.Repository.Create(&category)
	return res, err
}

func (s *ServiceCategory) Results() (*[]models.ModelCategory, error) {
	res, err := s.Repository.Results()

	return res, err
}

func (s *ServiceCategory) Result(input *domain.CategoryInput) (*models.ModelCategory, error) {
	var category domain.CategoryInput
	category.ID = input.ID

	res, err := s.Repository.Result(&category)

	return res, err
}

func (s *ServiceCategory) Delete(input *domain.CategoryInput) (*models.ModelCategory, error) {
	var category domain.CategoryInput
	category.ID = input.ID

	res, err := s.Repository.Delete(&category)

	return res, err
}

func (s *ServiceCategory) Update(input *domain.CategoryInput) (*models.ModelCategory, error) {
	var category domain.CategoryInput
	category.Name = input.Name

	res, err := s.Repository.Update(&category)

	return res, err
}
