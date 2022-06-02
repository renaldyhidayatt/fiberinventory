package repository

import (
	"fiberinventory/models"
	"fiberinventory/schemas"
	"net/http"

	"gorm.io/gorm"
)

type repositoryCustomer struct {
	db *gorm.DB
}

func NewRepositoryCustomer(db *gorm.DB) *repositoryCustomer {
	return &repositoryCustomer{db: db}
}

func (r *repositoryCustomer) EntityCreate(input *schemas.SchemaCustomer) (*models.ModelCustomer, schemas.SchemaDatabaseError) {
	var customer models.ModelCustomer

	customer.Name = input.Name
	customer.Telepon = input.Telepon
	customer.Email = input.Email
	customer.Alamat = input.Alamat

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&customer)

	checkCustomerName := db.Debug().First(&customer, "name=?", input.Name)

	if checkCustomerName.RowsAffected > 0 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_01",
		}
		return &customer, <-err
	}
	addCustomer := db.Debug().Create(&customer).Commit()

	if addCustomer.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_03",
		}
		return &customer, <-err
	}

	err <- schemas.SchemaDatabaseError{}

	return &customer, <-err

}

func (r *repositoryCustomer) EntityResults() (*[]models.ModelCustomer, schemas.SchemaDatabaseError) {
	var customer []models.ModelCustomer

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&customer)

	checkCustomerName := db.Debug().Find(&customer)

	if checkCustomerName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}
	}

	err <- schemas.SchemaDatabaseError{}

	return &customer, <-err
}

func (r *repositoryCustomer) EntityResult(input *schemas.SchemaCustomer) (*models.ModelCustomer, schemas.SchemaDatabaseError) {
	var customer models.ModelCustomer
	customer.ID = input.ID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&customer)

	checkCustomerName := db.Debug().First(&customer)

	if checkCustomerName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &customer, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &customer, <-err
}

func (r *repositoryCustomer) EntityDelete(input *schemas.SchemaCustomer) (*models.ModelCustomer, schemas.SchemaDatabaseError) {
	var customer models.ModelCustomer
	customer.ID = input.ID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&customer)

	checkCustomerName := db.Debug().First(&customer)

	if checkCustomerName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}

		return &customer, <-err
	}

	deleteCustomer := db.Debug().Delete(&customer)

	if deleteCustomer.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}

		return &customer, <-err
	}
	err <- schemas.SchemaDatabaseError{}
	return &customer, <-err
}

func (r *repositoryCustomer) EntityUpdate(input *schemas.SchemaCustomer) (*models.ModelCustomer, schemas.SchemaDatabaseError) {
	var customer models.ModelCustomer

	customer.ID = input.ID
	err := make(chan schemas.SchemaDatabaseError, 1)
	db := r.db.Model(&customer)

	checkCustomerName := db.Debug().First(&customer)

	if checkCustomerName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &customer, <-err
	}

	customer.Name = input.Name
	customer.Alamat = input.Alamat
	customer.Email = input.Email
	customer.Telepon = input.Telepon

	updateCustomer := db.Debug().Updates(&customer)

	if updateCustomer.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &customer, <-err
	}

	err <- schemas.SchemaDatabaseError{}

	return &customer, <-err
}
