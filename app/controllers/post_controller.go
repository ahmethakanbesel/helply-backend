package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/helply/backend/app/models"
	"github.com/helply/backend/platform/database"
)

// CreatePost CreateBook func for creates a new book.
// @Description Create a new book.
// @Summary create a new book
// @Tags Book
// @Accept json
// @Produce json
// @Param title body string true "Title"
// @Param author body string true "Author"
// @Param book_attrs body models.Post true "Book attributes"
// @Success 200 {object} models.Post
// @Security ApiKeyAuth
// @Router /api/v1/book [post]
func CreatePost(ctx *fiber.Ctx) error {
	post := models.Post{}
	if err := ctx.BodyParser(&post); err != nil {
		return err
	}

	if len(post.Content) < 1 {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Content is required",
		})
	}

	if err := database.Connection().Create(&post).Error; err != nil {
		return err
	}

	return ctx.JSON(post)
}

func DeletePost(ctx *fiber.Ctx) error {
	post := &models.Post{}

	err := database.Connection().First(&post, "id = ?", ctx.Params("id")).Error
	if err != nil {
		return err
	}

	err = database.Connection().Delete(&post).Error
	if err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"message": "Post deleted successfully",
	})
}

func GetPosts(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["sub"].(string)
	return ctx.Send([]byte(fmt.Sprintf("Hello user with id: %s", id)))

	var posts []models.Post
	database.Connection().Find(&posts)

	return ctx.JSON(posts)
}

func GetPost(ctx *fiber.Ctx) error {
	post := &models.Post{}

	err := database.Connection().First(&post, "id = ?", ctx.Params("id")).Error
	if err != nil {
		return err
	}

	return ctx.JSON(post)
}

func UpdatePost(ctx *fiber.Ctx) error {
	request := &models.Post{}

	if err := ctx.BodyParser(&request); err != nil {
		return err
	}

	post := &models.Post{}

	err := database.Connection().First(&post, "id = ?", ctx.Params("id")).Error
	if err != nil {
		return err
	}

	err = database.Connection().Model(&post).Updates(&models.Post{
		Title:   request.Title,
		Content: request.Content,
	}).Error

	if err != nil {
		return err
	}

	return ctx.JSON(post)
}
