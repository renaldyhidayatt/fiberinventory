package repository

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/models"

	"gorm.io/gorm"
)

type repositorySale struct {
	db *gorm.DB
}

func NewRepositorySale(db *gorm.DB) *repositorySale {
	return &repositorySale{db: db}
}

func (r *repositorySale) Create(input *domain.SaleInput) (*models.ModelSale, error) {
	var sale models.ModelSale

	sale.Name = input.Name
	sale.Telepon = input.Telepon
	sale.Email = input.Email
	sale.Alamat = input.Alamat

	db := r.db.Model(&sale)

	checkSaleName := db.Debug().First(&sale, "name=?", input.Name)

	if checkSaleName.RowsAffected > 0 {

		return &sale, domain.ErrorSaleAlready
	}
	addSale := db.Debug().Create(&sale).Commit()

	if addSale.RowsAffected < 1 {
		return &sale, domain.ErrorSaleCreateFailed
	}

	return &sale, nil

}

func (r *repositorySale) Results() (*[]models.ModelSale, error) {
	var sale []models.ModelSale

	db := r.db.Model(&sale)

	checkCustomerName := db.Debug().Find(&sale)

	if checkCustomerName.RowsAffected < 1 {

		return nil, domain.ErrorCustomerNotFound
	}

	return &sale, nil
}

func (r *repositorySale) Result(input *domain.SaleInput) (*models.ModelSale, error) {
	var sale models.ModelSale
	sale.ID = input.ID

	db := r.db.Model(&sale)

	checkSaleName := db.Debug().First(&sale)

	if checkSaleName.RowsAffected < 1 {

		return &sale, domain.ErrorSaleNotFound
	}

	return &sale, nil
}

func (r *repositorySale) Delete(input *domain.SaleInput) (*models.ModelSale, error) {
	var sale models.ModelSale
	sale.ID = input.ID

	db := r.db.Model(&sale)

	checkSaleName := db.Debug().First(&sale)

	if checkSaleName.RowsAffected < 1 {

		return &sale, domain.ErrorSaleNotFound
	}

	deleteSale := db.Debug().Delete(&sale)

	if deleteSale.RowsAffected < 1 {

		return &sale, domain.ErrorSaleDeleteFailed
	}

	return &sale, nil
}

func (r *repositorySale) Update(input *domain.SaleInput) (*models.ModelSale, error) {
	var sale models.ModelSale

	sale.ID = input.ID
	db := r.db.Model(&sale)

	checkSaleName := db.Debug().First(&sale)

	if checkSaleName.RowsAffected < 1 {

		return &sale, domain.ErrorSaleNotFound
	}

	sale.Name = input.Name
	sale.Alamat = input.Alamat
	sale.Email = input.Email
	sale.Telepon = input.Telepon

	updateSale := db.Debug().Updates(&sale)

	if updateSale.RowsAffected < 1 {

		return &sale, domain.ErrorSaleUpdateFailed
	}

	return &sale, nil
}
