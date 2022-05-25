package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/helply/backend/app/dto"
	"github.com/helply/backend/app/models"
	"github.com/helply/backend/platform/database"
)

// CreateProduct
// @Summary Create a new product.
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {object} models.Product
// @Security ApiKeyAuth
// @Router /api/v1/products [post]
func CreateProduct(ctx *fiber.Ctx) error {
	db := database.Connection()
	newProduct := new(dto.ProductDTO)
	if err := ctx.BodyParser(newProduct); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"status:": "error", "message:": "Invalid data given.", "data:": err})
	}
	product := new(models.Product)
	product.Name = newProduct.Name
	product.ImageID = newProduct.ImageID
	if err := db.Create(&product).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Couldn't create product.", "data:": err})
	}
	return ctx.JSON(fiber.Map{"status:": "success", "message:": "Product created", "data:": product})
}

func DeleteProduct(ctx *fiber.Ctx) error {
	product := &models.Product{}
	err := database.Connection().Delete(&product, "id = ?", ctx.Params("id")).Error
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Couldn't delete the product.", "data:": err})
	}

	return ctx.JSON(fiber.Map{"status": "success", "message": "Product deleted."})
}

func GetProducts(ctx *fiber.Ctx) error {
	var products []models.Product
	database.Connection().Joins("Image").Joins("Icon").Find(&products)

	return ctx.JSON(fiber.Map{"status": "success", "message": "", "data": products})
}

func GetProduct(ctx *fiber.Ctx) error {
	product := &models.Product{}
	err := database.Connection().First(&product, "id = ?", ctx.Params("id")).Error
	if err != nil {
		return err
	}

	return ctx.JSON(product)
}

func UpdateProduct(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"status:": "success", "message:": "Product created"})
}
