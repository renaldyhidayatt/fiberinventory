package repository

import (
	"fiberinventory/models"
	"fiberinventory/schemas"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type repositorySale struct {
	db *gorm.DB
}

func NewRepositorySale(db *gorm.DB) *repositorySale {
	return &repositorySale{db: db}
}

func (r *repositorySale) EntityCreate(input *schemas.SchemaSale) (*models.ModelSale, schemas.SchemaDatabaseError) {
	var sale models.ModelSale

	sale.Name = input.Name
	sale.Telepon = input.Telepon
	sale.Email = input.Email
	sale.Alamat = input.Alamat

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&sale)

	checkSaleName := db.Debug().First(&sale, "name=?", input.Name)

	if checkSaleName.RowsAffected > 0 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_01",
		}
		return &sale, <-err
	}
	addSale := db.Debug().Create(&sale).Commit()

	if addSale.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_03",
		}
		return &sale, <-err
	}

	err <- schemas.SchemaDatabaseError{}

	return &sale, <-err

}

func (r *repositorySale) EntityResults() (*[]models.ModelSale, schemas.SchemaDatabaseError) {
	var sale []models.ModelSale

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&sale)

	checkCustomerName := db.Debug().Find(&sale)

	if checkCustomerName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
	}

	err <- schemas.SchemaDatabaseError{}

	return &sale, <-err
}

func (r *repositorySale) EntityResult(input *schemas.SchemaSale) (*models.ModelSale, schemas.SchemaDatabaseError) {
	var sale models.ModelSale
	sale.ID = input.ID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&sale)

	checkSaleName := db.Debug().First(&sale)

	if checkSaleName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &sale, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &sale, <-err
}

func (r *repositorySale) EntityDelete(input *schemas.SchemaSale) (*models.ModelSale, schemas.SchemaDatabaseError) {
	var sale models.ModelSale
	sale.ID = input.ID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&sale)

	checkSaleName := db.Debug().First(&sale)

	if checkSaleName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: fiber.StatusForbidden,
			Type: "error_delete_01",
		}

		return &sale, <-err
	}

	deleteSale := db.Debug().Delete(&sale)

	if deleteSale.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: fiber.StatusForbidden,
			Type: "error_delete_02",
		}

		return &sale, <-err
	}
	err <- schemas.SchemaDatabaseError{}
	return &sale, <-err
}

func (r *repositorySale) EntityUpdate(input *schemas.SchemaSale) (*models.ModelSale, schemas.SchemaDatabaseError) {
	var sale models.ModelSale

	sale.ID = input.ID
	err := make(chan schemas.SchemaDatabaseError, 1)
	db := r.db.Model(&sale)

	checkSaleName := db.Debug().First(&sale)

	if checkSaleName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: fiber.StatusNotFound,
			Type: "error_update_01",
		}
		return &sale, <-err
	}

	sale.Name = input.Name
	sale.Alamat = input.Alamat
	sale.Email = input.Email
	sale.Telepon = input.Telepon

	updateSale := db.Debug().Updates(&sale)

	if updateSale.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: fiber.StatusForbidden,
			Type: "error_update_02",
		}
		return &sale, <-err
	}

	err <- schemas.SchemaDatabaseError{}

	return &sale, <-err
}
