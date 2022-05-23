package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/helply/backend/app/dto"
	"github.com/helply/backend/app/models"
	"github.com/helply/backend/pkg/helpers"
	"github.com/helply/backend/platform/database"
	"os"
	"time"
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
	db := database.Connection()
	newUser := new(dto.UserDTO)

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
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["identity"] = user.Email
	claims["expires"] = time.Now().Add(time.Hour * 72).Unix()
	claims["role"] = user.UserRoleID
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		err := ctx.SendStatus(fiber.StatusInternalServerError)
		if err != nil {
			return err
		}
		return ctx.JSON(fiber.Map{"status": "error", "message": "Couldn't sign token"})
	}

	return ctx.JSON(fiber.Map{"status": "success", "message": "User created.", "data": user, "token": t})
}

func GetUsers(ctx *fiber.Ctx) error {
	user := &models.User{}
	return ctx.JSON(user)
}

func GetUser(ctx *fiber.Ctx) error {
	claims, err := helpers.ExtractTokenMetadata(ctx)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't get user", "data": err})
	}
	id := claims.Identity
	return ctx.Send([]byte(fmt.Sprintf("Hello user with id: %s", id)))
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
