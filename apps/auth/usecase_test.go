package auth

import (
	"context"
	"mohhefni/go-online-shop/external/database"
	"mohhefni/go-online-shop/internal/config"
	"mohhefni/go-online-shop/test"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var ucs Usecase

func init() {
	filename := "../../cmd/api/config.yaml"
	err := config.LoadConfig(filename)

	if err != nil {
		panic(err)
	}

	db, err := database.Connection(config.Cfg.Db)
	if err != nil {
		panic(err)
	}

	repository := newRepository(db)
	svc := newService(repository)
	ucs = newUsecase(repository, svc)
}

var authTableTestHelper *test.AuthTableTestHelper

func TestMain(m *testing.M) {
	var err error
	authTableTestHelper, err = test.NewAuthTableTestHelper()
	if err != nil {
		panic(err)
	}

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestRegister(t *testing.T) {
	t.Cleanup(func() {
		if err := authTableTestHelper.CleanTable(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("should return an error when email already used", func(t *testing.T) {
		// Arrange
		payload := RegisterRequestPayload{
			Email:    "user1@gmail.com",
			Password: "123456789",
		}

		authTableTestHelper, err := test.NewAuthTableTestHelper()
		if err != nil {
			panic(err)
		}

		err = authTableTestHelper.AddUser(payload.Email, payload.Password)
		if err != nil {
			panic(err)
		}

		// Action
		_, err = ucs.RegisterUser(context.Background(), payload)

		// Assert
		require.NotNil(t, err)
	})

	t.Run("should not return an error when payload valid", func(t *testing.T) {
		// Arrange
		payload := RegisterRequestPayload{
			Email:    "user2@gmail.com",
			Password: "123456789",
		}

		// Action
		_, err := ucs.RegisterUser(context.Background(), payload)

		// Assert
		require.Nil(t, err)
	})
}

func TestLogin(t *testing.T) {
	t.Cleanup(func() {
		if err := authTableTestHelper.CleanTable(); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("should not return an error when credential valid", func(t *testing.T) {
		// Arrange
		email := "user2@gmail.com"
		password := "123456789"
		payloadRegister := RegisterRequestPayload{
			Email:    email,
			Password: password,
		}
		_, err := ucs.RegisterUser(context.Background(), payloadRegister)
		if err != nil {
			panic(err)
		}
		payloadLogin := LoginRequestPayload{
			Email:    email,
			Password: password,
		}

		// Action
		token, err := ucs.LoginUser(context.Background(), payloadLogin)

		// Assert
		require.Nil(t, err)
		require.NotEmpty(t, token)
		t.Log(token)
	})
}
