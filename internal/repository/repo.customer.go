package repository

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/models"

	"gorm.io/gorm"
)

type repositoryCustomer struct {
	db *gorm.DB
}

func NewRepositoryCustomer(db *gorm.DB) *repositoryCustomer {
	return &repositoryCustomer{db: db}
}

func (r *repositoryCustomer) Create(input *domain.CreateCustomerRequest) (*models.ModelCustomer, error) {
	var customer models.ModelCustomer

	customer.Name = input.Name
	customer.Telepon = input.Telepon
	customer.Email = input.Email
	customer.Alamat = input.Alamat

	db := r.db.Model(&customer)

	checkCustomerName := db.Debug().First(&customer, "name=?", input.Name)

	if checkCustomerName.RowsAffected > 0 {

		return &customer, domain.ErrorCustomerAlready
	}
	addCustomer := db.Debug().Create(&customer).Commit()

	if addCustomer.RowsAffected < 1 {
		return &customer, domain.ErrorCustomerCreateFailed
	}

	return &customer, nil

}

func (r *repositoryCustomer) Results() (*[]models.ModelCustomer, error) {
	var customer []models.ModelCustomer

	db := r.db.Model(&customer)

	checkCustomerName := db.Debug().Find(&customer)

	if checkCustomerName.RowsAffected < 1 {
		return nil, domain.ErrorCustomerNotFound
	}

	return &customer, nil
}

func (r *repositoryCustomer) Result(id string) (*models.ModelCustomer, error) {
	var customer models.ModelCustomer

	customer.ID = id

	db := r.db.Model(&customer)

	checkCustomerName := db.Debug().First(&customer)

	if checkCustomerName.RowsAffected < 1 {

		return &customer, domain.ErrorCustomerNotFound
	}

	return &customer, nil
}

func (r *repositoryCustomer) Delete(id string) (*models.ModelCustomer, error) {
	var customer models.ModelCustomer
	customer.ID = id

	db := r.db.Model(&customer)

	checkCustomerName := db.Debug().First(&customer)

	if checkCustomerName.RowsAffected < 1 {

		return &customer, domain.ErrorCustomerNotFound
	}

	deleteCustomer := db.Debug().Delete(&customer)

	if deleteCustomer.RowsAffected < 1 {

		return &customer, domain.ErrorCustomerDeleteFailed
	}
	return &customer, nil
}

func (r *repositoryCustomer) Update(input *domain.UpdateCustomerRequest) (*models.ModelCustomer, error) {
	var customer models.ModelCustomer

	customer.ID = input.ID
	db := r.db.Model(&customer)

	checkCustomerName := db.Debug().First(&customer)

	if checkCustomerName.RowsAffected < 1 {
		return &customer, domain.ErrorCustomerNotFound
	}

	customer.Name = input.Name
	customer.Alamat = input.Alamat
	customer.Email = input.Email
	customer.Telepon = input.Telepon

	updateCustomer := db.Debug().Updates(&customer)

	if updateCustomer.RowsAffected < 1 {
		return &customer, domain.ErrorCustomerUpdateFailed
	}

	return &customer, nil
}
