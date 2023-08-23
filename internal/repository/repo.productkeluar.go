package repository

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/models"

	"gorm.io/gorm"
)

type repositoryProductKeluar struct {
	db *gorm.DB
}

func NewRepositoryProductKeluar(db *gorm.DB) *repositoryProductKeluar {
	return &repositoryProductKeluar{db: db}
}

func (r *repositoryProductKeluar) Create(input *domain.CreateProductKeluarRequest) (*models.ModelProductKeluar, error) {
	var productkeluar models.ModelProductKeluar

	productkeluar.Qty = input.Qty
	productkeluar.ProductID = input.ProductID
	productkeluar.CategoryID = input.CategoryID

	db := r.db.Model(&productkeluar)

	addProductKeluar := db.Debug().Create(&productkeluar).Commit()

	if addProductKeluar.RowsAffected < 1 {

		return &productkeluar, domain.ErrorProductKeluarCreateFailed
	}

	return &productkeluar, nil
}

func (r *repositoryProductKeluar) Result(id string) (*models.ModelProductKeluar, error) {
	var productkeluar models.ModelProductKeluar

	db := r.db.Model(&productkeluar)

	checkProductKeluar := db.Debug().Where("id = ?", id).Find(&productkeluar)

	if checkProductKeluar.RowsAffected < 1 {
		return nil, domain.ErrorProductKeluarNotFound
	}

	return &productkeluar, nil
}

func (r *repositoryProductKeluar) Results() (*[]models.ModelProductKeluar, error) {
	var productkeluar []models.ModelProductKeluar

	db := r.db.Model(&productkeluar)

	checkProductKeluar := db.Debug().Find(&productkeluar)

	if checkProductKeluar.RowsAffected < 1 {
		return nil, domain.ErrorProductKeluarNotFound
	}

	return &productkeluar, nil
}

func (r *repositoryProductKeluar) Delete(id string) (*models.ModelProductKeluar, error) {
	var productkeluar models.ModelProductKeluar

	productkeluar.ID = id

	db := r.db.Model(&productkeluar)

	checkProductKeluar := db.Debug().First(&productkeluar)

	if checkProductKeluar.RowsAffected < 1 {

		return &productkeluar, domain.ErrorProductKeluarNotFound
	}

	return &productkeluar, nil
}

func (r *repositoryProductKeluar) Update(input *domain.UpdateProductKeluarRequest) (*models.ModelProductKeluar, error) {
	var productkeluar models.ModelProductKeluar

	productkeluar.ID = input.ID
	productkeluar.Qty = input.Qty
	productkeluar.ProductID = input.ProductID
	productkeluar.CategoryID = input.CategoryID

	db := r.db.Model(&productkeluar)

	checkProductKeluar := db.Debug().First(&productkeluar)

	if checkProductKeluar.RowsAffected < 1 {

		return &productkeluar, domain.ErrorProductKeluarNotFound
	}

	updateProductKeluar := db.Debug().Save(&productkeluar).Commit()

	if updateProductKeluar.RowsAffected < 1 {

		return &productkeluar, domain.ErrorProductKeluarUpdateFailed
	}

	return &productkeluar, nil
}
