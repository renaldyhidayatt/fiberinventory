package repository

import (
	"fiberinventory/models"
	"fiberinventory/schemas"
	"net/http"

	"gorm.io/gorm"
)

type repositoryCategory struct {
	db *gorm.DB
}

func NewRepositoryCategory(db *gorm.DB) *repositoryCategory {
	return &repositoryCategory{db: db}
}

func (r *repositoryCategory) EntityCreate(input *schemas.SchemaCategory) (*models.ModelCategory, schemas.SchemaDatabaseError) {
	var category models.ModelCategory

	category.Name = input.Name

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&category)

	checkCategoryName := db.Debug().First(&category, "name=?", input.Name)
	if checkCategoryName.RowsAffected > 0 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusConflict,
			Type: "error_create_01",
		}
		return &category, <-err
	}
	addCategory := db.Debug().Create(&category).Commit()

	if addCategory.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_create_03",
		}
		return &category, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &category, <-err
}

func (r *repositoryCategory) EntityResults() (*[]models.ModelCategory, schemas.SchemaDatabaseError) {
	var category []models.ModelCategory

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&category)

	checkCategoryName := db.Debug().Find(&category)

	if checkCategoryName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_results_01",
		}

	}

	err <- schemas.SchemaDatabaseError{}
	return &category, <-err
}

func (r *repositoryCategory) EntityResult(input *schemas.SchemaCategory) (*models.ModelCategory, schemas.SchemaDatabaseError) {
	var category models.ModelCategory

	category.ID = input.ID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&category)

	checkCategoryName := db.Debug().First(&category)

	if checkCategoryName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_result_01",
		}
		return &category, <-err
	}
	err <- schemas.SchemaDatabaseError{}
	return &category, <-err

}

func (r *repositoryCategory) EntityDelete(input *schemas.SchemaCategory) (*models.ModelCategory, schemas.SchemaDatabaseError) {
	var category models.ModelCategory
	category.ID = input.ID

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&category)

	checkCategoryName := db.Debug().First(&category)

	if checkCategoryName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_delete_01",
		}
		return &category, <-err
	}

	deleteCategory := db.Debug().Delete(&category)

	if deleteCategory.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_delete_02",
		}
		return &category, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &category, <-err
}

func (r *repositoryCategory) EntityUpdate(input *schemas.SchemaCategory) (*models.ModelCategory, schemas.SchemaDatabaseError) {
	var category models.ModelCategory

	category.ID = input.ID
	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&category)

	checkCategoryName := db.Debug().First(&category)

	if checkCategoryName.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_update_01",
		}
		return &category, <-err
	}

	category.Name = input.Name

	updateCategory := db.Debug().Updates(&category)

	if updateCategory.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_update_02",
		}
		return &category, <-err
	}

	err <- schemas.SchemaDatabaseError{}

	return &category, <-err

}
