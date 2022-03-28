package repository

import (
	"fiberinventory/models"
	"fiberinventory/schemas"
	"net/http"

	"gorm.io/gorm"
)

type repositoryProductMasuk struct {
	db *gorm.DB
}

func NewRepositoryProductMasuk(db *gorm.DB) *repositoryProductMasuk {
	return &repositoryProductMasuk{db: db}
}

func (r *repositoryProductMasuk) EntityCreate(input *schemas.SchemaProductMasuk) (*models.ModelProductMasuk, schemas.SchemaDatabaseError) {
	var productmasuk models.ModelProductMasuk

	productmasuk.Name = input.Name
	productmasuk.Qty = input.Qty
	productmasuk.ProductID = input.ProductID
	productmasuk.SupplierID = input.SupplierID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&productmasuk)

	addProdutMasuk := db.Debug().Create(&productmasuk).Commit()

	if addProdutMasuk.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_03",
		}
		return &productmasuk, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &productmasuk, <-err
}

func (r *repositoryProductMasuk) EntityResults() (*[]models.ModelProductMasuk, schemas.SchemaDatabaseError) {
	var productmasuk []models.ModelProductMasuk

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&productmasuk)

	checkProductMasukName := db.Debug().Find(&productmasuk)

	if checkProductMasukName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
	}

	err <- schemas.SchemaDatabaseError{}

	return &productmasuk, <-err

}

func (r *repositoryProductMasuk) EntityResult(input *schemas.SchemaProductMasuk) (*models.ModelProductMasuk, schemas.SchemaDatabaseError) {
	var productmasuk models.ModelProductMasuk
	productmasuk.ID = input.ID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&productmasuk)

	checkProductMasukName := db.Debug().First(&productmasuk)

	if checkProductMasukName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &productmasuk, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &productmasuk, <-err
}

func (r *repositoryProductMasuk) EntityDelete(input *schemas.SchemaProductMasuk) (*models.ModelProductMasuk, schemas.SchemaDatabaseError) {
	var productmasuk models.ModelProductMasuk
	productmasuk.ID = input.ID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&productmasuk)

	checkProductMasukName := db.Debug().First(&productmasuk)

	if checkProductMasukName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}

		return &productmasuk, <-err
	}

	deleteProductMasuk := db.Debug().Delete(&productmasuk)

	if deleteProductMasuk.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}

		return &productmasuk, <-err
	}
	err <- schemas.SchemaDatabaseError{}
	return &productmasuk, <-err
}

func (r *repositoryProductMasuk) EntityUpdate(input *schemas.SchemaProductMasuk) (*models.ModelProductMasuk, schemas.SchemaDatabaseError) {
	var productmasuk models.ModelProductMasuk

	productmasuk.ID = input.ID
	err := make(chan schemas.SchemaDatabaseError, 1)
	db := r.db.Model(&productmasuk)

	checkProductMasukName := db.Debug().First(&productmasuk)

	if checkProductMasukName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &productmasuk, <-err
	}

	productmasuk.Name = input.Name
	productmasuk.Qty = input.Qty
	productmasuk.ProductID = input.ProductID
	productmasuk.SupplierID = input.SupplierID

	updateProductMasuk := db.Debug().Updates(&productmasuk)

	if updateProductMasuk.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &productmasuk, <-err
	}

	err <- schemas.SchemaDatabaseError{}

	return &productmasuk, <-err
}
