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

func (s *ServiceProductMasuk) Create(input *domain.CreateProductMasukRequest) (*models.ModelProductMasuk, error) {
	var productmasuk domain.CreateProductMasukRequest

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

func (s *ServiceProductMasuk) Result(id string) (*models.ModelProductMasuk, error) {

	res, err := s.Repository.Result(id)

	return res, err

}

func (s *ServiceProductMasuk) Delete(id string) (*models.ModelProductMasuk, error) {

	res, err := s.Repository.Delete(id)

	return res, err
}

func (s *ServiceProductMasuk) Update(input *domain.UpdateProductMasukRequest) (*models.ModelProductMasuk, error) {
	var productmasuk domain.UpdateProductMasukRequest

	productmasuk.Name = input.Name
	productmasuk.Qty = input.Qty
	productmasuk.ProductID = input.ProductID
	productmasuk.SupplierID = input.SupplierID

	res, err := s.Repository.Update(&productmasuk)

	return res, err

}
