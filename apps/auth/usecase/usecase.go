package usecase

import (
	"context"
	"mohhefni/go-online-shop/apps/auth/entity"
	"mohhefni/go-online-shop/apps/auth/repository"
	"mohhefni/go-online-shop/apps/auth/request"
	"mohhefni/go-online-shop/infra/errorpkg"
	"mohhefni/go-online-shop/internal/config"
	"mohhefni/go-online-shop/utility"
)

type Usecase interface {
	RegisterUser(ctx context.Context, req request.RegisterRequestPayload) (id string, err error)
	LoginUser(ctx context.Context, req request.LoginRequestPayload) (token string, err error)
}

type usecase struct {
	repo repository.Repository
}

func NewUsecase(repo repository.Repository) Usecase {
	return &usecase{
		repo: repo,
	}
}

func (u *usecase) RegisterUser(ctx context.Context, req request.RegisterRequestPayload) (idUser string, err error) {
	authEntity := entity.NewFromRegisterRequest(req)
	if err = authEntity.RegisterValidate(); err != nil {
		return
	}

	err = u.repo.VerifyAvailableEmail(context.Background(), authEntity.Email)
	if err != nil {
		return
	}

	authEntity.Password, err = utility.EncryptPassword(authEntity.Password, uint8(config.Cfg.App.Encrytion.Salt))
	if err != nil {
		return
	}

	idUser, err = u.repo.AddUser(ctx, authEntity)

	return
}

func (u *usecase) LoginUser(ctx context.Context, req request.LoginRequestPayload) (token string, err error) {
	authEntity := entity.NewFromLoginRequest(req)
	if err = authEntity.LoginValidate(); err != nil {
		return
	}

	authEntity, err = u.repo.GetUserByEmail(context.Background(), req.Email)
	if err != nil {
		return
	}

	err = utility.VerifyPasswordFromPlain(authEntity.Password, req.Password)
	if err != nil {
		err = errorpkg.ErrPasswordNotMatch
		return
	}

	token, err = utility.GenerateToken(authEntity.PublicId.String(), string(authEntity.Role), config.Cfg.App.Encrytion.JWTSecret)

	return
}
