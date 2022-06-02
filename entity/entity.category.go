package entity

import (
	"fiberinventory/models"
	"fiberinventory/schemas"
)

type EntityCategory interface {
	EntityCreate(input *schemas.SchemaCategory) (*models.ModelCategory, schemas.SchemaDatabaseError)
	EntityResult(input *schemas.SchemaCategory) (*models.ModelCategory, schemas.SchemaDatabaseError)
	EntityResults() (*[]models.ModelCategory, schemas.SchemaDatabaseError)
	EntityDelete(input *schemas.SchemaCategory) (*models.ModelCategory, schemas.SchemaDatabaseError)
	EntityUpdate(input *schemas.SchemaCategory) (*models.ModelCategory, schemas.SchemaDatabaseError)
}
