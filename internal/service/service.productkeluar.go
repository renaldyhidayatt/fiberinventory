package service

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/models"
	"fiberinventory/internal/repository"
)

type ServiceProductKeluar struct {
	Repository repository.ProductKeluarRepository
}

func NewServiceProductKeluar(productkeluar repository.ProductKeluarRepository) *ServiceProductKeluar {
	return &ServiceProductKeluar{Repository: productkeluar}
}

func (s *ServiceProductKeluar) Create(input *domain.CreateProductKeluarRequest) (*models.ModelProductKeluar, error) {
	var productkeluar domain.CreateProductKeluarRequest

	productkeluar.Qty = input.Qty
	productkeluar.ProductID = input.ProductID
	productkeluar.CategoryID = input.CategoryID

	res, err := s.Repository.Create(&productkeluar)

	return res, err
}

func (s *ServiceProductKeluar) Results() (*[]models.ModelProductKeluar, error) {
	res, err := s.Repository.Results()

	return res, err
}

func (s *ServiceProductKeluar) Result(id string) (*models.ModelProductKeluar, error) {

	res, err := s.Repository.Result(id)

	return res, err
}

func (s *ServiceProductKeluar) Delete(id string) (*models.ModelProductKeluar, error) {

	res, err := s.Repository.Delete(id)

	return res, err
}

func (s *ServiceProductKeluar) Update(input *domain.UpdateProductKeluarRequest) (*models.ModelProductKeluar, error) {
	var productkeluar domain.UpdateProductKeluarRequest

	productkeluar.Qty = input.Qty
	productkeluar.ProductID = input.ProductID
	productkeluar.CategoryID = input.CategoryID

	res, err := s.Repository.Update(&productkeluar)

	return res, err
}
