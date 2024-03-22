package services

import (
	"errors"
	"github.com/google/uuid"
	"ithumans.com/coproxpert/helpers/security"
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

func UpdateUser(u *models.User) (*models.User, error) {
	userRepository, _ := repositories.NewUserRepository()

	err := userRepository.Update(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func DeleteUser(id uuid.UUID) error {
	userRepository, _ := repositories.NewUserRepository()
	if deleted := userRepository.DeleteByID(id); deleted != true {
		return errors.New("user not found")
	}
	return nil
}

func PasswordForget(username string) error {
	userRepository, _ := repositories.NewUserRepository()
	user, err := userRepository.FindByUsername(username)
	if err != nil {
		return err
	}

	token := uuid.New()
	resetTokenExpiresAt := time.Now().Add(models.PasswordResetTokenDurationMinutes * time.Minute)

	user.PasswordResetToken = &token
	user.ResetTokenExpiresAt = &resetTokenExpiresAt

	if err := userRepository.Update(user); err != nil {
		return err
	}
	return nil
}

func PasswordReset(u *models.User) error {
	userRepository, _ := repositories.NewUserRepository()
	existingUser, err := userRepository.FindByUsername(*u.Username)
	if err != nil {
		return err
	}

	if existingUser.PasswordResetToken == nil || *existingUser.PasswordResetToken != *u.PasswordResetToken {
		return errors.New("invalid token")
	}

	if existingUser.IsTokenExpired() {
		return errors.New("token expired")
	}

	hashedPassword, err := security.HashPassword(*u.Password)
	if err != nil {
		return err
	}

	existingUser.Password = &hashedPassword
	existingUser.PasswordResetToken = nil
	existingUser.ResetTokenExpiresAt = nil

	if err := userRepository.Update(existingUser); err != nil {
		return err
	}
	return nil
}

func Login(u *models.User) (*models.User, error) {
	userRepository, _ := repositories.NewUserRepository()
	user, err := userRepository.FindByUsername(*u.Username)
	if err != nil {
		return nil, err
	}

	if user.IsLocked() {
		return nil, errors.New("user locked")
	}

	if !security.IsPasswordHashValid(*u.Password, *user.Password) {
		user.Tries = new(int)
		*user.Tries++
		if *user.Tries >= 5 {
			lockExpiresAt := time.Now().Add(models.LockDurationMinutes * time.Minute)
			user.LockExpiresAt = &lockExpiresAt
		}
		if err := userRepository.Update(user); err != nil {
			return nil, err
		}
		return nil, errors.New("invalid password")
	}

	user.Tries = new(int)
	*user.Tries = 0
	user.RefreshToken()
	if err := userRepository.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func Logout(u *models.User) error {
	userRepository, _ := repositories.NewUserRepository()
	u.Token = nil
	u.TokenExpiresAt = nil
	if err := userRepository.Update(u); err != nil {
		return err
	}
	return nil
}
