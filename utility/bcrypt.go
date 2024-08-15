package utility

import "golang.org/x/crypto/bcrypt"

func VerifyPasswordFromPlain(encrypted string, password string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(password))
	if err != nil {
		return
	}

	return
}

func EncryptPassword(pass string, salt uint8) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(passwordHash), nil
}
