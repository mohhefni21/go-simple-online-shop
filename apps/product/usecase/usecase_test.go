package usecase

import (
	"context"
	"fmt"
	"mohhefni/go-online-shop/apps/product/repository"
	"mohhefni/go-online-shop/apps/product/request"
	"mohhefni/go-online-shop/external/database"
	"mohhefni/go-online-shop/internal/config"
	"mohhefni/go-online-shop/test"
	"os"
	"strconv"
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

var productTableTestHelper *test.ProductTableTestHelper

func TestMain(m *testing.M) {
	var err error

	productTableTestHelper, err = test.NewProductTableTestHelper()
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
	})
	t.Run("should not return an error when payload valid", func(t *testing.T) {
		// Arrange
		payload := request.AddProductPayload{
			Name:  "shampo",
			Stock: 21,
			Price: 25000,
		}

		// Action
		_, err := ucs.CreateProduct(context.Background(), payload)

		// Assert
		require.Nil(t, err)
	})
}

func TestGetProduct(t *testing.T) {
	t.Cleanup(func() {
		if err := productTableTestHelper.CleanTableProduct(); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("should not return an error when payload valid", func(t *testing.T) {
		// Arrange
		// insert 4 products
		for i := 0; i < 4; i++ {
			sku := uuid.NewString()
			payload := request.AddProductPayload{
				Name:  fmt.Sprintf("sampo-%s", strconv.Itoa(i)),
				Stock: 21,
				Price: 25000,
			}
			if err := productTableTestHelper.AddProduct(sku, payload.Name, payload.Stock, payload.Price); err != nil {
				panic(err)
			}
		}
		paginationPayload := request.GetProductsRequestPayload{
			Cursor: 0,
			Size:   2,
		}

		// Action
		products, err := ucs.GetProducts(context.Background(), paginationPayload)

		// Assert
		require.Nil(t, err)
		require.NotNil(t, products)
		require.Equal(t, 2, len(products))
	})
}

func TestGetDetailProduct(t *testing.T) {
	t.Cleanup(func() {
		if err := productTableTestHelper.CleanTableProduct(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("should not return an error when get detail product", func(t *testing.T) {
		// Arrange
		sku := uuid.NewString()

		payload := request.AddProductPayload{
			Name:  "sampo lifeboy",
			Stock: 21,
			Price: 25000,
		}
		if err := productTableTestHelper.AddProduct(sku, payload.Name, payload.Stock, payload.Price); err != nil {
			panic(err)
		}

		// Action
		product, err := ucs.GetDetailProduct(context.Background(), sku)

		// Assert
		require.Nil(t, err)
		require.NotNil(t, product)
	})
}
