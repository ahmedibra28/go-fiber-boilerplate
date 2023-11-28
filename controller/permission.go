package controller

import (
	"github.com/ahmedibra28/go-fiber-boilerplate/config"
	"github.com/ahmedibra28/go-fiber-boilerplate/models"
	"github.com/gofiber/fiber/v2"
)

func GetPermissions(c *fiber.Ctx) error {
	var permissions = []models.Permission{}

	config.DB.Find(&permissions)
	return c.JSON(permissions)
}

func CreatePermission(c *fiber.Ctx) error {
	var permission models.Permission

	if err := c.BodyParser(&permission); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err})
	}

	if err := config.DB.Create(&permission).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(permission)
}

func UpdatePermission(c *fiber.Ctx) error {
	id := c.Params("id")

	var permission models.Permission

	if err := config.DB.First(&permission, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	if err := c.BodyParser(&permission); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err})
	}

	if err := config.DB.Save(&permission).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(permission)

}

func DeletePermission(c *fiber.Ctx) error {
	id := c.Params("id")

	var permission models.Permission

	if err := config.DB.First(&permission, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	if err := config.DB.Delete(&permission, id).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(permission)
}
