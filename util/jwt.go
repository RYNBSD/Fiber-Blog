package util

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secret = "s"

func SignJwt(id string) (string, error) {
	add := time.Minute * 30;
	exp := time.Now().Add(add).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
		"exp": exp,
	})
	return token.SignedString([]byte(secret))
}

func VerifyJwt(authorization string) (string, error) {
	token, err := jwt.Parse(authorization, func (token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !token.Valid || !ok {
		return "", fmt.Errorf("Can't claims token")
	}

	exp := claims["exp"].(float64)
	expUnix := time.Unix(int64(exp), 0)
	if time.Now().After(expUnix) {
		return "", fmt.Errorf("Token has expired")
	}

	id,  ok := claims["id"].(string)
	if !ok {
		return "", fmt.Errorf("Invalid id type")
	}

	return id, nil
}