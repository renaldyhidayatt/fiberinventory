package repository

import (
	"fiberinventory/models"
	"fiberinventory/pkg"
	"fiberinventory/schemas"
	"net/http"

	"gorm.io/gorm"
)

type repositoryUser struct {
	db *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) *repositoryUser {
	return &repositoryUser{db: db}
}

func (r *repositoryUser) EntityRegister(input *schemas.SchemaUser) (*models.ModelUser, schemas.SchemaDatabaseError) {
	var user models.ModelUser
	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Email = input.Email
	user.Password = input.Password
	user.Role = input.Role

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&user)

	checkEmailExist := db.Debug().First(&user, "email = ?", input.Email)

	if checkEmailExist.RowsAffected > 0 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusConflict,
			Type: "error_register_01",
		}
		return &user, <-err
	}

	addNewUser := db.Debug().Create(&user).Commit()
	if addNewUser.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_register_02",
		}
		return &user, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &user, <-err
}

func (r *repositoryUser) EntityLogin(input *schemas.SchemaUser) (*models.ModelUser, schemas.SchemaDatabaseError) {
	var user models.ModelUser
	user.Email = input.Email
	user.Password = input.Password

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&user)

	checkEmailExist := db.Debug().First(&user, "email = ?", input.Email)

	if checkEmailExist.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_login_01",
		}
		return &user, <-err
	}

	checkPasswordMatch := pkg.ComparePassword(user.Password, input.Password)

	if checkPasswordMatch != nil {
		err <- schemas.SchemaDatabaseError{
			Code: http.StatusBadRequest,
			Type: "error_login_02",
		}
		return &user, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &user, <-err
}
