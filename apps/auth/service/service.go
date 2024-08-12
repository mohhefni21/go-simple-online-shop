package service

import (
	"mohhefni/go-online-shop/apps/auth/repository"
	"mohhefni/go-online-shop/utility"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	VerifyPasswordFromPlain(encrypted string, password string) (err error)
	EncryptPassword(pass string, salt uint8) (string, error)
	GenerateToken(id string, role string, secret string) (token string, err error)
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) VerifyPasswordFromPlain(encrypted string, password string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(password))
	if err != nil {
		return
	}

	return
}

func (s *service) EncryptPassword(pass string, salt uint8) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(passwordHash), nil
}

func (s *service) GenerateToken(id string, role string, secret string) (token string, err error) {
	return utility.GenerateToken(id, role, secret)
}
