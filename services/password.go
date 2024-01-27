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

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), passwordHashCost)
	return string(bytes), err
}

func IsPasswordHashValid(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
