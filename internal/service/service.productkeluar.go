package service

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/models"
	"fiberinventory/internal/repository"
)

type serviceProductKeluar struct {
	Repository repository.ProductKeluarRepository
}

func NewServiceProductKeluar(productkeluar repository.ProductKeluarRepository) *serviceProductKeluar {
	return &serviceProductKeluar{Repository: productkeluar}
}

func (s *serviceProductKeluar) Create(input *domain.ProductKeluarInput) (*models.ModelProductKeluar, error) {
	var productkeluar domain.ProductKeluarInput

	productkeluar.Qty = input.Qty
	productkeluar.ProductID = input.ProductID
	productkeluar.CategoryID = input.CategoryID

	res, err := s.Repository.Create(&productkeluar)

	return res, err
}

func (s *serviceProductKeluar) Results() (*[]models.ModelProductKeluar, error) {
	res, err := s.Repository.Results()

	return res, err
}

func (s *serviceProductKeluar) Result(input *domain.ProductKeluarInput) (*models.ModelProductKeluar, error) {
	var productkeluar domain.ProductKeluarInput
	productkeluar.ID = input.ID

	res, err := s.Repository.Result(&productkeluar)

	return res, err
}

func (s *serviceProductKeluar) Delete(input *domain.ProductKeluarInput) (*models.ModelProductKeluar, error) {
	var productkeluar domain.ProductKeluarInput
	productkeluar.ID = input.ID

	res, err := s.Repository.Delete(&productkeluar)

	return res, err
}

func (s *serviceProductKeluar) Update(input *domain.ProductKeluarInput) (*models.ModelProductKeluar, error) {
	var productkeluar domain.ProductKeluarInput

	productkeluar.Qty = input.Qty
	productkeluar.ProductID = input.ProductID
	productkeluar.CategoryID = input.CategoryID

	res, err := s.Repository.Update(&productkeluar)

	return res, err
}
