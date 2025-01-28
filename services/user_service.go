package services

import (
	"errors"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"ithumans.com/coproxpert/helpers"
	"ithumans.com/coproxpert/helpers/auth"
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/repositories"
)

func CreateUser(u *models.User) (*models.User, error) {

	if u.Password == nil {
		return nil, errors.New("password is required")
	}

	if u.Username == nil {
		return nil, errors.New("username is required")
	}

	userRepository := repositories.NewUserRepository()
	hashedPassword, _ := helpers.HashPassword(*u.Password)
	u.Password = &hashedPassword
	u.Tries = helpers.IntPointer(0)

	// Set refresh token expiration to 30 days from now
	expiresAt := time.Now().Add(30 * 24 * time.Hour)
	u.RefreshTokenExpiresAt = &expiresAt

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

	hashedPassword, err := helpers.HashPassword(*u.Password)
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
		return errors.New("invalid password token")
	}

	if existingUser.IsPasswordTokenExpired() {
		return errors.New("password token expired")
	}

	hashedPassword, err := helpers.HashPassword(*u.Password)
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

	if !helpers.IsPasswordHashValid(*u.Password, *user.Password) {
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

	user.Unlock()
	token, _ := auth.GenerateJWT(user.ID, user.CreatedAt)
	user.RefreshToken = auth.GenerateRefreshToken()
	user.Token = &token
	if err := userRepository.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func RefreshToken(refreshToken uuid.UUID) (*auth.RefreshTokenPayload, error) {
	var payload auth.RefreshTokenPayload
	userRepository := repositories.NewUserRepository()
	user, err := userRepository.FindByRefreshToken(refreshToken)
	if err != nil {
		return nil, errors.New("user not found for the refresh token")
	}

	if user.IsLocked() {
		return nil, errors.New("user is locked")
	}

	if user.IsRefreshTokenExpired() {
		return nil, errors.New("refresh token has expired")
	}

	token, err := auth.GenerateJWT(user.ID, user.CreatedAt)
	if err != nil {
		return nil, errors.New("failed to generate access token")
	}

	newRefreshToken := auth.GenerateRefreshToken()

	// Set new refresh token expiration to 30 days from now
	expiresAt := time.Now().Add(30 * 24 * time.Hour)
	user.RefreshTokenExpiresAt = &expiresAt

	user.RefreshToken = newRefreshToken
	user.Token = &token
	if err := userRepository.Update(user); err != nil {
		return nil, err
	}
	payload.RefreshToken = newRefreshToken.String()
	payload.Token = &token

	return &payload, nil
}

func Logout(u *models.User) error {
	if u == nil {
		return errors.New("user not logged in")
	}
	u.Token = nil
	userRepository := repositories.NewUserRepository()
	if err := userRepository.Update(u); err != nil {
		return err
	}
	return nil
}
