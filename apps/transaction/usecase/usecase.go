package usecase

import (
	"context"
	"mohhefni/go-online-shop/apps/transaction/entity"
	"mohhefni/go-online-shop/apps/transaction/repository"
	"mohhefni/go-online-shop/apps/transaction/request"
)

type Usecase interface {
	CreateTransaction(ctx context.Context, req request.AddTransactionPayload) (err error)
}

type usecase struct {
	repo repository.Repository
}

func NewUsecase(repo repository.Repository) Usecase {
	return &usecase{
		repo: repo,
	}
}

func (u *usecase) CreateTransaction(ctx context.Context, req request.AddTransactionPayload) (err error) {
	products, err := u.repo.GetDetailProductBySku(ctx, req.ProduckSku)
	if err != nil {
		return
	}

	trx := entity.NewTransactionFromRequest(req)
	trx.FromProductToTransaction(products)
	trx.SetPlatformFee(10000)

	err = u.repo.AddTransaction(ctx, *trx)

	return
}
