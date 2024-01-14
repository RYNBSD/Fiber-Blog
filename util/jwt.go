package util

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

const secret = "s"

func SignJwt(id string) (string, error) {
	add := time.Minute * 30
	exp := time.Now().Add(add).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": exp,
	})
	return token.SignedString([]byte(secret))
}

func VerifyJwt(key string) (string, error) {
	token, err := jwt.Parse(key, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !token.Valid || !ok {
		return "", fiber.NewError(fiber.StatusUnauthorized, "Can't claims token")
	}

	exp := claims["exp"].(float64)
	expUnix := time.Unix(int64(exp), 0)
	if time.Now().After(expUnix) {
		return "", fiber.NewError(fiber.StatusUnauthorized, "Token has expired")
	}

	id, ok := claims["id"].(string)
	if !ok {
		return "", fiber.NewError(fiber.StatusUnauthorized, "Invalid id type")
	}

	return id, nil
}
