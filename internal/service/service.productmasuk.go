package service

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/models"
	"fiberinventory/internal/repository"
)

type ServiceProductMasuk struct {
	Repository repository.ProductMasukRepository
}

func NewServiceProductMasuk(productmasuk repository.ProductMasukRepository) *ServiceProductMasuk {
	return &ServiceProductMasuk{Repository: productmasuk}
}

func (s *ServiceProductMasuk) Create(input *domain.ProductMasukInput) (*models.ModelProductMasuk, error) {
	var productmasuk domain.ProductMasukInput

	productmasuk.Name = input.Name
	productmasuk.Qty = input.Qty
	productmasuk.ProductID = input.ProductID
	productmasuk.SupplierID = input.SupplierID

	res, err := s.Repository.Create(&productmasuk)

	return res, err
}

func (s *ServiceProductMasuk) Results() (*[]models.ModelProductMasuk, error) {
	res, err := s.Repository.Results()

	return res, err
}

func (s *ServiceProductMasuk) Result(input *domain.ProductMasukInput) (*models.ModelProductMasuk, error) {
	var productmasuk domain.ProductMasukInput
	productmasuk.ID = input.ID

	res, err := s.Repository.Result(&productmasuk)

	return res, err

}

func (s *ServiceProductMasuk) Delete(input *domain.ProductMasukInput) (*models.ModelProductMasuk, error) {
	var productmasuk domain.ProductMasukInput
	productmasuk.ID = input.ID

	res, err := s.Repository.Delete(&productmasuk)

	return res, err
}

func (s *ServiceProductMasuk) Update(input *domain.ProductMasukInput) (*models.ModelProductMasuk, error) {
	var productmasuk domain.ProductMasukInput

	productmasuk.Name = input.Name
	productmasuk.Qty = input.Qty
	productmasuk.ProductID = input.ProductID
	productmasuk.SupplierID = input.SupplierID

	res, err := s.Repository.Update(&productmasuk)

	return res, err

}
