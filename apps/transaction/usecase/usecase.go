package usecase

import (
	"context"
	"fmt"
	"mohhefni/go-online-shop/apps/transaction/entity"
	"mohhefni/go-online-shop/apps/transaction/repository"
	"mohhefni/go-online-shop/apps/transaction/request"
)

type Usecase interface {
	CreateTransaction(ctx context.Context, req request.AddTransactionPayload) (err error)
	GetTransactionsHistory(ctx context.Context, publicId string) (transactions []entity.TransactionEntity, err error)
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
		fmt.Println("not-found")
		return
	}

	trx := entity.NewTransactionFromRequest(req).
		FromProductToTransaction(products).
		SetPlatformFee(10000).
		SetGrandTotal()

	err = trx.ValidateAmount()
	if err != nil {
		return
	}

	err = trx.ValidateStock(products.Stock)
	if err != nil {
		return
	}

	tx, err := u.repo.Begin(ctx)
	if err != nil {
		return err
	}

	defer u.repo.Roolback(ctx, tx)

	err = u.repo.AddTransaction(ctx, tx, *trx)
	if err != nil {
		return err
	}

	err = products.UpdateStockProduct(trx.Amount)
	if err != nil {
		return err
	}

	err = u.repo.UpdateStockProduct(ctx, tx, products)

	err = u.repo.Commit(ctx, tx)
	if err != nil {
		return err
	}

	return
}

func (u *usecase) GetTransactionsHistory(ctx context.Context, publicId string) (transactions []entity.TransactionEntity, err error) {
	transactions, err = u.repo.GetTransactionByUser(ctx, publicId)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
