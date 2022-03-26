package entity

import (
	"fiberinventory/models"
	"fiberinventory/schemas"
)

type EntityCustomer interface {
	EntityCreate(input *schemas.SchemaCustomer) (*models.ModelCustomer, schemas.SchemaDatabaseError)
	EntityResult(input *schemas.SchemaCustomer) (*models.ModelCustomer, schemas.SchemaDatabaseError)
	EntityResults() (*[]models.ModelCustomer, schemas.SchemaDatabaseError)
	EntityDelete(input *schemas.SchemaCustomer) (*models.ModelCustomer, schemas.SchemaDatabaseError)
	EntityUpdate(input *schemas.SchemaCustomer) (*models.ModelCustomer, schemas.SchemaDatabaseError)
}
