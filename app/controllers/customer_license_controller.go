package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/helply/backend/app/dto"
	"github.com/helply/backend/app/models"
	"github.com/helply/backend/pkg/helpers"
	"github.com/helply/backend/platform/database"
)

// CreateCustomerLicense
// @Summary Create a new licence.
// @Tags CustomerLicense
// @Accept json
// @Produce json
// @Param title body string true "Title"
// @Param content body string true "Content"
// @Param product_id body int true "Product ID"
// @Param category_id body int true "Category ID"
// @Success 200 {object} models.CustomerLicense
// @Security ApiKeyAuth
// @Router /api/v1/customer-licenses [post]
func CreateCustomerLicense(ctx *fiber.Ctx) error {
	db := database.Connection()
	data := new(dto.CustomerLicenseDto)
	if err := ctx.BodyParser(data); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"status:": "error", "message:": "Invalid data given.", "data:": err})
	}
	claims, err := helpers.ExtractTokenMetadata(ctx)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"status": "error", "message": "Couldn't get user", "data": err})
	}
	customerLicense := new(models.CustomerLicense)
	customerLicense.CustomerID = claims.ID
	license := &models.License{}
	err = database.Connection().First(&license, "code = ?", data.LicenseCode).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"status": "error", "message": "License code is not valid.", "data": err})
	}
	customerLicense.LicenseID = license.ID
	if err := db.Create(&customerLicense).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Couldn't create Customer License.", "data:": err})
	}

	return ctx.JSON(fiber.Map{"status:": "success", "message:": "Customer license created", "data:": customerLicense})

}

func DeleteCustomerLicense(ctx *fiber.Ctx) error {
	customerLicense := &models.CustomerLicense{}
	err := database.Connection().Delete(&customerLicense, "id = ?", ctx.Params("id")).Error
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{"status": "success", "message": "Customer License deleted."})
}

func GetCustomerLicenses(ctx *fiber.Ctx) error {
	var customerLicenses []models.CustomerLicense
	database.Connection().Find(&customerLicenses)

	return ctx.JSON(customerLicenses)
}

func GetCustomerLicense(ctx *fiber.Ctx) error {
	customerLicense := &models.CustomerLicense{}
	err := database.Connection().First(&customerLicense, "id = ?", ctx.Params("id")).Error
	if err != nil {
		return err
	}

	return ctx.JSON(customerLicense)
}

func UpdateCustomerLicense(ctx *fiber.Ctx) error {
	type NewCustomerLicense struct {
		CustomerID uint32 ` json:"customer_id" `
		LicenseID  uint32 `json:"license_id" `
	}
	db := database.Connection()
	newCustomerLicense := new(NewCustomerLicense)
	if err := ctx.BodyParser(newCustomerLicense); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Invalid input data.", "data:": err})
	}
	customerLicense := new(models.CustomerLicense)
	customerLicense.CustomerID = newCustomerLicense.CustomerID
	customerLicense.LicenseID = newCustomerLicense.LicenseID
	if err := db.Updates(&customerLicense).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Couldn't create customer License.", "data:": err})
	}

	return ctx.JSON(fiber.Map{"status:": "success", "message:": "customer License created", "data:": customerLicense})

}
