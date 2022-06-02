package entity

import (
	"fiberinventory/models"
	"fiberinventory/schemas"
)

type EntitySale interface {
	EntityCreate(input *schemas.SchemaSale) (*models.ModelSale, schemas.SchemaDatabaseError)
	EntityResult(input *schemas.SchemaSale) (*models.ModelSale, schemas.SchemaDatabaseError)
	EntityResults() (*[]models.ModelSale, schemas.SchemaDatabaseError)
	EntityDelete(input *schemas.SchemaSale) (*models.ModelSale, schemas.SchemaDatabaseError)
	EntityUpdate(input *schemas.SchemaSale) (*models.ModelSale, schemas.SchemaDatabaseError)
}
