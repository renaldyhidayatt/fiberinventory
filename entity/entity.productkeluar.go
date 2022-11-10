package entity

import (
	"fiberinventory/models"
	"fiberinventory/schemas"
)

type EntityProductKeluar interface {
	EntityCreate(input *schemas.SchemaProductKeluar) (*models.ModelProductKeluar, schemas.SchemaDatabaseError)
	EntityResult(input *schemas.SchemaProductKeluar) (*models.ModelProductKeluar, schemas.SchemaDatabaseError)
	EntityResults() (*[]models.ModelProductKeluar, schemas.SchemaDatabaseError)
	EntityDelete(input *schemas.SchemaProductKeluar) (*models.ModelProductKeluar, schemas.SchemaDatabaseError)
	EntityUpdate(input *schemas.SchemaProductKeluar) (*models.ModelProductKeluar, schemas.SchemaDatabaseError)
}
