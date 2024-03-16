package services

import (
	"github.com/google/uuid"
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/repositories"
	"time"
)

func CreateUser(u *models.User) (*models.User, error) {
	userRepository, _ := repositories.NewUserRepository()
	u.ID = uuid.New()
	err := userRepository.Create(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func GetUser(id uuid.UUID) (*models.User, error) {
	userRepository, _ := repositories.NewUserRepository()
	user, err := userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByUsername(username string) (*models.User, error) {
	userRepository, _ := repositories.NewUserRepository()
	user, err := userRepository.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, nil

}

func GetUserByToken(token string) (*models.User, error) {
	userRepository, _ := repositories.NewUserRepository()
	user, err := userRepository.FindByToken(token)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(u *models.User) (*models.User, error) {
	userRepository, _ := repositories.NewUserRepository()

	err := userRepository.Update(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func CreatePasswordForgetToken(u *models.User) (*models.User, error) {
	userRepository, _ := repositories.NewUserRepository()

	token := uuid.New()
	resetTokenExpiresAt := u.CreatedAt.Add(models.PasswordResetTokenDurationMinutes * time.Minute)

	u.PasswordResetToken = &token
	u.ResetTokenExpiresAt = &resetTokenExpiresAt

	err := userRepository.Update(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func DeleteUser(id uuid.UUID) bool {
	userRepository, _ := repositories.NewUserRepository()
	return userRepository.DeleteByID(id)
}
