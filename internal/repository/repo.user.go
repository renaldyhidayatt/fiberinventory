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

func (r *repositoryUser) Register(input *domain.RegisterInput) (*models.ModelUser, error) {
	var user models.ModelUser
	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Email = input.Email
	user.Password = input.Password
	user.Role = input.Role

	db := r.db.Model(&user)

	_, err := r.FindByEmail(input.Email)

	if err != nil {
		return nil, domain.ErrorUserNotFound
	}

	addNewUser := db.Debug().Create(&user).Commit()
	if addNewUser.RowsAffected > 1 {
		return nil, domain.ErrorRegisterFailed
	}

	return &user, nil
}

func (r *repositoryUser) Login(input *domain.LoginInput) (*models.ModelUser, error) {
	var user models.ModelUser
	user.Email = input.Email
	user.Password = input.Password

	db := r.db.Model(&user)

	checkEmailExist := db.Debug().First(&user, "email=?", input.Email)

	if checkEmailExist.RowsAffected > 1 {
		return &user, nil
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

func (r *repositoryUser) Result(id string) (*models.ModelUser, error) {
	var user models.ModelUser

	user.ID = id

	db := r.db.Model(&user)

	checkUser := db.Debug().First(&user)

	if checkUser.RowsAffected < 1 {
		return &user, domain.ErrorUserNotFound
	}

	return &user, nil
}

func (r *repositoryUser) Delete(id string) (*models.ModelUser, error) {
	var user models.ModelUser

	user.ID = id

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

	checkEmailExist := db.Debug().First(&user, "email=?", email)

	if checkEmailExist.RowsAffected > 1 {
		return nil, checkEmailExist.Error
	}

	return &user, nil
}

func (r *repositoryUser) Update(input *domain.UpdateUserRequest) (*models.ModelUser, error) {
	var user models.ModelUser

	user.ID = input.ID

	db := r.db.Model(&user)

	checkUser := db.Debug().First(&user)

	if checkUser.RowsAffected < 1 {
		return &user, domain.ErrorUserNotFound
	}

	return &user, nil
}
