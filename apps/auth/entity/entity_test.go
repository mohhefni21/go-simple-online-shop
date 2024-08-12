package entity

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEntityRegister(t *testing.T) {
	t.Run("should return an error when email is not provided", func(t *testing.T) {
		// Arrange
		authEntity := AuthEntity{
			Email:    "",
			Password: "supersecret",
		}

		// Action
		err := authEntity.RegisterValidate()

		// Assert
		require.NotNil(t, err)
	})

	t.Run("should return an error when email is invalid", func(t *testing.T) {
		// Arrange
		authEntity := AuthEntity{
			Email:    "admin",
			Password: "supersecret",
		}

		// Action
		err := authEntity.RegisterValidate()

		// Assert
		require.NotNil(t, err)
	})

	t.Run("should return an error when password is not provided", func(t *testing.T) {
		// Arrange
		authEntity := AuthEntity{
			Email:    "admin@gmail.com",
			Password: "",
		}

		// Action
		err := authEntity.RegisterValidate()

		// Assert
		require.NotNil(t, err)
	})

	t.Run("should return an error when password provided has less than minimum characters", func(t *testing.T) {
		// Arrange
		authEntity := AuthEntity{
			Email:    "admin@gmail.com",
			Password: "123",
		}

		// Action
		err := authEntity.RegisterValidate()

		// Assert
		require.NotNil(t, err)
	})

	t.Run("should not return an error for valid payload", func(t *testing.T) {
		// Arrange
		authEntity := AuthEntity{
			Email:    "admin@gmail.com",
			Password: "supersecret",
		}

		// Action
		err := authEntity.RegisterValidate()

		// Assert
		require.Nil(t, err)
	})
}
