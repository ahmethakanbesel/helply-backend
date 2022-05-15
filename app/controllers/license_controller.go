package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/helply/backend/app/models"
	"github.com/helply/backend/platform/database"
	"time"
)

// CreateLicense
// @Summary Create a new licence.
// @Tags License
// @Accept json
// @Produce json
// @Param title body string true "Title"
// @Param content body string true "Content"
// @Param product_id body int true "Product ID"
// @Param category_id body int true "Category ID"
// @Success 200 {object} models.License
// @Security ApiKeyAuth
// @Router /api/v1/licences [post]
func CreateLicense(ctx *fiber.Ctx) error {
	type NewLicense struct {
		ExpiresAt time.Time `json:"expires_at"`
		Code      string    `json:"code"`
		ProductID uint32    `json:"product_id"`
		IsActive  bool      `json:"is_active"`
	}
	db := database.Connection()
	newLicense := new(NewLicense)
	if err := ctx.BodyParser(newLicense); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Invalid input data.", "data:": err})
	}
	license := new(models.License)
	license.ExpiresAt = newLicense.ExpiresAt
	license.Code = newLicense.Code
	license.ProductID = newLicense.ProductID
	license.IsActive = newLicense.IsActive
	if err := db.Create(&license).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Couldn't create license.", "data:": err})
	}

	return ctx.JSON(fiber.Map{"status:": "success", "message:": "License created", "data:": license})

}

func DeleteLicense(ctx *fiber.Ctx) error {
	license := &models.License{}
	err := database.Connection().Delete(&license, "id = ?", ctx.Params("id")).Error
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{"status": "success", "message": "License deleted."})
}

func GetLicenses(ctx *fiber.Ctx) error {
	var licenses []models.License
	database.Connection().Find(&licenses)

	return ctx.JSON(licenses)
}

func GetLicense(ctx *fiber.Ctx) error {
	license := &models.License{}
	err := database.Connection().First(&license, "id = ?", ctx.Params("id")).Error
	if err != nil {
		return err
	}

	return ctx.JSON(license)
}

func UpdateLicense(ctx *fiber.Ctx) error {
	type NewLicense struct {
		ExpiresAt time.Time `json:"expires_at"`
		Code      string    `json:"code"`
		ProductID uint32    `json:"product_id"`
		IsActive  bool      `json:"is_active"`
	}
	db := database.Connection()
	newLicense := new(NewLicense)
	if err := ctx.BodyParser(newLicense); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Invalid input data.", "data:": err})
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
