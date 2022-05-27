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
	db := database.Connection()
	newProduct := new(dto.ProductDTO)
	if err := ctx.BodyParser(newProduct); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"status:": "error", "message:": "Invalid data given.", "data:": err})
	}
	product := new(models.Product)
	product.Name = newProduct.Name
	product.ImageID = newProduct.ImageID
	if err := db.Create(&product).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Couldn't create product.", "data:": err})
	}
	return ctx.JSON(fiber.Map{"status:": "success", "message:": "Product created", "data:": product})
}

func DeleteTicket(ctx *fiber.Ctx) error {
	product := &models.Product{}
	err := database.Connection().Delete(&product, "id = ?", ctx.Params("id")).Error
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Couldn't delete the product.", "data:": err})
	}

	return ctx.JSON(fiber.Map{"status": "success", "message": "Product deleted."})
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
	return ctx.JSON(fiber.Map{"status:": "success", "message:": "Product created"})
}

func CloseTicket(ctx *fiber.Ctx) error {
	ticket := &models.Ticket{}
	err := database.Connection().First(&ticket, "tickets.id = ?", ctx.Params("id")).Error
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Could not get the ticket.", "data:": err})
	}
	ticket.TicketStatusID = 2
	database.Connection().Save(ticket)
	return ctx.JSON(fiber.Map{"status": "success", "message": "Ticket closed.", "data": ticket})
}
