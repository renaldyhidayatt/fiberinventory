package service

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/models"
	"fiberinventory/internal/repository"
)

type serviceCustomer struct {
	Repository repository.CustomerRepository
}

func NewServiceCustomer(customer repository.CustomerRepository) *serviceCustomer {
	return &serviceCustomer{Repository: customer}
}

func (s *serviceCustomer) Create(input *domain.CustomerInput) (*models.ModelCustomer, error) {
	var customer domain.CustomerInput
	customer.Name = input.Name
	customer.Telepon = input.Telepon
	customer.Email = input.Email
	customer.Alamat = input.Alamat

	res, err := s.Repository.Create(&customer)

	return res, err
}

func (s *serviceCustomer) Results() (*[]models.ModelCustomer, error) {
	res, err := s.Repository.Results()

	return res, err
}

func (s *serviceCustomer) Result(input *domain.CustomerInput) (*models.ModelCustomer, error) {
	var customer domain.CustomerInput
	customer.ID = input.ID

	res, err := s.Repository.Result(&customer)

	return res, err
}

func (s *serviceCustomer) Delete(input *domain.CustomerInput) (*models.ModelCustomer, error) {
	var customer domain.CustomerInput
	customer.ID = input.ID

	res, err := s.Repository.Delete(&customer)

	return res, err
}

func (s *serviceCustomer) Update(input *domain.CustomerInput) (*models.ModelCustomer, error) {
	var customer domain.CustomerInput
	customer.Name = input.Name
	customer.Telepon = input.Telepon
	customer.Email = input.Email
	customer.Alamat = input.Alamat

	res, err := s.Repository.Update(&customer)

	return res, err

}
