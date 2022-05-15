package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/helply/backend/app/models"
	"github.com/helply/backend/platform/database"
)

// CreateUser
// @Summary Create a new user.
// @Tags User
// @Accept json
// @Produce json
// @Param name body string true "Name"
// @Param email body string true "Email"
// @Param password body int true "Password"
// @Success 200 {object} models.User
// @Router /api/v1/users [post]
func CreateUser(ctx *fiber.Ctx) error {
	type NewUser struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	db := database.Connection()
	newUser := new(NewUser)

	if err := ctx.BodyParser(newUser); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Invalid input data.", "data": err})

	}
	user := new(models.User)
	user.Name = newUser.Name
	user.Email = newUser.Email
	user.Password = models.HashPassword(newUser.Password)
	user.UserRoleID = 3
	if err := db.Create(&user).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err})
	}

	return ctx.JSON(fiber.Map{"status": "success", "message": "User created.", "data": user})
}

func GetUsers(ctx *fiber.Ctx) error {
	user := &models.User{}
	return ctx.JSON(user)
}

func GetUser(ctx *fiber.Ctx) error {
	/*
		tokenData, err := utils.ExtractTokenMetadata(ctx)
		if err != nil {
			return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't extract token metadata", "data": err})
		}
		return ctx.Send([]byte(fmt.Sprintf("Hello user with id: %s", tokenData.Identity)))*/

	user := &models.User{}
	return ctx.JSON(user)
}

// UpdateUser
// @Summary Create a new user.
// @Tags User
// @Accept json
// @Produce json
// @Param name body string true "Name"
// @Param email body string true "Email"
// @Param password body int true "Password"
// @Success 200 {object} models.User
// @Security ApiKeyAuth
// @Router /api/v1/users [post]
func UpdateUser(ctx *fiber.Ctx) error {
	user := &models.User{}
	return ctx.JSON(user)
}

func DeleteUser(ctx *fiber.Ctx) error {
	user := &models.User{}
	return ctx.JSON(user)
}
