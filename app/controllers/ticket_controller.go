package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/helply/backend/app/dto"
	"github.com/helply/backend/app/models"
	"github.com/helply/backend/pkg/helpers"
	"github.com/helply/backend/platform/database"
)

// CreateTicket
// @Summary Create a new ticket.
// @Tags Ticket
// @Accept json
// @Produce json
// @Success 200 {object} models.Ticket
// @Security ApiKeyAuth
// @Router /api/v1/products [post]
func CreateTicket(ctx *fiber.Ctx) error {
	claims, err := helpers.ExtractTokenMetadata(ctx)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"status:": "error", "message:": "Couldn't get the user information.", "data:": err})
	}
	newTicket := new(dto.TicketDTO)
	if err = ctx.BodyParser(newTicket); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"status:": "error", "message:": "Invalid data given.", "data:": err})
	}
	db := database.Connection()
	ticket := new(models.Ticket)
	ticket.TicketTopicID = newTicket.TopicID
	ticket.ProductID = newTicket.ProductID
	ticket.TicketStatusID = 1
	ticket.CustomerID = claims.ID
	if err = db.Create(&ticket).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Couldn't create a ticket.", "data:": err})
	}
	ticketReply := new(models.TicketReply)
	ticketReply.TicketID = ticket.ID
	ticketReply.UserID = claims.ID
	ticketReply.Content = newTicket.Content
	if err = db.Create(&ticketReply).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Couldn't create a ticket reply.", "data:": err})
	}
	return ctx.JSON(fiber.Map{"status:": "success", "message:": "Ticket created.", "data:": ticket})
}

func DeleteTicket(ctx *fiber.Ctx) error {
	ticket := &models.Ticket{}
	err := database.Connection().Delete(&ticket, "id = ?", ctx.Params("id")).Error
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Couldn't delete the ticket.", "data:": err})
	}
	return ctx.JSON(fiber.Map{"status": "success", "message": "Ticket deleted."})
}

func GetTickets(ctx *fiber.Ctx) error {
	var tickets []models.Ticket
	claims, err := helpers.ExtractTokenMetadata(ctx)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"status:": "error", "message:": "Couldn't get the user information.", "data:": err})
	}
	if claims.Role == "customer" {
		database.Connection().Joins("Customer").Joins("Product").Joins("TicketTopic").Joins("TicketStatus").Find(&tickets, "\"Customer\".\"id\" = ?", claims.ID)
	} else {
		database.Connection().Joins("Customer").Joins("Product").Joins("TicketTopic").Joins("TicketStatus").Find(&tickets)
	}
	return ctx.JSON(fiber.Map{"status": "success", "message": "", "data": tickets})
}

func GetTicket(ctx *fiber.Ctx) error {
	ticket := &models.Ticket{}
	err := database.Connection().Joins("Customer").Joins("Product").Joins("TicketTopic").Joins("TicketStatus").First(&ticket, "tickets.id = ?", ctx.Params("id")).Error
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Could not get the ticket.", "data:": err})
	}
	return ctx.JSON(fiber.Map{"status": "success", "message": "", "data": ticket})
}

func UpdateTicket(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"status:": "success", "message:": "Ticket updated."})
}

func CloseTicket(ctx *fiber.Ctx) error {
	ticket := &models.Ticket{}
	database.Connection().Model(&models.Ticket{}).Where("id = ?", ctx.Params("id")).Update("ticket_status_id", 2)
	return ctx.JSON(fiber.Map{"status": "success", "message": "Ticket closed.", "data": ticket})
}
