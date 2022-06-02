package service

import (
	"fiberinventory/entity"
	"fiberinventory/models"
	"fiberinventory/schemas"
)

type serviceSale struct {
	sale entity.EntitySale
}

func NewServiceSale(sale entity.EntitySale) *serviceSale {
	return &serviceSale{sale: sale}
}

func (s *serviceSale) EntityCreate(input *schemas.SchemaSale) (*models.ModelSale, schemas.SchemaDatabaseError) {
	var sale schemas.SchemaSale

	sale.Name = input.Name
	sale.Telepon = input.Telepon
	sale.Email = input.Email
	sale.Alamat = input.Alamat

	res, err := s.sale.EntityCreate(&sale)
	return res, err

}

func (s *serviceSale) EntityResults() (*[]models.ModelSale, schemas.SchemaDatabaseError) {
	res, err := s.sale.EntityResults()
	return res, err
}

func (s *serviceSale) EntityResult(input *schemas.SchemaSale) (*models.ModelSale, schemas.SchemaDatabaseError) {
	var sale schemas.SchemaSale

	sale.ID = input.ID

	res, err := s.sale.EntityResult(&sale)

	return res, err
}

func (s *serviceSale) EntityDelete(input *schemas.SchemaSale) (*models.ModelSale, schemas.SchemaDatabaseError) {
	var sale schemas.SchemaSale

	sale.ID = input.ID

	res, err := s.sale.EntityDelete(&sale)

	return res, err
}

func (s *serviceSale) EntityUpdate(input *schemas.SchemaSale) (*models.ModelSale, schemas.SchemaDatabaseError) {
	var sale schemas.SchemaSale
	sale.Name = input.Name
	sale.Telepon = input.Telepon
	sale.Email = input.Email
	sale.Alamat = input.Alamat

	res, err := s.sale.EntityUpdate(&sale)

	return res, err
}
