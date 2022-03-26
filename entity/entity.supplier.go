package entity

import (
	"fiberinventory/models"
	"fiberinventory/schemas"
)

type EntitySupplier interface {
	EntityCreate(input *schemas.SchemaSupplier) (*models.ModelSupplier, schemas.SchemaDatabaseError)
	EntityResult(input *schemas.SchemaSupplier) (*models.ModelSupplier, schemas.SchemaDatabaseError)
	EntityResults() (*[]models.ModelSupplier, schemas.SchemaDatabaseError)
	EntityDelete(input *schemas.SchemaSupplier) (*models.ModelSupplier, schemas.SchemaDatabaseError)
	EntityUpdate(input *schemas.SchemaSupplier) (*models.ModelSupplier, schemas.SchemaDatabaseError)
}
