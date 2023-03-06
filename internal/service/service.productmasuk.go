package service

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/models"
	"fiberinventory/internal/repository"
)

type serviceProductMasuk struct {
	Repository repository.ProductMasukRepository
}

func NewServiceProductMasuk(productmasuk repository.ProductMasukRepository) *serviceProductMasuk {
	return &serviceProductMasuk{Repository: productmasuk}
}

func (s *serviceProductMasuk) Create(input *domain.ProductMasukInput) (*models.ModelProductMasuk, error) {
	var productmasuk domain.ProductMasukInput

	productmasuk.Name = input.Name
	productmasuk.Qty = input.Qty
	productmasuk.ProductID = input.ProductID
	productmasuk.SupplierID = input.SupplierID

	res, err := s.Repository.Create(&productmasuk)

	return res, err
}

func (s *serviceProductMasuk) Results() (*[]models.ModelProductMasuk, error) {
	res, err := s.Repository.Results()

	return res, err
}

func (s *serviceProductMasuk) Result(input *domain.ProductMasukInput) (*models.ModelProductMasuk, error) {
	var productmasuk domain.ProductMasukInput
	productmasuk.ID = input.ID

	res, err := s.Repository.Result(&productmasuk)

	return res, err

}

func (s *serviceProductMasuk) Delete(input *domain.ProductMasukInput) (*models.ModelProductMasuk, error) {
	var productmasuk domain.ProductMasukInput
	productmasuk.ID = input.ID

	res, err := s.Repository.Delete(&productmasuk)

	return res, err
}

func (s *serviceProductMasuk) Update(input *domain.ProductMasukInput) (*models.ModelProductMasuk, error) {
	var productmasuk domain.ProductMasukInput

	productmasuk.Name = input.Name
	productmasuk.Qty = input.Qty
	productmasuk.ProductID = input.ProductID
	productmasuk.SupplierID = input.SupplierID

	res, err := s.Repository.Update(&productmasuk)

	return res, err

}
