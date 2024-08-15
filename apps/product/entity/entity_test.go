package entity

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProductEntity(t *testing.T) {
	t.Run("should return an error when name product is not provided", func(t *testing.T) {
		// Arrange
		productEntity := ProductEntity{
			Name:  "",
			Stock: 10,
			Price: 25000,
		}

		// Action
		err := productEntity.Validate()

		// Assert
		require.NotNil(t, err)
	})

	t.Run("should return an error when name product have minimum character", func(t *testing.T) {
		// Arrange
		productEntity := ProductEntity{
			Name:  "no",
			Stock: 10,
			Price: 25000,
		}

		// Action
		err := productEntity.Validate()

		// Assert
		require.NotNil(t, err)
	})

	t.Run("should return an error when stock greater than 0", func(t *testing.T) {
		// Arrange
		productEntity := ProductEntity{
			Name:  "shampo",
			Stock: -1,
			Price: 25000,
		}

		// Action
		err := productEntity.Validate()

		// Assert
		require.NotNil(t, err)
	})

	t.Run("should return an error when price greater than 0", func(t *testing.T) {
		// Arrange
		productEntity := ProductEntity{
			Name:  "shampo",
			Stock: 10,
			Price: -25000,
		}

		// Action
		err := productEntity.Validate()

		// Assert
		require.NotNil(t, err)
	})

	t.Run("should not return an error when product payload valid", func(t *testing.T) {
		// Arrange
		productEntity := ProductEntity{
			Name:  "shampo",
			Stock: 10,
			Price: 25000,
		}

		// Action
		err := productEntity.Validate()

		// Assert
		require.Nil(t, err)
	})
}
