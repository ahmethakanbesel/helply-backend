package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/helply/backend/app/models"
	"github.com/helply/backend/platform/database"
)

// CreateArticleTag
// @Summary Create a new article tag.
// @Tags Product
// @Accept json
// @Produce json
// @Param title body string true "Title"
// @Param content body string true "Content"
// @Param product_id body int true "Product ID"
// @Param category_id body int true "Category ID"
// @Success 200 {object} models.ArticleTag
// @Security ApiKeyAuth
// @Router /api/v1/articletags [post]
func CreateArticleTag(ctx *fiber.Ctx) error {
	type NewArticleTag struct {
		Name      string `json:"name"`
		ArticleID uint32 `json:"article_id"`
	}
	db := database.Connection()
	newArticleTag := new(NewArticleTag)
	if err := ctx.BodyParser(newArticleTag); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Invalid input data.", "data:": err})
	}
	articleTag := new(models.ArticleTag)
	articleTag.Name = newArticleTag.Name
	articleTag.ArticleID = newArticleTag.ArticleID
	if err := db.Create(&articleTag).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Couldn't create article tag.", "data:": err})
	}

	return ctx.JSON(fiber.Map{"status:": "success", "message:": "Article tag created", "data:": articleTag})

}

func DeleteArticleTag(ctx *fiber.Ctx) error {
	articleTag := &models.ArticleTag{}
	err := database.Connection().Delete(&articleTag, "id = ?", ctx.Params("id")).Error
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{"status": "success", "message": "Article Tag deleted."})
}

func GetArticleTags(ctx *fiber.Ctx) error {
	var articleTags []models.ArticleTag
	database.Connection().Find(&articleTags)

	return ctx.JSON(articleTags)
}

func GetArticleTag(ctx *fiber.Ctx) error {
	articleTag := &models.ArticleTag{}
	err := database.Connection().First(&articleTag, "id = ?", ctx.Params("id")).Error
	if err != nil {
		return err
	}

	return ctx.JSON(articleTag)
}

func UpdateArticleTag(ctx *fiber.Ctx) error {
	type NewArticleTag struct {
		Name      string `json:"name"`
		ArticleID uint32 `json:"article_id"`
	}
	db := database.Connection()
	newArticleTag := new(NewArticleTag)
	if err := ctx.BodyParser(newArticleTag); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Invalid input data.", "data:": err})
	}
	articleTag := new(models.ArticleTag)
	articleTag.Name = newArticleTag.Name
	articleTag.ArticleID = newArticleTag.ArticleID
	if err := db.Updates(&articleTag).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Couldn't create article tag.", "data:": err})
	}

	return ctx.JSON(fiber.Map{"status:": "success", "message:": "Article tag created", "data:": articleTag})

}
