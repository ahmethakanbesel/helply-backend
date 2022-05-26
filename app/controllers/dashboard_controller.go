package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/helply/backend/app/models"
	"github.com/helply/backend/platform/database"
)

func GetAdminStats(ctx *fiber.Ctx) error {
	stats := new(models.Stats)
	database.Connection().Model(&models.Ticket{}).Count(&stats.WidgetData.TicketCount)
	database.Connection().Model(&models.User{}).Count(&stats.WidgetData.UserCount)
	database.Connection().Model(&models.Product{}).Count(&stats.WidgetData.ProductCount)
	database.Connection().Model(&models.License{}).Count(&stats.WidgetData.LicenseCount)
	return ctx.JSON(fiber.Map{"status:": "success", "message:": "", "data": stats})
}
