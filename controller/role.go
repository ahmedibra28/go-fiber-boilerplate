package controller

import (
	"strings"

	"github.com/ahmedibra28/go-fiber-boilerplate/config"
	"github.com/ahmedibra28/go-fiber-boilerplate/models"
	"github.com/gofiber/fiber/v2"
)

func GetRoles(c *fiber.Ctx) error {
	var roles = []models.Role{}

	config.DB.Find(&roles)
	return c.JSON(roles)
}

func validate(role *models.Role) string {
	if strings.TrimSpace(role.Name) == "" {
		return "Name field is required"
	}
	return ""
}

func CreateRole(c *fiber.Ctx) error {
	var role models.Role

	if err := c.BodyParser(&role); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err})
	}

	validationError := validate(&role)
	if validationError != "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": validationError,
		})
	}

	if err := config.DB.Create(&role).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(role)
}

func UpdateRole(c *fiber.Ctx) error {
	id := c.Params("id")

	var role models.Role

	if err := config.DB.First(&role, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	if err := c.BodyParser(&role); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err})
	}

	if err := config.DB.Save(&role).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(role)

}

func DeleteRole(c *fiber.Ctx) error {
	id := c.Params("id")

	var role models.Role

	if err := config.DB.First(&role, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	if err := config.DB.Delete(&role, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(role)
}
