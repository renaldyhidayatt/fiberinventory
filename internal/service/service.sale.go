package service

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/models"
	"fiberinventory/internal/repository"
)

type ServiceSale struct {
	Repository repository.SaleRepository
}

func NewServiceSale(sale repository.SaleRepository) *ServiceSale {
	return &ServiceSale{Repository: sale}
}

func (s *ServiceSale) Create(input *domain.CreateSaleRequest) (*models.ModelSale, error) {
	var sale domain.CreateSaleRequest

	sale.Name = input.Name
	sale.Telepon = input.Telepon
	sale.Email = input.Email
	sale.Alamat = input.Alamat

	res, err := s.Repository.Create(&sale)
	return res, err

}

func (s *ServiceSale) Results() (*[]models.ModelSale, error) {
	res, err := s.Repository.Results()
	return res, err
}

func (s *ServiceSale) Result(id string) (*models.ModelSale, error) {

	res, err := s.Repository.Result(id)

	return res, err
}

func (s *ServiceSale) Delete(id string) (*models.ModelSale, error) {

	res, err := s.Repository.Delete(id)

	return res, err
}

func (s *ServiceSale) Update(input *domain.UpdateSaleRequest) (*models.ModelSale, error) {
	var sale domain.UpdateSaleRequest

	sale.Name = input.Name
	sale.Telepon = input.Telepon
	sale.Email = input.Email
	sale.Alamat = input.Alamat

	res, err := s.Repository.Update(&sale)

	return res, err
}
