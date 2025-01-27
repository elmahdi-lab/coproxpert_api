package helpers_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	password := "password123"
	hashedPassword, err := HashPassword(password)

	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)
	assert.NotEqual(t, password, hashedPassword)
}

func TestHashPasswordWithInvalidCost(t *testing.T) {
	err := os.Setenv("PASSWORD_HASH_COST", "invalid")
	if err != nil {
		return
	}
	password := "password123"
	hashedPassword, err := HashPassword(password)

	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)
	assert.NotEqual(t, password, hashedPassword)
}

func TestIsPasswordHashValid(t *testing.T) {
	password := "password123"
	hashedPassword, _ := HashPassword(password)
	isValid := IsPasswordHashValid(password, hashedPassword)

	assert.True(t, isValid)
}

func TestIsPasswordHashInvalid(t *testing.T) {
	password := "password123"
	hashedPassword, _ := HashPassword(password)
	isValid := IsPasswordHashValid("wrongpassword", hashedPassword)

	assert.False(t, isValid)
}
