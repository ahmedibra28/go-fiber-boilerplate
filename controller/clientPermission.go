package controller

import (
	"github.com/ahmedibra28/go-fiber-boilerplate/config"
	"github.com/ahmedibra28/go-fiber-boilerplate/models"
	"github.com/gofiber/fiber/v2"
)

func GetClientPermissions(c *fiber.Ctx) error {
	var clientPermissions = []models.ClientPermission{}

	config.DB.Find(&clientPermissions)
	return c.JSON(clientPermissions)
}

func CreateClientPermission(c *fiber.Ctx) error {
	var clientPermission models.ClientPermission

	if err := c.BodyParser(&clientPermission); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err})
	}

	if err := config.DB.Create(&clientPermission).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(clientPermission)
}

func UpdateClientPermission(c *fiber.Ctx) error {
	id := c.Params("id")

	var clientPermission models.ClientPermission

	if err := config.DB.First(&clientPermission, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	if err := c.BodyParser(&clientPermission); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err})
	}

	if err := config.DB.Save(&clientPermission).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(clientPermission)

}

func DeleteClientPermission(c *fiber.Ctx) error {
	id := c.Params("id")

	var clientPermission models.ClientPermission

	if err := config.DB.First(&clientPermission, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	if err := config.DB.Delete(&clientPermission, id).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(clientPermission)
}
