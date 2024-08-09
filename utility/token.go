package utility

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id string, role string, secret string) (jwtToken string, err error) {
	claims := jwt.MapClaims{
		"id":   id,
		"role": role,
		"exp":  jwt.NewNumericDate(time.Now().Add(10 * time.Minute)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtToken, err = token.SignedString([]byte(secret))
	if err != nil {
		return
	}

	return
}

// func ValidateToken(tokenString string, secret string) (id string, role string, err error) {
// 	tokens, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
// 		if _, ok := t.Method.(*jwt.SigningMethodECDSA); !ok {
// 			return nil, fmt.Errorf("unexpected signing method")
// 		}

// 		return secret, nil
// 	})

// 	if err != nil {
// 		return
// 	}

// }
