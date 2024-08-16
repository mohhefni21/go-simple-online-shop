package usecase

import (
	"context"
	"mohhefni/go-online-shop/apps/transaction/repository"
)

type Usecase interface {
	CreateTransaction(ctx context.Context, email string, productSku string)(err, error)
}

type usecase struct {
	repository.ProductRepository
	repository.TransactionRepository
}

func NewTransaction(productRepo repository.ProductRepository, transactionRepo repository.TransactionRepository) Usecase {
	return &usecase{
		ProductRepository: productRepo,
		TransactionRepository: transactionRepo,
	}
}

func (u *usecase) CreateTransaction(ctx context.Context) {
	
}
