package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/helply/backend/app/dto"
	"github.com/helply/backend/app/models"
	"github.com/helply/backend/pkg/helpers"
	"github.com/helply/backend/platform/database"
	_ "strconv"
)

// CreateArticle
// @Summary Create a new article.
// @Tags Article
// @Accept json
// @Produce json
// @Success 200 {object} models.Article
// @Security ApiKeyAuth
// @Router /api/v1/articles [post]
func CreateArticle(ctx *fiber.Ctx) error {
	db := database.Connection()
	newArticle := new(dto.ArticleDTO)
	if err := ctx.BodyParser(newArticle); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"status:": "error", "message:": "Invalid data given.", "data:": err})
	}
	article := new(models.Article)
	article.Title = newArticle.Title
	article.Content = newArticle.Content
	article.ProductID = newArticle.ProductID
	article.CategoryID = newArticle.CategoryID
	article.ImageID = newArticle.ImageID
	claims, err := helpers.ExtractTokenMetadata(ctx)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Couldn't get the user information.", "data:": err})
	}
	article.AuthorID = claims.ID
	if err := db.Create(&article).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Couldn't create an article.", "data:": err})
	}
	return ctx.JSON(fiber.Map{"status:": "success", "message:": "Article created", "data:": article})
}

func DeleteArticle(ctx *fiber.Ctx) error {
	article := &models.Article{}
	err := database.Connection().Delete(&article, "id = ?", ctx.Params("id")).Error
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{"status": "success", "message": "Article deleted."})
}

func GetArticles(ctx *fiber.Ctx) error {
	var articles []models.Article
	//database.Connection().Joins(clause.Associations).Find(&articles)
	database.Connection().Joins("Category").Joins("Product").Joins("Author").Joins("Image").Preload("Author.Photo").Find(&articles)
	return ctx.JSON(fiber.Map{"status": "success", "message": "", "data": articles})
}

func GetArticle(ctx *fiber.Ctx) error {
	article := &models.Article{}
	err := database.Connection().Joins("Category").Joins("Product").Joins("Author").Joins("Image").Preload("Author.Photo").First(&article, "articles.id = ?", ctx.Params("id")).Error
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Could not get the article.", "data:": err})
	}
	return ctx.JSON(fiber.Map{"status": "success", "message": "", "data": article})
}

func UpdateArticle(ctx *fiber.Ctx) error {
	newArticle := new(dto.ArticleUpdateDTO)
	if err := ctx.BodyParser(newArticle); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid data given.", "data": err})
	}
	article := &models.Article{}
	err := database.Connection().First(&article, "id = ?", ctx.Params("id")).Error
	if err != nil || article.ID <= 0 {
		return ctx.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Couldn't get the article.", "data:": err})
	}
	if newArticle.Title != "" {
		article.Title = newArticle.Title
	}
	if newArticle.ProductID > 0 {
		article.ProductID = newArticle.ProductID
	}
	if newArticle.CategoryID > 0 {
		article.CategoryID = newArticle.CategoryID
	}
	if newArticle.Content != "" {
		article.Content = newArticle.Content
	}
	database.Connection().Save(article)
	return ctx.JSON(fiber.Map{"status": "success", "message": "Article updated.", "data": article})
}

func VoteArticle(ctx *fiber.Ctx) error {
	vote := new(dto.ArticleVoteDTO)
	if err := ctx.BodyParser(vote); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"status:": "error", "message:": "Invalid input data.", "data:": err})
	}
	article := &models.Article{}
	err := database.Connection().First(&article, "id = ?", ctx.Params("id")).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"status:": "error", "message:": "Article not found.", "data:": err})
	}
	if vote.Type == 1 {
		article.Votes = article.Votes + 1
	} else {
		article.Votes = article.Votes - 1
	}
	err = database.Connection().Model(&article).Updates(&models.Article{
		Votes: article.Votes,
	}).Error
	return ctx.JSON(fiber.Map{"status:": "success", "message:": "Article voted.", "data:": article})
}
