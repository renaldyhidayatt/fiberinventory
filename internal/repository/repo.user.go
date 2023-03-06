package repository

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/models"

	"gorm.io/gorm"
)

type repositoryUser struct {
	db *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) *repositoryUser {
	return &repositoryUser{db: db}
}

func (r *repositoryUser) Register(input *domain.UserInput) (*models.ModelUser, error) {
	var user models.ModelUser
	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Email = input.Email
	user.Password = input.Password
	user.Role = input.Role

	db := r.db.Model(&user)

	checkEmailExist := db.Debug().Where("email = ?", input.Email)

	if checkEmailExist.RowsAffected > 0 {

		return &user, domain.ErrorRegisterEmail
	}

	addNewUser := db.Debug().Create(&user).Commit()
	if addNewUser.RowsAffected < 1 {
		return &user, domain.ErrorRegisterFailed
	}

	return &user, nil
}

func (r *repositoryUser) Login(input *domain.UserInput) (*models.ModelUser, error) {
	var user models.ModelUser
	user.Email = input.Email
	user.Password = input.Password

	db := r.db.Model(&user)

	checkEmailAndPassword := db.Debug().Where("email = ?", input.Email).Where("password = ?", input.Password)

	if checkEmailAndPassword.RowsAffected < 1 {

		return &user, domain.ErrorLoginFailed
	}

	return &user, nil
}

func (r *repositoryUser) Results() (*[]models.ModelUser, error) {

	var user []models.ModelUser

	db := r.db.Model(&user)

	checkResult := db.Debug().Find(&user)

	if checkResult.RowsAffected < 1 {
		return nil, domain.ErrorUsersNotFound
	}

	return &user, nil
}

func (r *repositoryUser) Result(input *domain.UserInput) (*models.ModelUser, error) {
	var user models.ModelUser

	user.ID = input.ID

	db := r.db.Model(&user)

	checkUser := db.Debug().First(&user)

	if checkUser.RowsAffected < 1 {
		return &user, domain.ErrorUserNotFound
	}

	return &user, nil
}

func (r *repositoryUser) Delete(input *domain.UserInput) (*models.ModelUser, error) {
	var user models.ModelUser

	user.ID = input.ID

	db := r.db.Model(&user)

	checkUser := db.Debug().First(&user)

	if checkUser.RowsAffected < 1 {
		return &user, domain.ErrorUserNotFound
	}

	return &user, nil
}

func (r *repositoryUser) FindByEmail(email string) (*models.ModelUser, error) {
	var user models.ModelUser

	db := r.db.Model(&user)

	checkEmail := db.Debug().Where("email = ?", email)

	if checkEmail.RowsAffected < 1 {
		return nil, domain.ErrorUserNotFound
	}

	return &user, nil
}

func (r *repositoryUser) Update(input *domain.UserInput) (*models.ModelUser, error) {
	var user models.ModelUser

	user.ID = input.ID

	db := r.db.Model(&user)

	checkUser := db.Debug().First(&user)

	if checkUser.RowsAffected < 1 {
		return &user, domain.ErrorUserNotFound
	}

	return &user, nil
}
