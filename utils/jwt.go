package utils

import (
	"github.com/dgrijalva/jwt-go"
)

// GenerateJWT creates a signed token
func GenerateJWT(secret string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	return token.SignedString([]byte(secret))
}
