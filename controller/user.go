package controller

import (
	"github.com/ahmedibra28/go-fiber-boilerplate/config"
	"github.com/ahmedibra28/go-fiber-boilerplate/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User

	if err := config.DB.Preload("Role",
		func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name", "description")
		}).Find(&users).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err})
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err})
	}

	if err := config.DB.Save(&user).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(user)

}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	if err := config.DB.Delete(&user, id).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(user)
}
