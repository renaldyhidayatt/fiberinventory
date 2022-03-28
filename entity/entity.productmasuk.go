package entity

import (
	"fiberinventory/models"
	"fiberinventory/schemas"
)

type EntityProductMasuk interface {
	EntityCreate(input *schemas.SchemaProductMasuk) (*models.ModelProductMasuk, schemas.SchemaDatabaseError)
	EntityResult(input *schemas.SchemaProductMasuk) (*models.ModelProductMasuk, schemas.SchemaDatabaseError)
	EntityResults() (*[]models.ModelProductMasuk, schemas.SchemaDatabaseError)
	EntityDelete(input *schemas.SchemaProductMasuk) (*models.ModelProductMasuk, schemas.SchemaDatabaseError)
	EntityUpdate(input *schemas.SchemaProductMasuk) (*models.ModelProductMasuk, schemas.SchemaDatabaseError)
}
