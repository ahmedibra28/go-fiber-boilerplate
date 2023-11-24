package controller

import (
	"github.com/ahmedibra28/go-fiber-boilerplate/config"
	"github.com/ahmedibra28/go-fiber-boilerplate/models"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {

	var users = []models.User{}

	config.DB.Find(&users)

	return c.JSON(fiber.Map{
		"status": 200,
		"total":  len(users),
		"data":   users,
	})
}
