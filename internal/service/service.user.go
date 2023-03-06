package service

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/models"
	"fiberinventory/internal/repository"
	"fiberinventory/pkg/auth"
	"fiberinventory/pkg/hash"
	"fmt"
)

type serviceUser struct {
	Repository repository.UserRepository
	Hash       hash.Hashing
	Token      auth.TokenManager
}

func NewServiceUser(auth repository.UserRepository, hash hash.Hashing, token auth.TokenManager) *serviceUser {
	return &serviceUser{Repository: auth, Hash: hash, Token: token}
}

func (s *serviceUser) Register(input *domain.UserInput) (*models.ModelUser, error) {
	var schema domain.UserInput
	schema.FirstName = input.FirstName
	schema.LastName = input.LastName
	schema.Email = input.Email
	schema.Password = input.Password
	schema.Role = input.Role

	res, err := s.Repository.Register(&schema)

	return res, err
}

func (s *serviceUser) Login(input *domain.UserInput) (domain.Token, error) {
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

func (s *serviceUser) createJwt(id string) (domain.Token, error) {
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

func (s *serviceUser) Results() (*[]models.ModelUser, error) {
	res, err := s.Repository.Results()

	return res, err
}

func (s *serviceUser) Result(input *domain.UserInput) (*models.ModelUser, error) {
	var user domain.UserInput

	user.ID = input.ID

	res, err := s.Repository.Result(&user)

	return res, err
}

func (s *serviceUser) Delete(input *domain.UserInput) (*models.ModelUser, error) {
	var user domain.UserInput

	user.ID = input.ID

	res, err := s.Repository.Delete(&user)

	return res, err
}

func (s *serviceUser) Update(input *domain.UserInput) (*models.ModelUser, error) {
	var user domain.UserInput

	user.ID = input.ID

	res, err := s.Repository.Update(&user)

	return res, err
}
