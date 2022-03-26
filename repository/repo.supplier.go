package repository

import (
	"fiberinventory/models"
	"fiberinventory/schemas"
	"net/http"

	"gorm.io/gorm"
)

type repositorySupplier struct {
	db *gorm.DB
}

func NewRepositorySupplier(db *gorm.DB) *repositorySupplier {
	return &repositorySupplier{db: db}
}

func (r *repositorySupplier) EntityCreate(input *schemas.SchemaSupplier) (*models.ModelSupplier, schemas.SchemaDatabaseError) {
	var supplier models.ModelSupplier

	supplier.Name = input.Name
	supplier.Telepon = input.Telepon
	supplier.Email = input.Email
	supplier.Alamat = input.Alamat

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&supplier)

	checkSupplierName := db.Debug().First(&supplier, "name=?", input.Name)

	if checkSupplierName.RowsAffected > 0 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_01",
		}
		return &supplier, <-err
	}
	addSupplier := db.Debug().Create(&supplier).Commit()

	if addSupplier.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_03",
		}
		return &supplier, <-err
	}

	err <- schemas.SchemaDatabaseError{}

	return &supplier, <-err

}

func (r *repositorySupplier) EntityResults() (*[]models.ModelSupplier, schemas.SchemaDatabaseError) {
	var supplier []models.ModelSupplier

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&supplier)

	checksupplierName := db.Debug().Find(&supplier)

	if checksupplierName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
	}

	err <- schemas.SchemaDatabaseError{}

	return &supplier, <-err
}

func (r *repositorySupplier) EntityResult(input *schemas.SchemaSupplier) (*models.ModelSupplier, schemas.SchemaDatabaseError) {
	var supplier models.ModelSupplier
	supplier.ID = input.ID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&supplier)

	checkSupplierName := db.Debug().First(&supplier)

	if checkSupplierName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &supplier, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &supplier, <-err
}

func (r *repositorySupplier) EntityDelete(input *schemas.SchemaSupplier) (*models.ModelSupplier, schemas.SchemaDatabaseError) {
	var supplier models.ModelSupplier
	supplier.ID = input.ID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&supplier)

	checksupplierName := db.Debug().First(&supplier)

	if checksupplierName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}

		return &supplier, <-err
	}

	deletesupplier := db.Debug().Delete(&supplier)

	if deletesupplier.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}

		return &supplier, <-err
	}
	err <- schemas.SchemaDatabaseError{}
	return &supplier, <-err
}

func (r *repositorySupplier) EntityUpdate(input *schemas.SchemaSupplier) (*models.ModelSupplier, schemas.SchemaDatabaseError) {
	var supplier models.ModelSupplier

	supplier.ID = input.ID
	err := make(chan schemas.SchemaDatabaseError, 1)
	db := r.db.Model(&supplier)

	checksupplierName := db.Debug().First(&supplier)

	if checksupplierName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &supplier, <-err
	}

	supplier.Name = input.Name
	supplier.Alamat = input.Alamat
	supplier.Email = input.Email
	supplier.Telepon = input.Telepon

	updateSupplier := db.Debug().Updates(&supplier)

	if updateSupplier.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &supplier, <-err
	}

	err <- schemas.SchemaDatabaseError{}

	return &supplier, <-err
}
