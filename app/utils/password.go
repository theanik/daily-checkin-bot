package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func MakeHash(str string) (string, error) {
	hashStr, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("could not hash password %w", err)
	}

	return string(hashStr), nil
}

func VerifyHash(hashStr string, rawStr string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashStr), []byte(rawStr))
}