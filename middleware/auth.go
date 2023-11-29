package middleware

import (
	"strings"

	"github.com/ahmedibra28/go-fiber-boilerplate/config"
	"github.com/ahmedibra28/go-fiber-boilerplate/models"
	"github.com/ahmedibra28/go-fiber-boilerplate/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

func AuthMiddleware(c *fiber.Ctx) error {
	breakToken := c.Get("Authorization")

	if breakToken == "" {
		return c.Status(401).JSON(fiber.Map{
			"error": "Please provide a valid token",
		})
	}

	pureToken := strings.Split(breakToken, " ")[1]

	token, err := utils.ParseJWT(pureToken)

	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user_id := token.Claims.(jwt.MapClaims)["sub"]

	var user models.User

	if err := config.DB.
		Preload("Role",
			func(db *gorm.DB) *gorm.DB {
				return db.Select("id", "name", "type")
			}).
		First(&user, user_id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "User not found",
		})

	}

	c.Locals("userID", user.ID)
	c.Locals("userName", user.Name)
	c.Locals("userMobile", user.Mobile)
	c.Locals("userEmail", user.Email)
	c.Locals("roleID", user.Role.ID)
	c.Locals("roleName", user.Role.Name)
	c.Locals("roleType", user.Role.Type)

	c.Next()
	return nil
}
