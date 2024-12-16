package main

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func createToken(refreshToken string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"refreshToken": refreshToken,
		"exp":          time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(os.Getenv("SECRET_KEY"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			return os.Getenv("SECRET_KEY"), nil
		},
	)

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
