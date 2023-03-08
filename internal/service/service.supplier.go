package service

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/models"
	"fiberinventory/internal/repository"
)

type ServiceSupplier struct {
	Repository repository.SupplierRepository
}

func NewServiceSupplier(supplier repository.SupplierRepository) *ServiceSupplier {
	return &ServiceSupplier{Repository: supplier}
}

func (s *ServiceSupplier) Create(input *domain.SupplierInput) (*models.ModelSupplier, error) {
	var supplier domain.SupplierInput
	supplier.Name = input.Name
	supplier.Telepon = input.Telepon
	supplier.Email = input.Email
	supplier.Alamat = input.Alamat

	res, err := s.Repository.Create(&supplier)

	return res, err
}

func (s *ServiceSupplier) Results() (*[]models.ModelSupplier, error) {
	res, err := s.Repository.Results()

	return res, err
}

func (s *ServiceSupplier) Result(input *domain.SupplierInput) (*models.ModelSupplier, error) {
	var supplier domain.SupplierInput
	supplier.ID = input.ID

	res, err := s.Repository.Result(&supplier)

	return res, err
}

func (s *ServiceSupplier) Delete(input *domain.SupplierInput) (*models.ModelSupplier, error) {
	var supplier domain.SupplierInput
	supplier.ID = input.ID

	res, err := s.Repository.Delete(&supplier)

	return res, err
}

func (s *ServiceSupplier) Update(input *domain.SupplierInput) (*models.ModelSupplier, error) {
	var supplier domain.SupplierInput
	supplier.Name = input.Name
	supplier.Telepon = input.Telepon
	supplier.Email = input.Email
	supplier.Alamat = input.Alamat

	res, err := s.Repository.Update(&supplier)

	return res, err

}
