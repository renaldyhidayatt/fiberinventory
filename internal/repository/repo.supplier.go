package repository

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/models"

	"gorm.io/gorm"
)

type repositorySupplier struct {
	db *gorm.DB
}

func NewRepositorySupplier(db *gorm.DB) *repositorySupplier {
	return &repositorySupplier{db: db}
}

func (r *repositorySupplier) Create(input *domain.SupplierInput) (*models.ModelSupplier, error) {
	var supplier models.ModelSupplier

	supplier.Name = input.Name
	supplier.Telepon = input.Telepon
	supplier.Email = input.Email
	supplier.Alamat = input.Alamat

	db := r.db.Model(&supplier)

	checkSupplierName := db.Debug().First(&supplier, "name=?", input.Name)

	if checkSupplierName.RowsAffected > 0 {

		return &supplier, domain.ErrorSupplierAlready
	}
	addSupplier := db.Debug().Create(&supplier).Commit()

	if addSupplier.RowsAffected < 1 {
		return &supplier, domain.ErrorSupplierCreateFailed
	}

	return &supplier, nil

}

func (r *repositorySupplier) Results() (*[]models.ModelSupplier, error) {
	var supplier []models.ModelSupplier

	db := r.db.Model(&supplier)

	checksupplierName := db.Debug().Find(&supplier)

	if checksupplierName.RowsAffected < 1 {
		return nil, domain.ErrorSuppliersNotFound
	}

	return &supplier, nil
}

func (r *repositorySupplier) Result(input *domain.SupplierInput) (*models.ModelSupplier, error) {
	var supplier models.ModelSupplier
	supplier.ID = input.ID

	db := r.db.Model(&supplier)

	checkSupplierName := db.Debug().First(&supplier)

	if checkSupplierName.RowsAffected < 1 {

		return &supplier, domain.ErrorSupplierNotFound
	}

	return &supplier, nil
}

func (r *repositorySupplier) Delete(input *domain.SupplierInput) (*models.ModelSupplier, error) {
	var supplier models.ModelSupplier
	supplier.ID = input.ID

	db := r.db.Model(&supplier)

	checksupplierName := db.Debug().First(&supplier)

	if checksupplierName.RowsAffected < 1 {

		return &supplier, domain.ErrorSupplierNotFound
	}

	deletesupplier := db.Debug().Delete(&supplier)

	if deletesupplier.RowsAffected < 1 {

		return &supplier, domain.ErrorSupplierDeleteFailed
	}
	return &supplier, nil
}

func (r *repositorySupplier) Update(input *domain.SupplierInput) (*models.ModelSupplier, error) {
	var supplier models.ModelSupplier

	supplier.ID = input.ID
	db := r.db.Model(&supplier)

	checksupplierName := db.Debug().First(&supplier)

	if checksupplierName.RowsAffected < 1 {

		return &supplier, domain.ErrorSupplierNotFound
	}

	supplier.Name = input.Name
	supplier.Alamat = input.Alamat
	supplier.Email = input.Email
	supplier.Telepon = input.Telepon

	updateSupplier := db.Debug().Updates(&supplier)

	if updateSupplier.RowsAffected < 1 {

		return &supplier, domain.ErrorSupplierUpdateFailed
	}

	return &supplier, nil
}
