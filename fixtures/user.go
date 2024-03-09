package fixtures

import (
	"github.com/google/uuid"
	"ithumans.com/coproxpert/helpers"
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/repositories"
	"ithumans.com/coproxpert/services"
)

// CreateUser Create a new user with it's credential, permission, token and contact
func CreateUser() {

	user := &models.User{

		ID:         uuid.New(),
		Username:   helpers.StringPointer("elmahdi@example.com"),
		IsVerified: helpers.BoolPointer(true),
	}

	userRepository, _ := repositories.NewUserRepository()

	err := userRepository.Create(user)
	if err != nil {
		return
	}

	userContact := &models.Contact{
		ID:          uuid.New(),
		User:        user,
		IsDefault:   helpers.BoolPointer(true),
		PhoneNumber: helpers.StringPointer("555-555-5555"), // Use stringPtr for consistent pointers
		Address:     helpers.StringPointer("123 Fake Street"),
		City:        helpers.StringPointer("Anytown"),
		Province:    helpers.StringPointer("CA"),
		ZipCode:     helpers.StringPointer("12345"),
		Country:     helpers.StringPointer("US"),
	}

	contactRepository, _ := repositories.NewContactRepository()
	err = contactRepository.Create(userContact)

	token := &models.Token{
		ID:   uuid.New(),
		User: user,
	}
	token.GenerateToken()

	tokenRepository, _ := repositories.NewTokenRepository()
	err = tokenRepository.Create(token)

	hashedPassword, _ := services.HashPassword("password")

	credential := &models.Credential{
		ID:       uuid.New(),
		User:     user,
		Password: helpers.StringPointer(hashedPassword),
		Tries:    helpers.IntPointer(0),
	}

	credentialRepository, _ := repositories.NewCredentialRepository()
	err = credentialRepository.Create(credential)

	if userContact == nil {
		panic("userContact is nil")
	}

}
