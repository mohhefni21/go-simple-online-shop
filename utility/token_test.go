package utility

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestGenerateToken(t *testing.T) {
	t.Run("should generate the token correctly", func(t *testing.T) {
		// Arrange
		id := uuid.NewString()

		// Action
		jwtToken, err := GenerateToken(id, "user", "very-secret")

		// Assert
		require.NotNil(t, jwtToken)
		require.Nil(t, err)
	})
}

func TestValidateToken(t *testing.T) {
	t.Run("should verify the token correctly", func(t *testing.T) {
		// Arrange
		idPayload := uuid.NewString()
		rolePayload := "user"
		secret := "very-secret"
		jwtToken, _ := GenerateToken(idPayload, rolePayload, secret)

		// Action
		id, role, err := ValidateToken(jwtToken, secret)

		// Assert
		require.Nil(t, err)
		require.Equal(t, idPayload, id)
		require.Equal(t, rolePayload, role)

	})
}
