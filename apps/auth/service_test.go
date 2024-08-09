package auth

import (
	"context"
	"mohhefni/go-online-shop/external/database"
	"mohhefni/go-online-shop/internal/config"
	"mohhefni/go-online-shop/test"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

var srv service

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
	srv = newService(repository)
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
			t.Fatalf("Failed to clean table: %v", err)
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
		err = srv.RegisterUser(context.Background(), payload)

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
		err := srv.RegisterUser(context.Background(), payload)

		// Assert
		require.Nil(t, err)
	})
}

func TestEncryptPassword(t *testing.T) {
	t.Run("should encrypt password successfully and not be the same as plain password", func(t *testing.T) {
		// Arrange
		password := "plaintext"

		// Action
		encryptedPass, err := srv.EncryptPassword(password, config.Cfg.App.Encrytion.Salt)

		// Assert
		require.Nil(t, err)
		require.NotEqual(t, password, encryptedPass)
	})

	t.Run("should successfully verify the encrypted password", func(t *testing.T) {
		// Arrange
		password := "plaintext"
		encryptedPass, _ := srv.EncryptPassword(password, config.Cfg.App.Encrytion.Salt)

		// Action
		err := bcrypt.CompareHashAndPassword([]byte(encryptedPass), []byte(password))
		require.Nil(t, err)
	})
}
