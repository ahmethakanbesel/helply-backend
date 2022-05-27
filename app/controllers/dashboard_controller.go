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
	var licenseStats []models.LicenseStats
	database.Connection().Raw("SELECT p.name as id, p.name as label, count(*) as value FROM products p, licenses l WHERE p.id=l.product_id GROUP BY p.name, p.id").Scan(&licenseStats)
	stats.LicenseStats = licenseStats
	return ctx.JSON(fiber.Map{"status:": "success", "message:": "", "data": stats})
}
