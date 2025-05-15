package security

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	var hash, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func CompareHashPassword(pivotPassword, inputPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(pivotPassword), []byte(inputPassword))
	if err != nil {
		return fmt.Errorf("Password invalid")
	}
	return nil
}
