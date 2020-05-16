package security

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(hash), nil
}

func ComparePassword(hash string, password []byte) bool {
	byteHash := []byte(hash)
	err := bcrypt.CompareHashAndPassword(byteHash, password)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
