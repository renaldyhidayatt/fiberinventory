package service

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/models"
	"fiberinventory/internal/repository"
	"fiberinventory/pkg/auth"
	"fiberinventory/pkg/hash"
	"fmt"
)

type ServiceUser struct {
	Repository repository.UserRepository
	Hash       hash.Hashing
	Token      auth.TokenManager
}

func NewServiceUser(auth repository.UserRepository, hash hash.Hashing, token auth.TokenManager) *ServiceUser {
	return &ServiceUser{Repository: auth, Hash: hash, Token: token}
}

func (s *ServiceUser) Register(input *domain.UserInput) (*models.ModelUser, error) {
	var schema domain.UserInput
	schema.FirstName = input.FirstName
	schema.LastName = input.LastName
	schema.Email = input.Email
	schema.Password = input.Password
	schema.Role = input.Role

	res, err := s.Repository.Register(&schema)

	return res, err
}

func (s *ServiceUser) Login(input *domain.UserInput) (domain.Token, error) {
	var schema domain.UserInput

	res, err := s.Repository.FindByEmail(input.Email)

	if err != nil {
		return domain.Token{}, fmt.Errorf("failed error %w", err)
	}

	err = s.Hash.ComparePassword(res.Password, input.Password)

	if err != nil {
		return domain.Token{}, fmt.Errorf("failed error :%w", err)
	}

	schema.Email = input.Email
	schema.Password = res.Password

	res, err = s.Repository.Login(&schema)

	if err != nil {
		return domain.Token{}, err
	}

	return s.createJwt(res.ID)
}

func (s *ServiceUser) createJwt(id string) (domain.Token, error) {
	var (
		res domain.Token
		err error
	)

	res.Jwt, err = s.Token.NewJWT(id)

	if err != nil {
		return res, err
	}

	return res, err
}

func (s *ServiceUser) Results() (*[]models.ModelUser, error) {
	res, err := s.Repository.Results()

	return res, err
}

func (s *ServiceUser) Result(input *domain.UserInput) (*models.ModelUser, error) {
	var user domain.UserInput

	user.ID = input.ID

	res, err := s.Repository.Result(&user)

	return res, err
}

func (s *ServiceUser) Delete(input *domain.UserInput) (*models.ModelUser, error) {
	var user domain.UserInput

	user.ID = input.ID

	res, err := s.Repository.Delete(&user)

	return res, err
}

func (s *ServiceUser) Update(input *domain.UserInput) (*models.ModelUser, error) {
	var user domain.UserInput

	user.ID = input.ID

	res, err := s.Repository.Update(&user)

	return res, err
}
