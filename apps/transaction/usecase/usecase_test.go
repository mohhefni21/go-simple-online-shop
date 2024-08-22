package usecase

import (
	"context"
	authrequest "mohhefni/go-online-shop/apps/auth/request"
	productrequest "mohhefni/go-online-shop/apps/product/request"
	"mohhefni/go-online-shop/apps/transaction/repository"
	transactionrequest "mohhefni/go-online-shop/apps/transaction/request"
	"mohhefni/go-online-shop/external/database"
	"mohhefni/go-online-shop/internal/config"
	"mohhefni/go-online-shop/test"
	"os"
	"testing"

	"github.com/google/uuid"
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

var authTableTestHelper *test.AuthTableTestHelper
var productTableTestHelper *test.ProductTableTestHelper
var transactionTableTestHelper *test.TransactionTableTestHelper

func TestMain(m *testing.M) {
	var err error
	authTableTestHelper, err = test.NewAuthTableTestHelper()
	if err != nil {
		panic(err)
	}
	productTableTestHelper, err = test.NewProductTableTestHelper()
	if err != nil {
		panic(err)
	}
	transactionTableTestHelper, err = test.NewTransactionTableTestHelper()
	if err != nil {
		panic(err)
	}

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestAddProduct(t *testing.T) {
	t.Cleanup(func() {
		if err := productTableTestHelper.CleanTableProduct(); err != nil {
			t.Fatal(err)
		}
		if err := authTableTestHelper.CleanTableUser(); err != nil {
			t.Fatal(err)
		}
		if err := transactionTableTestHelper.CleanTableTransaction(); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("should not return an error when payload valid", func(t *testing.T) {
		// Arrange
		// Adduser
		payloadUser := authrequest.RegisterRequestPayload{
			Email:    "user1@gmail.com",
			Password: "123456789",
		}
		authTableTestHelper, err := test.NewAuthTableTestHelper()
		if err != nil {
			panic(err)
		}
		publicId, err := authTableTestHelper.AddUser(payloadUser.Email, payloadUser.Password)
		if err != nil {
			panic(err)
		}
		// Addproduct
		sku := uuid.NewString()
		payloadProduct := productrequest.AddProductPayload{
			Name:  "sampo lifeboy",
			Stock: 21,
			Price: 25000,
		}
		if err := productTableTestHelper.AddProduct(sku, payloadProduct.Name, payloadProduct.Stock, payloadProduct.Price); err != nil {
			panic(err)
		}


		transaction := transactionrequest.AddTransactionPayload{
			ProduckSku:   sku,
			UserPublicId: publicId,
			Amount:       2,
		}

		// Action
		err = ucs.CreateTransaction(context.Background(), transaction)

		// Assert
		require.Nil(t, err)
	})
}
func TestGetTransactionHistory(t *testing.T) {
	t.Cleanup(func() {
		if err := authTableTestHelper.CleanTableUser(); err != nil {
			t.Fatal(err)
		}
		if err := transactionTableTestHelper.CleanTableTransaction(); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("should return transactions history by users", func(t *testing.T) {
		// Arrange
		// Adduser
		payloadUser := authrequest.RegisterRequestPayload{
			Email:    "user1@gmail.com",
			Password: "123456789",
		}
		authTableTestHelper, err := test.NewAuthTableTestHelper()
		if err != nil {
			panic(err)
		}
		publicId, err := authTableTestHelper.AddUser(payloadUser.Email, payloadUser.Password)
		if err != nil {
			panic(err)
		}

		// Action
		transactions, err := ucs.GetTransactionsHistory(context.Background(), publicId)

		// Assert
		require.Nil(t, err)
		t.Logf("%v", transactions)
	})
}
