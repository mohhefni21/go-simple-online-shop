package utility

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id string, role string, secret string) (jwtToken string, err error) {
	claims := jwt.MapClaims{
		"public_id": id,
		"role":      role,
		"exp":       jwt.NewNumericDate(time.Now().Add(10 * time.Minute)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtToken, err = token.SignedString([]byte(secret))
	if err != nil {
		return
	}

	return
}

func ValidateToken(tokenString string, secret string) (publicId string, role string, err error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return []byte(secret), nil
	})

	if err != nil {
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		publicId = fmt.Sprintf("%v", claims["public_id"])
		role = fmt.Sprintf("%v", claims["role"])
		return
	}

	err = fmt.Errorf("unable to extract claims")

	return
}
