package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/helply/backend/app/dto"
	"github.com/helply/backend/app/models"
	"github.com/helply/backend/pkg/helpers"
	"github.com/helply/backend/platform/database"
)

func CreateTicketReply(ctx *fiber.Ctx) error {
	db := database.Connection()
	newReply := new(dto.TicketReplyDTO)
	if err := ctx.BodyParser(newReply); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"status:": "error", "message:": "Invalid data given.", "data:": err})
	}
	reply := new(models.TicketReply)
	reply.TicketID = newReply.TicketID
	reply.Content = newReply.Content
	claims, err := helpers.ExtractTokenMetadata(ctx)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Couldn't get the user information.", "data:": err})
	}
	// @TODO: Check user permissions
	reply.UserID = claims.ID
	if err := db.Create(&reply).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Couldn't create a ticket reply.", "data:": err})
	}
	return ctx.JSON(fiber.Map{"status:": "success", "message:": "Ticket reply created.", "data:": reply})
}

func GetTicketReplies(ctx *fiber.Ctx) error {
	var ticketReplies []models.TicketReply
	database.Connection().Order("created_at desc").Joins("User").Preload("User.Photo").Preload("Ticket.Product").Preload("Ticket.TicketTopic").Find(&ticketReplies, "\"User\".\"user_role_id\" = ?", 3)

	return ctx.JSON(fiber.Map{"status": "success", "message": "", "data": ticketReplies})
}

func GetTicketRepliesById(ctx *fiber.Ctx) error {
	var ticketReplies []models.TicketReply
	database.Connection().Order("created_at desc").Joins("User").Preload("User.Photo").Find(&ticketReplies, "ticket_id = ?", ctx.Params("id"))

	return ctx.JSON(fiber.Map{"status": "success", "message": "", "data": ticketReplies})
}
