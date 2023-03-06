package service

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/models"
	"fiberinventory/internal/repository"
)

type serviceProduct struct {
	Repository repository.ProductRepository
}

func NewServiceProduct(product repository.ProductRepository) *serviceProduct {
	return &serviceProduct{Repository: product}
}

func (s *serviceProduct) Create(input *domain.ProductInput) (*models.ModelProduct, error) {
	var product domain.ProductInput
	product.Name = input.Name
	product.Image = input.Image
	product.Qty = input.Qty
	product.CategoryID = input.CategoryID

	res, err := s.Repository.Create(&product)

	return res, err
}

func (s *serviceProduct) Results() (*[]models.ModelProduct, error) {
	res, err := s.Repository.Results()
	return res, err
}

func (s *serviceProduct) Result(input *domain.ProductInput) (*models.ModelProduct, error) {
	var product domain.ProductInput
	product.ID = input.ID

	res, err := s.Repository.Result(&product)

	return res, err
}

func (s *serviceProduct) Delete(input *domain.ProductInput) (*models.ModelProduct, error) {
	var product domain.ProductInput
	product.ID = input.ID

	res, err := s.Repository.Delete(&product)

	return res, err
}

func (s *serviceProduct) Update(input *domain.ProductInput) (*models.ModelProduct, error) {
	var product domain.ProductInput

	product.Name = input.Name
	product.Image = input.Image
	product.Qty = input.Qty
	product.CategoryID = input.CategoryID

	res, err := s.Repository.Update(&product)

	return res, err
}
