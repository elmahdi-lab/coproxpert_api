package services

import (
	"errors"
	"github.com/google/uuid"
	"ithumans.com/coproxpert/helpers/security"
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/repositories"
	"log/slog"
	"time"
)

func CreateUser(u *models.User) (*models.User, error) {

	if u.Password == nil {
		return nil, errors.New("password is required")
	}

	if u.Username == nil {
		return nil, errors.New("username is required")
	}

	userRepository := repositories.NewUserRepository()
	hashedPassword, _ := security.HashPassword(*u.Password)
	u.Password = &hashedPassword
	err := userRepository.Create(u)
	if err != nil {
		slog.Error("error creating user", u)
		return nil, err
	}
	return u, nil
}

func GetUser(id uuid.UUID) (*models.User, error) {
	userRepository := repositories.NewUserRepository()
	user, err := userRepository.FindByID(id)
	if err != nil {
		slog.Error("error getting user", id)
		return nil, err
	}
	return user, nil
}

func UpdateUser(u *models.User) (*models.User, error) {
	userRepository := repositories.NewUserRepository()
	err := userRepository.Update(u)
	if err != nil {
		slog.Error("error updating user", u)
		return nil, err
	}
	return u, nil
}

func UpdatePassword(u *models.User) error {
	userRepository := repositories.NewUserRepository()
	user, err := userRepository.FindByID(u.ID)
	if err != nil {
		return err
	}

	hashedPassword, err := security.HashPassword(*u.Password)
	if err != nil {
		return err
	}
	user.Password = &hashedPassword
	if err := userRepository.Update(user); err != nil {
		return err
	}
	return nil
}

func DeleteUser(id uuid.UUID) error {
	userRepository := repositories.NewUserRepository()
	if deleted := userRepository.DeleteByID(id); deleted != true {
		slog.Error("error deleting user", id)
		return errors.New("user not found")
	}
	return nil
}

func PasswordForget(username string) error {
	userRepository := repositories.NewUserRepository()
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
	userRepository := repositories.NewUserRepository()
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
	userRepository := repositories.NewUserRepository()
	user, err := userRepository.FindByUsername(*u.Username)
	if err != nil {
		return nil, err
	}

	if user.IsLocked() {
		return nil, errors.New("user locked")
	}

	if !security.IsPasswordHashValid(*u.Password, *user.Password) {
		errMessage := "invalid credentials"
		user.IncrementTries()
		if *user.Tries >= 5 {
			user.Lock()
			errMessage = "user locked"
		}
		if err := userRepository.Update(user); err != nil {
			return nil, err
		}
		return nil, errors.New(errMessage)
	}

	user.Tries = new(int)
	user.Unlock()
	user.RefreshToken()
	if err := userRepository.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func Logout(u *models.User) error {
	if u == nil {
		return errors.New("user not logged in")
	}
	u.Token = nil
	u.TokenExpiresAt = nil
	userRepository := repositories.NewUserRepository()
	if err := userRepository.Update(u); err != nil {
		return err
	}
	return nil
}
