package repository

import (
	"fiberinventory/models"
	"fiberinventory/schemas"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type repositoryProduct struct {
	db *gorm.DB
}

func NewRepositoryProduct(db *gorm.DB) *repositoryProduct {
	return &repositoryProduct{db: db}
}

func (r *repositoryProduct) EntityCreate(input *schemas.SchemaProduct) (*models.ModelProduct, schemas.SchemaDatabaseError) {
	var product models.ModelProduct

	product.Name = input.Name
	product.Image = input.Image
	product.Qty = input.Qty
	product.CategoryID = input.CategoryID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&product)

	checkProductName := db.Debug().First(&product, "name=?", input.Name)

	if checkProductName.RowsAffected > 0 {
		err <- schemas.SchemaDatabaseError{
			Code: fiber.StatusConflict,
			Type: "error_create_01",
		}
		return &product, <-err
	}

	addProduct := db.Debug().Create(&product).Commit()

	if addProduct.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: fiber.StatusForbidden,
			Type: "error_create_03",
		}
		return &product, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &product, <-err
}

func (r *repositoryProduct) EntityResults() (*[]models.ModelProduct, schemas.SchemaDatabaseError) {
	var product []models.ModelProduct

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&product)

	checkProduct := db.Debug().Find(&product)

	if checkProduct.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: fiber.StatusNotFound,
			Type: "error_results_01",
		}
	}

	err <- schemas.SchemaDatabaseError{}
	return &product, <-err
}

func (r *repositoryProduct) EntityResult(input *schemas.SchemaProduct) (*models.ModelProduct, schemas.SchemaDatabaseError) {
	var product models.ModelProduct
	product.ID = input.ID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&product)

	checkProductName := db.Debug().First(&product)

	if checkProductName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: fiber.StatusNotFound,
			Type: "error_result_01",
		}
		return &product, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &product, <-err
}

func (r *repositoryProduct) EntityDelete(input *schemas.SchemaProduct) (*models.ModelProduct, schemas.SchemaDatabaseError) {
	var product models.ModelProduct
	product.ID = input.ID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&product)

	checkProduct := db.Debug().First(&product)

	if checkProduct.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: fiber.StatusForbidden,
			Type: "error_delete_01",
		}

		return &product, <-err
	}

	deleteProduct := db.Debug().Delete(&product)

	if deleteProduct.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: fiber.StatusForbidden,
			Type: "error_delete_02",
		}

		return &product, <-err
	}
	err <- schemas.SchemaDatabaseError{}
	return &product, <-err
}

func (r *repositoryProduct) EntityUpdate(input *schemas.SchemaProduct) (*models.ModelProduct, schemas.SchemaDatabaseError) {
	var product models.ModelProduct

	product.ID = input.ID
	err := make(chan schemas.SchemaDatabaseError, 1)
	db := r.db.Model(&product)

	checkProductName := db.Debug().First(&product)

	if checkProductName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: fiber.StatusNotFound,
			Type: "error_update_01",
		}
		return &product, <-err
	}

	product.Name = input.Name
	product.Image = input.Image
	product.Qty = input.Qty
	product.CategoryID = input.CategoryID

	updateProduct := db.Debug().Updates(&product)

	if updateProduct.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: fiber.StatusForbidden,
			Type: "error_update_02",
		}
		return &product, <-err
	}

	err <- schemas.SchemaDatabaseError{}

	return &product, <-err
}
