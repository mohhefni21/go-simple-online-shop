package utility

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestEntityRegister(t *testing.T) {
	t.Run("should create token correctly", func(t *testing.T) {
		// Arrange
		id := uuid.NewString()

		// Action
		jwtToken, err := GenerateToken(id, "user", "ini-rahasia-sekali")

		// Assert
		require.NotNil(t, jwtToken)
		require.Nil(t, err)
		t.Log(jwtToken)
	})
}
