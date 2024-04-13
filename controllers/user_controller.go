package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"ithumans.com/coproxpert/models"
	"ithumans.com/coproxpert/services"
)

func handleError(c *fiber.Ctx, err error, statusCode int) error {
	return c.Status(statusCode).JSON(fiber.Map{"error": err.Error()})
}

func CreateUserAction(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return handleError(c, err, fiber.StatusBadRequest)
	}

	createdUser, err := services.CreateUser(user)
	if err != nil {
		return handleError(c, err, fiber.StatusBadRequest)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"user": fiber.Map{"id": createdUser.ID, "username": *createdUser.Username}})
}

func GetUserAction(c *fiber.Ctx) error {
	id := c.Params("id")
	userUuid, err := uuid.Parse(id)

	if err != nil {
		return handleError(c, err, fiber.StatusBadRequest)
	}

	user, err := services.GetUser(userUuid)
	if err != nil {
		return handleError(c, err, fiber.StatusBadRequest)
	}
	user.Anonymize()

	return c.JSON(fiber.Map{"user": user})
}

func UpdateUserAction(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return handleError(c, err, fiber.StatusBadRequest)
	}

	updatedUser, err := services.UpdateUser(user)
	updatedUser.Anonymize()
	if err != nil {
		return handleError(c, err, fiber.StatusBadRequest)
	}
	return c.JSON(fiber.Map{"message": "User updated successfully", "user": updatedUser})
}

func DeleteUserAction(c *fiber.Ctx) error {
	id := c.Params("id")
	userUuid, err := uuid.Parse(id)

	if err != nil {
		return handleError(c, err, fiber.StatusBadRequest)
	}

	if deleted := services.DeleteUser(userUuid); deleted != nil {
		return handleError(c, err, fiber.StatusBadRequest)
	}
	return c.JSON(fiber.Map{"message": "User deleted successfully"})
}

func UpdatePasswordAction(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return handleError(c, err, fiber.StatusBadRequest)
	}

	if err := services.UpdatePassword(user); err != nil {
		return handleError(c, err, fiber.StatusBadRequest)
	}
	return c.JSON(fiber.Map{"message": "Password updated successfully"})

}

func PasswordForgetAction(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return handleError(c, err, fiber.StatusBadRequest)
	}

	if err := services.PasswordForget(*user.Username); err != nil {
		return handleError(c, err, fiber.StatusBadRequest)
	}
	return c.JSON(fiber.Map{"message": "Password reset token sent successfully"})
}

func PasswordResetAction(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return handleError(c, err, fiber.StatusBadRequest)
	}

	if err := services.PasswordReset(user); err != nil {
		return handleError(c, err, fiber.StatusBadRequest)
	}
	return c.JSON(fiber.Map{"message": "Password reset successfully"})
}

func LoginAction(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return handleError(c, err, fiber.StatusBadRequest)
	}

	loggedUser, err := services.Login(user)
	if err != nil {
		return handleError(c, err, fiber.StatusBadRequest)
	}
	loggedUser.Anonymize()
	return c.JSON(fiber.Map{"user": loggedUser})
}

func LogoutAction(c *fiber.Ctx) error {
	loggedUser := c.Locals("user").(*models.User)
	if err := services.Logout(loggedUser); err != nil {
		return handleError(c, err, fiber.StatusBadRequest)
	}
	return c.JSON(fiber.Map{"message": "Logged out successfully"})
}
