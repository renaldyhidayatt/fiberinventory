package entity

import (
	"fiberinventory/models"
	"fiberinventory/schemas"
)

type EntityUser interface {
	EntityRegister(input *schemas.SchemaUser) (*models.ModelUser, schemas.SchemaDatabaseError)
	EntityLogin(input *schemas.SchemaUser) (*models.ModelUser, schemas.SchemaDatabaseError)
}
