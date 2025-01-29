package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"ithumans.com/coproxpert/data/services"
)

func UploadFileAction(c *fiber.Ctx) error {
	header, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "No file found",
		})
	}

	fileRecord, e := services.CreateFile(header)
	if e.Code != "" {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusCreated).JSON(fileRecord)
}

func GetFileAction(ctx *fiber.Ctx) error {

	fileId := ctx.Params("id")
	fileUuid := uuid.MustParse(fileId)
	fileRecord, err := services.GetFileByID(fileUuid)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return ctx.Status(fiber.StatusOK).JSON(fileRecord)
}

func UpdateFileAction(ctx *fiber.Ctx) error {

	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error:": "Not implemented"})
}

func DeleteFileAction(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error:": "Not implemented"})
}
