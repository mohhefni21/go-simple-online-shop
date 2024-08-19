package usecase

import (
	"context"
	"mohhefni/go-online-shop/apps/transaction/repository"
	"mohhefni/go-online-shop/apps/transaction/request"
	"mohhefni/go-online-shop/external/database"
	"mohhefni/go-online-shop/internal/config"
	"testing"

	"github.com/stretchr/testify/require"
)

var ucs Usecase

func init() {
	filename := "../../../config.yaml"
	err := config.LoadConfig(filename)

	if err != nil {
		panic(err)
	}

	db, err := database.Connection(config.Cfg.Db)
	if err != nil {
		panic(err)
	}

	repository := repository.NewRepository(db)
	ucs = NewUsecase(repository)
}

func TestAddProduct(t *testing.T) {
	t.Run("should not return an error when payload valid", func(t *testing.T) {
		// Arrange
		transaction := request.AddTransactionPayload{
			ProduckSku:   "ef5ce3b1-c91e-44d6-945c-d44423c4e6de",
			UserPublicId: "9ad51038-87a5-4fee-b805-afbba2ee78df",
			Amount:       2,
		}

		err := ucs.CreateTransaction(context.Background(), transaction)

		require.Nil(t, err)
	})
}
