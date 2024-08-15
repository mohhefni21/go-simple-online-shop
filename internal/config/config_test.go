package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	t.Run("should return an error if file does not exits", func(t *testing.T) {
		// Arrange
		filename := "file.yaml"

		// Action & Assert
		require.NotNil(t, LoadConfig(filename), "should return an error when loading a valid config file")
	})

	t.Run("should load configuration successfully", func(t *testing.T) {
		// Arrange
		filename := "../../config.yaml"

		// Action & Assert
		require.Nil(t, LoadConfig(filename), "should not return an error when loading a valid config file")
	})
}
