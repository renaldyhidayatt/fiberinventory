package service

import (
	"fiberinventory/entity"
	"fiberinventory/models"
	"fiberinventory/schemas"
)

type serviceCustomer struct {
	customer entity.EntityCustomer
}

func NewServiceCustomer(customer entity.EntityCustomer) *serviceCustomer {
	return &serviceCustomer{customer: customer}
}

func (s *serviceCustomer) EntityCreate(input *schemas.SchemaCustomer) (*models.ModelCustomer, schemas.SchemaDatabaseError) {
	var customer schemas.SchemaCustomer
	customer.Name = input.Name
	customer.Telepon = input.Telepon
	customer.Email = input.Email
	customer.Alamat = input.Alamat

	res, err := s.customer.EntityCreate(&customer)

	return res, err
}

func (s *serviceCustomer) EntityResults() (*[]models.ModelCustomer, schemas.SchemaDatabaseError) {
	res, err := s.customer.EntityResults()

	return res, err
}

func (s *serviceCustomer) EntityResult(input *schemas.SchemaCustomer) (*models.ModelCustomer, schemas.SchemaDatabaseError) {
	var customer schemas.SchemaCustomer
	customer.ID = input.ID

	res, err := s.customer.EntityResult(&customer)

	return res, err
}

func (s *serviceCustomer) EntityDelete(input *schemas.SchemaCustomer) (*models.ModelCustomer, schemas.SchemaDatabaseError) {
	var customer schemas.SchemaCustomer
	customer.ID = input.ID

	res, err := s.customer.EntityDelete(&customer)

	return res, err
}

func (s *serviceCustomer) EntityUpdate(input *schemas.SchemaCustomer) (*models.ModelCustomer, schemas.SchemaDatabaseError) {
	var customer schemas.SchemaCustomer
	customer.Name = input.Name
	customer.Telepon = input.Telepon
	customer.Email = input.Email
	customer.Alamat = input.Alamat

	res, err := s.customer.EntityUpdate(&customer)

	return res, err

}
