package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/helply/backend/app/models"
	"github.com/helply/backend/platform/database"
)

func GetTicketTopics(ctx *fiber.Ctx) error {
	var topics []models.TicketTopic
	database.Connection().Find(&topics)
	return ctx.JSON(fiber.Map{"status": "success", "message": "", "data": topics})
}
