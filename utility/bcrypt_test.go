package utility

import (
	"mohhefni/go-online-shop/internal/config"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestEncryptPassword(t *testing.T) {
	t.Run("should encrypt password successfully and not be the same as plain password", func(t *testing.T) {
		// Arrange
		password := "plaintext"

		// Action
		encryptedPass, err := EncryptPassword(password, config.Cfg.App.Encrytion.Salt)

		// Assert
		require.Nil(t, err)
		require.NotEqual(t, password, encryptedPass)
	})

	t.Run("should successfully verify the encrypted password", func(t *testing.T) {
		// Arrange
		password := "plaintext"
		encryptedPass, _ := EncryptPassword(password, config.Cfg.App.Encrytion.Salt)

		// Action
		err := bcrypt.CompareHashAndPassword([]byte(encryptedPass), []byte(password))
		require.Nil(t, err)
	})
}
