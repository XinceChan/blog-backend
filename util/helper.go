package util

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const SECRET_KEY = "SECRET_LOVE"

func GenerateJwt(issuer string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    issuer,
		ExpiresAt: time.Now().Add(time.Hour * time.Duration(24)).Unix(),
	})
	return claims.SignedString([]byte(SECRET_KEY))
}

func ParseJwt(cookie string) (string, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	if err != nil || !token.Valid {
		return "", err
	}
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return "", errors.New("invalid token")
	}
	return claims.Issuer, nil
}
