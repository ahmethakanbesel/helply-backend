package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/helply/backend/app/dto"
	"github.com/helply/backend/app/models"
	"github.com/helply/backend/platform/database"
)

// CreateLicense
// @Summary Create a new licence.
// @Tags License
// @Accept json
// @Produce json
// @Success 200 {object} models.License
// @Security ApiKeyAuth
// @Router /api/v1/licences [post]
func CreateLicense(ctx *fiber.Ctx) error {
	db := database.Connection()
	newLicense := new(dto.LicenseDTO)
	if err := ctx.BodyParser(newLicense); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"status:": "error", "message:": "Invalid data given.", "data:": err})
	}
	license := new(models.License)
	license.ExpiresAt = newLicense.ExpiresAt
	license.Code = newLicense.Code
	license.ProductID = newLicense.ProductID
	license.IsActive = newLicense.IsActive
	if err := db.Create(&license).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Couldn't create a license.", "data:": err})
	}

	return ctx.JSON(fiber.Map{"status:": "success", "message:": "License created.", "data:": license})

}

func DeleteLicense(ctx *fiber.Ctx) error {
	license := &models.License{}
	err := database.Connection().Delete(&license, "id = ?", ctx.Params("id")).Error
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Couldn't delete the license.", "data:": err})
	}

	return ctx.JSON(fiber.Map{"status": "success", "message": "License deleted."})
}

func GetLicenses(ctx *fiber.Ctx) error {
	var licenses []models.License
	database.Connection().Joins("Product").Preload("Product.Image").Find(&licenses)
	return ctx.JSON(fiber.Map{"status": "success", "message": "Licenses listed.", "data": licenses})
}

func GetLicense(ctx *fiber.Ctx) error {
	license := &models.License{}
	err := database.Connection().First(&license, "id = ?", ctx.Params("id")).Error
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Couldn't get the license.", "data:": err})
	}
	return ctx.JSON(fiber.Map{"status": "success", "message": "", "data": license})
}

func UpdateLicense(ctx *fiber.Ctx) error {
	db := database.Connection()
	newLicense := new(dto.LicenseDTO)
	if err := ctx.BodyParser(newLicense); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"status:": "error", "message:": "Invalid data given.", "data:": err})
	}
	license := new(models.License)
	license.ExpiresAt = newLicense.ExpiresAt
	license.Code = newLicense.Code
	license.ProductID = newLicense.ProductID
	license.IsActive = newLicense.IsActive
	if err := db.Updates(&license).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Couldn't create license.", "data:": err})
	}

	return ctx.JSON(fiber.Map{"status:": "success", "message:": "License created", "data:": license})

}
