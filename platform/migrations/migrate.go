package migrations

import (
	"github.com/helply/backend/app/models"
	"github.com/helply/backend/platform/database"
)

func Migrate() {
	database.Connection().AutoMigrate(&models.Article{})
	database.Connection().AutoMigrate(&models.ArticleCategory{})
	database.Connection().AutoMigrate(&models.ArticleTag{})
	database.Connection().AutoMigrate(&models.File{})
	database.Connection().AutoMigrate(&models.Language{})
	database.Connection().AutoMigrate(&models.LanguageString{})
	database.Connection().AutoMigrate(&models.License{})
	database.Connection().AutoMigrate(&models.CustomerLicense{})
	database.Connection().AutoMigrate(&models.Page{})
	database.Connection().AutoMigrate(&models.Permission{})
	database.Connection().AutoMigrate(&models.RolePermission{})
	database.Connection().AutoMigrate(&models.Product{})
	database.Connection().AutoMigrate(&models.Ticket{})
	database.Connection().AutoMigrate(&models.TicketStatus{})
	database.Connection().AutoMigrate(&models.TicketReply{})
	database.Connection().AutoMigrate(&models.TicketTopic{})
	database.Connection().AutoMigrate(&models.TicketReplyAttachment{})
	database.Connection().AutoMigrate(&models.User{})
	database.Connection().AutoMigrate(&models.UserRole{})
}
