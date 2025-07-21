package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashedPassword(plaintextPassword string) (string, error) {
	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), bcrypt.DefaultCost)
	if err != nil{
		return "", fmt.Errorf("unable to hash the password: %w", err);
	}
	return string(HashedPassword), nil
}