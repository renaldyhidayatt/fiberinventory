package service

import (
	"fiberinventory/entity"
	"fiberinventory/models"
	"fiberinventory/schemas"
)

type serviceSupplier struct {
	supplier entity.EntitySupplier
}

func NewServiceSupplier(supplier entity.EntitySupplier) *serviceSupplier {
	return &serviceSupplier{supplier: supplier}
}

func (s *serviceSupplier) EntityCreate(input *schemas.SchemaSupplier) (*models.ModelSupplier, schemas.SchemaDatabaseError) {
	var supplier schemas.SchemaSupplier
	supplier.Name = input.Name
	supplier.Telepon = input.Telepon
	supplier.Email = input.Email
	supplier.Alamat = input.Alamat

	res, err := s.supplier.EntityCreate(&supplier)

	return res, err
}

func (s *serviceSupplier) EntityResults() (*[]models.ModelSupplier, schemas.SchemaDatabaseError) {
	res, err := s.supplier.EntityResults()

	return res, err
}

func (s *serviceSupplier) EntityResult(input *schemas.SchemaSupplier) (*models.ModelSupplier, schemas.SchemaDatabaseError) {
	var supplier schemas.SchemaSupplier
	supplier.ID = input.ID

	res, err := s.supplier.EntityResult(&supplier)

	return res, err
}

func (s *serviceSupplier) EntityDelete(input *schemas.SchemaSupplier) (*models.ModelSupplier, schemas.SchemaDatabaseError) {
	var supplier schemas.SchemaSupplier
	supplier.ID = input.ID

	res, err := s.supplier.EntityDelete(&supplier)

	return res, err
}

func (s *serviceSupplier) EntityUpdate(input *schemas.SchemaSupplier) (*models.ModelSupplier, schemas.SchemaDatabaseError) {
	var supplier schemas.SchemaSupplier
	supplier.Name = input.Name
	supplier.Telepon = input.Telepon
	supplier.Email = input.Email
	supplier.Alamat = input.Alamat

	res, err := s.supplier.EntityUpdate(&supplier)

	return res, err

}
