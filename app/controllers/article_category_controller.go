package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/helply/backend/app/models"
	"github.com/helply/backend/platform/database"
)

// CreateArticleCategory
// @Summary Create a new article category.
// @Tags ArticleCategory
// @Accept json
// @Produce json
// @Param title body string true "Title"
// @Param content body string true "Content"
// @Param product_id body int true "Product ID"
// @Param category_id body int true "Category ID"
// @Success 200 {object} models.ArticleCategory
// @Security ApiKeyAuth
// @Router /api/v1/article-category [post]
func CreateArticleCategory(ctx *fiber.Ctx) error {
	type NewArticleCategory struct {
		Name string `json:"name"`
	}
	db := database.Connection()
	newArticleCategory := new(NewArticleCategory)
	if err := ctx.BodyParser(newArticleCategory); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"status:": "error", "message:": "Invalid data given.", "data:": err})
	}
	articleCategory := new(models.ArticleCategory)
	articleCategory.Name = newArticleCategory.Name
	if err := db.Create(&articleCategory).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Couldn't create article category.", "data:": err})
	}

	return ctx.JSON(fiber.Map{"status:": "success", "message:": "Article category created", "data:": articleCategory})

}

func DeleteArticleCategory(ctx *fiber.Ctx) error {
	articleCategory := &models.ArticleCategory{}
	err := database.Connection().Delete(&articleCategory, "id = ?", ctx.Params("id")).Error
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{"status": "success", "message": "Article Category deleted."})
}

func GetArticleCategories(ctx *fiber.Ctx) error {
	var categories []models.ArticleCategory
	database.Connection().Find(&categories)
	return ctx.JSON(fiber.Map{"status": "success", "message": "", "data": categories})
}

func GetArticleCategory(ctx *fiber.Ctx) error {
	articleCategory := &models.ArticleCategory{}
	err := database.Connection().First(&articleCategory, "id = ?", ctx.Params("id")).Error
	if err != nil {
		return err
	}

	return ctx.JSON(articleCategory)
}

func UpdateArticleCategory(ctx *fiber.Ctx) error {
	type NewArticleCategory struct {
		Name string `json:"name"`
	}
	db := database.Connection()
	newArticleCategory := new(NewArticleCategory)
	if err := ctx.BodyParser(newArticleCategory); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"status:": "error", "message:": "Invalid data given.", "data:": err})
	}
	articleCategory := new(models.ArticleCategory)
	articleCategory.Name = newArticleCategory.Name
	if err := db.Updates(&articleCategory).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Couldn't create article category.", "data:": err})
	}

	return ctx.JSON(fiber.Map{"status:": "success", "message:": "Article category created", "data:": articleCategory})

}
