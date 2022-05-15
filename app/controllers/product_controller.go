package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/helply/backend/app/models"
	"github.com/helply/backend/platform/database"
)

// CreateProduct
// @Summary Create a new product.
// @Tags Product
// @Accept json
// @Produce json
// @Param title body string true "Title"
// @Param content body string true "Content"
// @Param product_id body int true "Product ID"
// @Param category_id body int true "Category ID"
// @Success 200 {object} models.Product
// @Security ApiKeyAuth
// @Router /api/v1/products [post]
func CreateProduct(ctx *fiber.Ctx) error {
	type NewProduct struct {
		Name   string `json:"name"`
		Icon   string `json:"icon"`
		Image  string `json:"image"`
		PageID uint32 `json:"page_id"`
	}
	db := database.Connection()
	newProduct := new(NewProduct)
	if err := ctx.BodyParser(newProduct); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Invalid input data.", "data:": err})
	}
	product := new(models.Product)
	product.Name = newProduct.Name
	product.Icon = newProduct.Icon
	product.Image = newProduct.Image
	product.PageID = newProduct.PageID
	if err := db.Create(&product).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Couldn't create product.", "data:": err})
	}

	return ctx.JSON(fiber.Map{"status:": "success", "message:": "Product created", "data:": product})

}

func DeleteProduct(ctx *fiber.Ctx) error {
	product := &models.Product{}
	err := database.Connection().Delete(&product, "id = ?", ctx.Params("id")).Error
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{"status": "success", "message": "Product deleted."})
}

func GetProducts(ctx *fiber.Ctx) error {
	var products []models.Product
	database.Connection().Find(&products)

	return ctx.JSON(products)
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
	type NewProduct struct {
		Name   string `json:"name"`
		Icon   string `json:"icon"`
		Image  string `json:"image"`
		PageID uint32 `json:"page_id"`
	}
	db := database.Connection()
	newProduct := new(NewProduct)
	if err := ctx.BodyParser(newProduct); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Invalid input data.", "data:": err})
	}
	product := new(models.Product)
	product.Name = newProduct.Name
	product.Icon = newProduct.Icon
	product.Image = newProduct.Image
	product.PageID = newProduct.PageID
	if err := db.Updates(&product).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Couldn't create product.", "data:": err})
	}

	return ctx.JSON(fiber.Map{"status:": "success", "message:": "Product created", "data:": product})

}
