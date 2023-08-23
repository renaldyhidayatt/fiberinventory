package repository

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/models"

	"gorm.io/gorm"
)

type repositoryCategory struct {
	db *gorm.DB
}

func NewRepositoryCategory(db *gorm.DB) *repositoryCategory {
	return &repositoryCategory{db: db}
}

func (r *repositoryCategory) Create(input *domain.CreateCategoryRequest) (*models.ModelCategory, error) {
	var category models.ModelCategory

	category.Name = input.Name

	db := r.db.Model(&category)

	checkCategoryName := db.Debug().First(&category, "name=?", input.Name)
	if checkCategoryName.RowsAffected > 0 {

		return &category, domain.ErrorCategoryAlready
	}
	addCategory := db.Debug().Create(&category).Commit()

	if addCategory.RowsAffected < 1 {

		return &category, domain.ErrorCategoryCreateFailed
	}

	return &category, nil
}

func (r *repositoryCategory) Results() (*[]models.ModelCategory, error) {
	var category []models.ModelCategory

	db := r.db.Model(&category)

	checkCategoryName := db.Debug().Find(&category)

	if checkCategoryName.RowsAffected < 1 {
		return nil, domain.ErrorCategoryNotFound

	}

	return &category, nil
}

func (r *repositoryCategory) Result(id string) (*models.ModelCategory, error) {
	var category models.ModelCategory

	category.ID = id

	db := r.db.Model(&category)

	checkCategoryName := db.Debug().First(&category)

	if checkCategoryName.RowsAffected < 1 {
		return &category, domain.ErrorCategoryNotFound
	}
	return &category, nil

}

func (r *repositoryCategory) Delete(id string) (*models.ModelCategory, error) {
	var category models.ModelCategory
	category.ID = id

	db := r.db.Model(&category)

	checkCategoryName := db.Debug().First(&category)

	if checkCategoryName.RowsAffected < 1 {

		return &category, domain.ErrorCategoryNotFound
	}

	deleteCategory := db.Debug().Delete(&category)

	if deleteCategory.RowsAffected < 1 {
		return &category, domain.ErrorCategorDeleteFailed
	}

	return &category, nil
}

func (r *repositoryCategory) Update(input *domain.UpdateCategoryRequest) (*models.ModelCategory, error) {
	var category models.ModelCategory

	category.ID = input.ID

	db := r.db.Model(&category)

	checkCategoryName := db.Debug().First(&category)

	if checkCategoryName.RowsAffected < 1 {

		return &category, domain.ErrorCategoryNotFound
	}

	category.Name = input.Name

	updateCategory := db.Debug().Updates(&category)

	if updateCategory.RowsAffected < 1 {
		return &category, domain.ErrorCategoryUpdateFailed
	}

	return &category, nil

}
