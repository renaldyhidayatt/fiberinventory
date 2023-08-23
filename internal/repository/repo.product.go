package repository

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/models"

	"gorm.io/gorm"
)

type repositoryProduct struct {
	db *gorm.DB
}

func NewRepositoryProduct(db *gorm.DB) *repositoryProduct {
	return &repositoryProduct{db: db}
}

func (r *repositoryProduct) Create(input *domain.CreateProductRequest) (*models.ModelProduct, error) {
	var product models.ModelProduct

	product.Name = input.Name
	product.Image = input.Image
	product.Qty = input.Qty
	product.CategoryID = input.CategoryID

	db := r.db.Model(&product)

	checkProductName := db.Debug().First(&product, "name=?", input.Name)

	if checkProductName.RowsAffected > 0 {

		return &product, domain.ErrorProductAlready
	}

	addProduct := db.Debug().Create(&product).Commit()

	if addProduct.RowsAffected < 1 {

		return &product, domain.ErrorProductCreateFailed
	}

	return &product, nil
}

func (r *repositoryProduct) Results() (*[]models.ModelProduct, error) {
	var product []models.ModelProduct

	db := r.db.Model(&product)

	checkProduct := db.Debug().Find(&product)

	if checkProduct.RowsAffected < 1 {
		return nil, domain.ErrorProductNotFound
	}

	return &product, nil
}

func (r *repositoryProduct) Result(id string) (*models.ModelProduct, error) {
	var product models.ModelProduct
	product.ID = id

	db := r.db.Model(&product)

	checkProductName := db.Debug().First(&product)

	if checkProductName.RowsAffected < 1 {

		return &product, domain.ErrorProductAlready
	}

	return &product, nil
}

func (r *repositoryProduct) Delete(id string) (*models.ModelProduct, error) {
	var product models.ModelProduct
	product.ID = id

	db := r.db.Model(&product)

	checkProduct := db.Debug().First(&product)

	if checkProduct.RowsAffected < 1 {

		return &product, domain.ErrorProductNotFound
	}

	deleteProduct := db.Debug().Delete(&product)

	if deleteProduct.RowsAffected < 1 {

		return &product, domain.ErrorProductDeleteFailed
	}
	return &product, nil
}

func (r *repositoryProduct) Update(input *domain.UpdateProductRequest) (*models.ModelProduct, error) {
	var product models.ModelProduct

	product.ID = input.ID
	db := r.db.Model(&product)

	checkProductName := db.Debug().First(&product)

	if checkProductName.RowsAffected < 1 {

		return &product, domain.ErrorProductNotFound
	}

	product.Name = input.Name
	product.Image = input.Image
	product.Qty = input.Qty
	product.CategoryID = input.CategoryID

	updateProduct := db.Debug().Updates(&product)

	if updateProduct.RowsAffected < 1 {
		return &product, domain.ErrorProductUpdateFailed
	}

	return &product, nil
}
