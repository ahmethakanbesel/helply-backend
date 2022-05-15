package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/helply/backend/app/models"
	"github.com/helply/backend/platform/database"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

// Login Function for creating JWT
// @Description Login and get JWT.
// @Summary If credentials are correct, return JWT.
// @Tags Auth
// @Accept json
// @Produce json
// @Param identity body string true "Identity (email or username)"
// @Param password body string true "Password"
// @Success 200 {string} status "ok"
// @Router /api/v1/auth/login [post]
func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Identity string `json:"identity"`
		Password string `json:"password"`
	}
	input := new(LoginInput)
	if err := c.BodyParser(&input); err != nil {
		err := c.SendStatus(fiber.StatusUnauthorized)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{"status": "error", "message": "Invalid email or password"})
	}
	identity := input.Identity
	pass := input.Password
	user := &models.User{}
	err := database.Connection().First(&user, "email = ?", identity).Error
	if err != nil {
		return c.JSON(fiber.Map{"status": "error", "message": "Invalid email or password"})
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
	if user.Email != identity || err != nil {
		err := c.SendStatus(fiber.StatusUnauthorized)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{"status": "error", "message": "Invalid email or password"})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["identity"] = identity
	claims["expires"] = time.Now().Add(time.Hour * 72).Unix()
	claims["role"] = ""
	userRole := &models.UserRole{}
	err = database.Connection().First(&userRole, "id = ?", user.UserRoleID).Error
	if err == nil {
		claims["role"] = userRole.Name
	}

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		err := c.SendStatus(fiber.StatusInternalServerError)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{"status": "error", "message": "Couldn't sign token"})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": user, "token": t})
}
