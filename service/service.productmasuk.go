package service

import (
	"fiberinventory/entity"
	"fiberinventory/models"
	"fiberinventory/schemas"
)

type serviceProductMasuk struct {
	productmasuk entity.EntityProductMasuk
}

func NewServiceProductMasuk(productmasuk entity.EntityProductMasuk) *serviceProductMasuk {
	return &serviceProductMasuk{productmasuk: productmasuk}
}

func (s *serviceProductMasuk) EntityCreate(input *schemas.SchemaProductMasuk) (*models.ModelProductMasuk, schemas.SchemaDatabaseError) {
	var productmasuk schemas.SchemaProductMasuk

	productmasuk.Name = input.Name
	productmasuk.Qty = input.Qty
	productmasuk.ProductID = input.ProductID
	productmasuk.SupplierID = input.SupplierID

	res, err := s.productmasuk.EntityCreate(&productmasuk)

	return res, err
}

func (s *serviceProductMasuk) EntityResults() (*[]models.ModelProductMasuk, schemas.SchemaDatabaseError) {
	res, err := s.productmasuk.EntityResults()

	return res, err
}

func (s *serviceProductMasuk) EntityResult(input *schemas.SchemaProductMasuk) (*models.ModelProductMasuk, schemas.SchemaDatabaseError) {
	var productmasuk schemas.SchemaProductMasuk
	productmasuk.ID = input.ID

	res, err := s.productmasuk.EntityResult(&productmasuk)

	return res, err

}

func (s *serviceProductMasuk) EntityDelete(input *schemas.SchemaProductMasuk) (*models.ModelProductMasuk, schemas.SchemaDatabaseError) {
	var productmasuk schemas.SchemaProductMasuk
	productmasuk.ID = input.ID

	res, err := s.productmasuk.EntityDelete(&productmasuk)

	return res, err
}

func (s *serviceProductMasuk) EntityUpdate(input *schemas.SchemaProductMasuk) (*models.ModelProductMasuk, schemas.SchemaDatabaseError) {
	var productmasuk schemas.SchemaProductMasuk

	productmasuk.Name = input.Name
	productmasuk.Qty = input.Qty
	productmasuk.ProductID = input.ProductID
	productmasuk.SupplierID = input.SupplierID

	res, err := s.productmasuk.EntityUpdate(&productmasuk)

	return res, err

}
