package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/helply/backend/app/models"
	"github.com/helply/backend/pkg/helpers"
	"github.com/helply/backend/platform/database"
)

func UploadFile(c *fiber.Ctx) error {
	claims, err := helpers.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Couldn't get user details.", "data": err})
	}
	file, err := c.FormFile("file")
	fileUUID := uuid.New()
	// Check for errors:
	if err == nil {
		err = c.SaveFile(file, fmt.Sprintf("./uploads/%s", fileUUID))
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Couldn't upload the file.", "data:": err})
		}
	}
	newFile := new(models.File)
	newFile.Name = file.Filename
	newFile.MimeType = file.Header.Get("Content-Type")
	newFile.Size = uint32(file.Size)
	newFile.IsPublic = true
	newFile.Path = fileUUID.String()
	newFile.OwnerID = claims.ID
	if err = database.Connection().Create(&newFile).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Couldn't create a file.", "data:": err})
	}
	return c.JSON(fiber.Map{"status:": "success", "message:": "File uploaded.", "data": newFile})
}
