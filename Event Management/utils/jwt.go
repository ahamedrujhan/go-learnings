package utils

import (
	"errors"
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

func VerifyToken(token string) error {
	parseToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// check token signing method
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Unexpected signing method")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		errors.New("Colud not parse token")
	}

	isTokenValid := parseToken.Valid

	if !isTokenValid {
		return errors.New("Invalid token")
	}

	// jwt token claims type check

	_, ok := parseToken.Claims.(jwt.MapClaims)

	if !ok {
		return errors.New("Invalid token claims type")
	}

	// token claims check
	//email, ok := claims["email"].(string)
	//userId := claims["userId"].(int64)

	// token is verified
	return nil

}
