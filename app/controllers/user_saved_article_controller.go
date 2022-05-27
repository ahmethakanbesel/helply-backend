package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/helply/backend/app/dto"
	"github.com/helply/backend/app/models"
	"github.com/helply/backend/pkg/helpers"
	"github.com/helply/backend/platform/database"
)

func CreateUserSavedArticle(ctx *fiber.Ctx) error {
	claims, err := helpers.ExtractTokenMetadata(ctx)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"status:": "error", "message:": "Couldn't get the user information.", "data:": err})
	}
	db := database.Connection()
	newSavedArticle := new(dto.UserSavedArticleDTO)
	if err = ctx.BodyParser(newSavedArticle); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid data given.", "data": err})
	}
	savedArticle := new(models.UserSavedArticle)
	savedArticle.ArticleID = newSavedArticle.ArticleID
	savedArticle.UserID = claims.ID
	if err = db.Create(&savedArticle).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't save the article.", "data": err})
	}
	return ctx.JSON(fiber.Map{"status": "success", "message": "Article saved.", "data": savedArticle})
}
