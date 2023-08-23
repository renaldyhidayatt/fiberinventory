package service

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/models"
	"fiberinventory/internal/repository"
)

type ServiceCustomer struct {
	Repository repository.CustomerRepository
}

func NewServiceCustomer(customer repository.CustomerRepository) *ServiceCustomer {
	return &ServiceCustomer{Repository: customer}
}

func (s *ServiceCustomer) Create(input *domain.CreateCustomerRequest) (*models.ModelCustomer, error) {
	var customer domain.CreateCustomerRequest

	customer.Name = input.Name
	customer.Telepon = input.Telepon
	customer.Email = input.Email
	customer.Alamat = input.Alamat

	res, err := s.Repository.Create(&customer)

	return res, err
}

func (s *ServiceCustomer) Results() (*[]models.ModelCustomer, error) {
	res, err := s.Repository.Results()

	return res, err
}

func (s *ServiceCustomer) Result(id string) (*models.ModelCustomer, error) {

	res, err := s.Repository.Result(id)

	return res, err
}

func (s *ServiceCustomer) Delete(id string) (*models.ModelCustomer, error) {

	res, err := s.Repository.Delete(id)

	return res, err
}

func (s *ServiceCustomer) Update(input *domain.UpdateCustomerRequest) (*models.ModelCustomer, error) {
	var customer domain.UpdateCustomerRequest

	customer.Name = input.Name
	customer.Telepon = input.Telepon
	customer.Email = input.Email
	customer.Alamat = input.Alamat

	res, err := s.Repository.Update(&customer)

	return res, err

}
