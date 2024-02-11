package hash

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type Hashing struct{}

func NewHashingPassword() *Hashing {
	return &Hashing{}
}

func (h Hashing) HashPassword(password string) (string, error) {
	pw := []byte(password)
	hashedPw, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPw), nil
}

func (h Hashing) ComparePassword(hashPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}
