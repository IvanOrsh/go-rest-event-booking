package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// TODO: move to .env?
const secretKey = "secretKey"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(tokenString string) error {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return errors.New("could not parse token")
	}

	if !parsedToken.Valid {
		return errors.New("invalid token")
	}

	// extract data from token: email, userId
	// claims, ok := parsedToken.Claims.(jwt.MapClaims)
	// if !ok {
	// 	return errors.New("invalid token claims")
	// }

	// email := claims["email"].(string)
	// userId := claims["userId"].(float64)

	return nil
}
