package repository

import (
	"fiberinventory/models"
	"fiberinventory/schemas"
	"net/http"

	"gorm.io/gorm"
)

type repositoryProductKeluar struct {
	db *gorm.DB
}

func NewRepositoryProductKeluar(db *gorm.DB) *repositoryProductKeluar {
	return &repositoryProductKeluar{db: db}
}

func (r *repositoryProductKeluar) EntityCreate(input *schemas.SchemaProductKeluar) (*models.ModelProductKeluar, schemas.SchemaDatabaseError) {
	var productkeluar models.ModelProductKeluar

	productkeluar.Qty = input.Qty
	productkeluar.ProductID = input.ProductID
	productkeluar.CategoryID = input.CategoryID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&productkeluar)

	addProductKeluar := db.Debug().Create(&productkeluar).Commit()

	if addProductKeluar.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_03",
		}
		return &productkeluar, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &productkeluar, <-err
}

func (r *repositoryProductKeluar) EntityResult(input *schemas.SchemaProductKeluar) (*models.ModelProductKeluar, schemas.SchemaDatabaseError) {
	var productkeluar models.ModelProductKeluar

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&productkeluar)

	checkProductKeluar := db.Debug().Where("id = ?", input.ID).Find(&productkeluar)

	if checkProductKeluar.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
	}

	err <- schemas.SchemaDatabaseError{}
	return &productkeluar, <-err
}

func (r *repositoryProductKeluar) EntityResults() (*[]models.ModelProductKeluar, schemas.SchemaDatabaseError) {
	var productkeluar []models.ModelProductKeluar

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&productkeluar)

	checkProductKeluar := db.Debug().Find(&productkeluar)

	if checkProductKeluar.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
	}

	err <- schemas.SchemaDatabaseError{}

	return &productkeluar, <-err
}

func (r *repositoryProductKeluar) EntityDelete(input *schemas.SchemaProductKeluar) (*models.ModelProductKeluar, schemas.SchemaDatabaseError) {
	var productkeluar models.ModelProductKeluar

	productkeluar.ID = input.ID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&productkeluar)

	checkProductKeluar := db.Debug().First(&productkeluar)

	if checkProductKeluar.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &productkeluar, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &productkeluar, <-err
}

func (r *repositoryProductKeluar) EntityUpdate(input *schemas.SchemaProductKeluar) (*models.ModelProductKeluar, schemas.SchemaDatabaseError) {
	var productkeluar models.ModelProductKeluar

	productkeluar.ID = input.ID
	productkeluar.Qty = input.Qty
	productkeluar.ProductID = input.ProductID
	productkeluar.CategoryID = input.CategoryID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&productkeluar)

	checkProductKeluar := db.Debug().First(&productkeluar)

	if checkProductKeluar.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &productkeluar, <-err
	}

	updateProductKeluar := db.Debug().Save(&productkeluar).Commit()

	if updateProductKeluar.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_01",
		}
		return &productkeluar, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &productkeluar, <-err
}
