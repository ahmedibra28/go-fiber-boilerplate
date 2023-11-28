package controller

import (
	"github.com/ahmedibra28/go-fiber-boilerplate/config"
	"github.com/ahmedibra28/go-fiber-boilerplate/models"
	"github.com/ahmedibra28/go-fiber-boilerplate/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Login(c *fiber.Ctx) error {
	var login struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var user models.User

	if err := c.BodyParser(&login); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err})
	}

	if err := config.DB.Preload("Role",
		func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name", "type")
		},
	).First(&user, "email = ?", login.Email).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Incorrect Credentials"})
	}

	if check := utils.CheckPassword(user.Password, login.Password); !check {
		return c.Status(401).JSON(fiber.Map{"error": "Incorrect Credentials"})
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err})
	}

	return c.JSON(fiber.Map{
		"id":     user.ID,
		"name":   user.Name,
		"email":  user.Email,
		"image":  user.Image,
		"mobile": user.Mobile,
		"role": fiber.Map{
			"id":   user.Role.ID,
			"type": user.Role.Type,
		},
		"token": token,
	})
}
