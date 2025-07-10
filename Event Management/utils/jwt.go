package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const secretKey = "xizccva2ss5r5"

func GenarateToken(userId int64, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"email":  email,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}
