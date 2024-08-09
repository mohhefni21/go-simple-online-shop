package auth

import (
	"context"
	"mohhefni/go-online-shop/internal/config"

	"golang.org/x/crypto/bcrypt"
)

// Dependency Inversion
type Repository interface {
	AddUser(ctx context.Context, model AuthEntity) (err error)
	VerifyAvailableEmail(ctx context.Context, email string) (err error)
	GetUserByEmail(ctx context.Context, email string) (authEntity AuthEntity, err error)
}

type service struct {
	repo Repository
}

func newService(repo Repository) service {
	return service{
		repo: repo,
	}
}

func (s *service) RegisterUser(ctx context.Context, req RegisterRequestPayload) (err error) {
	authEntity := NewFromRegisterRequest(req)
	if err = authEntity.RegisterValidate(); err != nil {
		return
	}

	err = s.repo.VerifyAvailableEmail(context.Background(), authEntity.Email)
	if err != nil {
		return
	}

	authEntity.Password, err = s.EncryptPassword(authEntity.Password, uint8(config.Cfg.App.Encrytion.Salt))
	if err != nil {
		return
	}

	return s.repo.AddUser(ctx, authEntity)
}

func (s *service) EncryptPassword(pass string, salt uint8) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(passwordHash), nil
}

// func (s *service) LoginUser(ctx context.Context, req LoginRequestPayload) (token string, err error) {
// 	authEntity := NewFromLoginRequest(req)
// 	if err = authEntity.LoginValidate(); err != nil {
// 		return
// 	}

// 	authEntity, err = s.repo.GetUserByEmail(context.Background(), req.Email)
// 	if err != nil {
// 		return
// 	}

// 	err = s.VerifyPasswordFromPlain(authEntity.Password, req.Password)
// 	if err != nil {
// 		err = response.ErrPasswordNotMatch
// 		return
// 	}
// }

func (s *service) VerifyPasswordFromPlain(encrypted string, password string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(password))
	if err != nil {
		return
	}

	return
}
