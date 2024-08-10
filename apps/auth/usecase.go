package auth

import (
	"context"
	"mohhefni/go-online-shop/infra/response"
	"mohhefni/go-online-shop/internal/config"
)

type Usecase interface {
	RegisterUser(ctx context.Context, req RegisterRequestPayload) (err error)
	LoginUser(ctx context.Context, req LoginRequestPayload) (token string, err error)
}

type usecase struct {
	repo Repository
	svc  Service
}

func newUsecase(repo Repository, service Service) Usecase {
	return &usecase{
		repo: repo,
		svc:  service,
	}
}

func (u *usecase) RegisterUser(ctx context.Context, req RegisterRequestPayload) (err error) {
	authEntity := NewFromRegisterRequest(req)
	if err = authEntity.RegisterValidate(); err != nil {
		return
	}

	err = u.repo.VerifyAvailableEmail(context.Background(), authEntity.Email)
	if err != nil {
		return
	}

	authEntity.Password, err = u.svc.EncryptPassword(authEntity.Password, uint8(config.Cfg.App.Encrytion.Salt))
	if err != nil {
		return
	}

	return u.repo.AddUser(ctx, authEntity)
}

func (u *usecase) LoginUser(ctx context.Context, req LoginRequestPayload) (token string, err error) {
	authEntity := NewFromLoginRequest(req)
	if err = authEntity.LoginValidate(); err != nil {
		return
	}

	authEntity, err = u.repo.GetUserByEmail(context.Background(), req.Email)
	if err != nil {
		return
	}

	err = u.svc.VerifyPasswordFromPlain(authEntity.Password, req.Password)
	if err != nil {
		err = response.ErrPasswordNotMatch
		return
	}

	token, err = u.svc.GenerateToken(authEntity.PublicId.String(), string(authEntity.Role), config.Cfg.App.Encrytion.JWTSecret)

	return
}
