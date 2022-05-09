package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/helply/backend/app/models"
)

// CreateArticle
// @Summary Create a new article.
// @Tags Article
// @Accept json
// @Produce json
// @Param title body string true "Title"
// @Param content body string true "Content"
// @Param product_id body int true "Product ID"
// @Param category_id body int true "Category ID"
// @Success 200 {object} models.Article
// @Security ApiKeyAuth
// @Router /api/v1/articles [post]
func CreateArticle(ctx *fiber.Ctx) error {
	article := &models.Article{}
	return ctx.JSON(article)
}

func DeleteArticle(ctx *fiber.Ctx) error {
	article := &models.Article{}
	return ctx.JSON(article)
}

func GetArticles(ctx *fiber.Ctx) error {
	article := &models.Article{}
	return ctx.JSON(article)
}

func GetArticle(ctx *fiber.Ctx) error {
	article := &models.Article{}
	return ctx.JSON(article)
}

func UpdateArticle(ctx *fiber.Ctx) error {
	article := &models.Article{}
	return ctx.JSON(article)
}
