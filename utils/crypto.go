package utils

import "golang.org/x/crypto/bcrypt"

// GenerateHash creates bcrypt hash of a password
func GenerateHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}
