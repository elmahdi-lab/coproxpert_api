package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"ithumans.com/coproxpert/helpers/security"
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/services"
)

func CreateUserAction(c *fiber.Ctx) error {

	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	hashedPassword, err := security.HashPassword(*user.Password)
	user.Password = &hashedPassword

	createdUser, err := services.CreateUser(user)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "An error occurred while creating the user"})
	}

	return c.JSON(fiber.Map{"message": "User created successfully", "user": createdUser})
}

func GetUserAction(c *fiber.Ctx) error {

	// loggedUser := c.Locals("user").(*models.User)

	id := c.Params("id")
	userUuid, err := uuid.Parse(id)
	user, err := services.GetUser(userUuid)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"user": user})
}

func UpdateUserAction(c *fiber.Ctx) error {

	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	updatedUser, err := services.UpdateUser(user)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "User updated successfully", "user": updatedUser})
}

func DeleteUserAction(c *fiber.Ctx) error {
	loggedUser := c.Locals("user").(*models.User)

	err := security.Guard(c, loggedUser, models.AdminRole, nil, nil)
	if err != nil {
		return err
	}

	id := c.Params("id")
	userUuid, err := uuid.Parse(id)
	deleted := services.DeleteUser(userUuid)

	if deleted != true {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "User deleted successfully"})
}

func LoginAction(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	existingUser, err := services.GetUserByUsername(*user.Username)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Invalid username or password"})
	}

	isPasswordValid := security.IsPasswordHashValid(*user.Password, *existingUser.Password)

	if isPasswordValid != true {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid username or password"})
	}

	existingUser.GenerateToken()

	updatedUser, err := services.UpdateUser(existingUser)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"user": fiber.Map{"username": *updatedUser.Username, "token": *updatedUser.Token}})
}

func LogoutAction(c *fiber.Ctx) error {
	// TODO: this does not work, because we don't have a concept of a session
	// TODO: probably logout should be on auth middleware so we can require the token
	loggedUser := c.Locals("user").(*models.User)

	loggedUser.Token = nil
	loggedUser.TokenExpiresAt = nil

	_, err := services.UpdateUser(loggedUser)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "User logged out successfully"})
}

func PasswordForgetAction(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := services.GetUserByUsername(*user.Username)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Invalid username"})
	}

	updatedUser, err := services.CreatePasswordForgetToken(user)
	// TODO: add func to make this more reusable
	// TODO: queue email sending

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Password reset token sent successfully", "user": updatedUser})
}

func PasswordResetAction(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	existingUser, err := services.GetUserByUsername(*user.Username)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Invalid username"})
	}

	if existingUser.PasswordResetToken == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid token"})
	}

	if *existingUser.PasswordResetToken != *user.PasswordResetToken {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid token"})
	}

	if existingUser.IsExpired() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Token expired"})
	}

	hashedPassword, err := security.HashPassword(*user.Password)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	existingUser.Password = &hashedPassword
	existingUser.PasswordResetToken = nil
	existingUser.ResetTokenExpiresAt = nil

	updatedUser, err := services.UpdateUser(existingUser)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Password reset successfully", "user": updatedUser})
}
