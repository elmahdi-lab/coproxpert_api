package services

import (
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
)

func HashPassword(password string) (string, error) {
	passwordHashCostString := os.Getenv("PASSWORD_HASH_COST")
	passwordHashCost, err := strconv.Atoi(passwordHashCostString)
	if err != nil {
		passwordHashCost = bcrypt.DefaultCost
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), passwordHashCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func IsPasswordHashValid(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
