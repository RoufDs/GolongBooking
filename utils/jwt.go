package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func VerifyToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Unsupported signing method")
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return errors.New("Could not parse token")
	}

	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return errors.New("Token is not valid")
	}

	// claims, ok := parsedToken.Claims.(jwt.MapClaims)
	// if !ok {
	// 	return errors.New("Invalid token claims")
	// }

	// email := claims["email"].(string)
	// userId := claims["userId"].(int64)

	return nil
}
