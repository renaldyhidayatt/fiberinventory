package repository

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/models"

	"gorm.io/gorm"
)

type repositoryProductMasuk struct {
	db *gorm.DB
}

func NewRepositoryProductMasuk(db *gorm.DB) *repositoryProductMasuk {
	return &repositoryProductMasuk{db: db}
}

func (r *repositoryProductMasuk) Create(input *domain.ProductMasukInput) (*models.ModelProductMasuk, error) {
	var productmasuk models.ModelProductMasuk

	productmasuk.Name = input.Name
	productmasuk.Qty = input.Qty
	productmasuk.ProductID = input.ProductID
	productmasuk.SupplierID = input.SupplierID

	db := r.db.Model(&productmasuk)

	addProdutMasuk := db.Debug().Create(&productmasuk).Commit()

	if addProdutMasuk.RowsAffected < 1 {

		return &productmasuk, domain.ErrorProductMasukCreateFailed
	}

	return &productmasuk, nil
}

func (r *repositoryProductMasuk) Results() (*[]models.ModelProductMasuk, error) {
	var productmasuk []models.ModelProductMasuk

	db := r.db.Model(&productmasuk)

	checkProductMasukName := db.Debug().Find(&productmasuk)

	if checkProductMasukName.RowsAffected < 1 {
		return nil, domain.ErrorProductMasukNotFound
	}

	return &productmasuk, nil

}

func (r *repositoryProductMasuk) Result(input *domain.ProductMasukInput) (*models.ModelProductMasuk, error) {
	var productmasuk models.ModelProductMasuk
	productmasuk.ID = input.ID

	db := r.db.Model(&productmasuk)

	checkProductMasukName := db.Debug().First(&productmasuk)

	if checkProductMasukName.RowsAffected < 1 {

		return &productmasuk, domain.ErrorProductMasukNotFound
	}

	return &productmasuk, nil
}

func (r *repositoryProductMasuk) Delete(input *domain.ProductMasukInput) (*models.ModelProductMasuk, error) {
	var productmasuk models.ModelProductMasuk
	productmasuk.ID = input.ID

	db := r.db.Model(&productmasuk)

	checkProductMasukName := db.Debug().First(&productmasuk)

	if checkProductMasukName.RowsAffected < 1 {

		return &productmasuk, domain.ErrorProductMasukNotFound
	}

	deleteProductMasuk := db.Debug().Delete(&productmasuk)

	if deleteProductMasuk.RowsAffected < 1 {

		return &productmasuk, domain.ErrorProductDeleteFailed
	}
	return &productmasuk, nil
}

func (r *repositoryProductMasuk) Update(input *domain.ProductMasukInput) (*models.ModelProductMasuk, error) {
	var productmasuk models.ModelProductMasuk

	productmasuk.ID = input.ID

	db := r.db.Model(&productmasuk)

	checkProductMasukName := db.Debug().First(&productmasuk)

	if checkProductMasukName.RowsAffected < 1 {

		return &productmasuk, domain.ErrorProductMasukNotFound
	}

	productmasuk.Name = input.Name
	productmasuk.Qty = input.Qty
	productmasuk.ProductID = input.ProductID
	productmasuk.SupplierID = input.SupplierID

	updateProductMasuk := db.Debug().Updates(&productmasuk)

	if updateProductMasuk.RowsAffected < 1 {

		return &productmasuk, domain.ErrorProductMasukUpdateFailed
	}

	return &productmasuk, nil
}
