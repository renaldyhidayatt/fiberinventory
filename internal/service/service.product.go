package service

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/models"
	"fiberinventory/internal/repository"
)

type ServiceProduct struct {
	Repository repository.ProductRepository
}

func NewServiceProduct(product repository.ProductRepository) *ServiceProduct {
	return &ServiceProduct{Repository: product}
}

func (s *ServiceProduct) Create(input *domain.ProductInput) (*models.ModelProduct, error) {
	var product domain.ProductInput
	product.Name = input.Name
	product.Image = input.Image
	product.Qty = input.Qty
	product.CategoryID = input.CategoryID

	res, err := s.Repository.Create(&product)

	return res, err
}

func (s *ServiceProduct) Results() (*[]models.ModelProduct, error) {
	res, err := s.Repository.Results()
	return res, err
}

func (s *ServiceProduct) Result(input *domain.ProductInput) (*models.ModelProduct, error) {
	var product domain.ProductInput
	product.ID = input.ID

	res, err := s.Repository.Result(&product)

	return res, err
}

func (s *ServiceProduct) Delete(input *domain.ProductInput) (*models.ModelProduct, error) {
	var product domain.ProductInput
	product.ID = input.ID

	res, err := s.Repository.Delete(&product)

	return res, err
}

func (s *ServiceProduct) Update(input *domain.ProductInput) (*models.ModelProduct, error) {
	var product domain.ProductInput

	product.Name = input.Name
	product.Image = input.Image
	product.Qty = input.Qty
	product.CategoryID = input.CategoryID

	res, err := s.Repository.Update(&product)

	return res, err
}
