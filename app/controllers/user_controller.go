package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/helply/backend/app/dto"
	"github.com/helply/backend/app/models"
	"github.com/helply/backend/pkg/helpers"
	"github.com/helply/backend/platform/database"
	"os"
	"strconv"
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
	newUser := new(dto.UserRegisterDTO)
	if err := ctx.BodyParser(newUser); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Invalid data given.", "data": err})
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
	var users []models.User
	database.Connection().Joins("UserRole").Joins("Photo").Find(&users)
	return ctx.JSON(fiber.Map{"status": "success", "message": "", "data": users})
}

func GetUser(ctx *fiber.Ctx) error {
	claims, err := helpers.ExtractTokenMetadata(ctx)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't get user", "data": err})
	}
	if claims.Role == "customer" && ctx.Params("id") != strconv.Itoa(int(claims.ID)) {
		return ctx.Status(403).JSON(fiber.Map{"status": "error", "message": "You can't access the user."})
	}
	user := &models.User{}
	err = database.Connection().First(&user, "id = ?", ctx.Params("id")).Error
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Couldn't get the user.", "data:": err})
	}

	return ctx.JSON(fiber.Map{"status": "success", "message": "", "data": user})
}

func UpdateUser(ctx *fiber.Ctx) error {
	claims, err := helpers.ExtractTokenMetadata(ctx)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't get user", "data": err})
	}
	newUser := new(dto.UserDTO)
	if err = ctx.BodyParser(newUser); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status": "error", "message": "Invalid data given.", "data": err})

	}
	user := &models.User{}
	err = database.Connection().First(&user, "id = ?", claims.ID).Error
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"status:": "error", "message:": "Couldn't get the user.", "data:": err})
	}
	if newUser.Name != "" {
		user.Name = newUser.Name
	}
	if newUser.Email != "" {
		user.Email = newUser.Email
	}
	if newUser.Phone != "" {
		user.Phone = newUser.Phone
	}
	if newUser.PhotoID > 0 {
		user.PhotoID = newUser.PhotoID
	}
	if newUser.Password != "" {
		user.Password = models.HashPassword(newUser.Password)
	}
	database.Connection().Save(user)
	return ctx.JSON(fiber.Map{"status": "success", "message": "User updated.", "data": user})
}

func DeleteUser(ctx *fiber.Ctx) error {
	user := &models.User{}
	err := database.Connection().Delete(&user, "id = ?", ctx.Params("id")).Error
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"status:": "error", "message:": "Couldn't delete the user.", "data:": err})
	}
	return ctx.JSON(fiber.Map{"status": "success", "message": "User deleted."})
}
