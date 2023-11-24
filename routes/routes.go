package routes

import (
	"github.com/ahmedibra28/go-fiber-boilerplate/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/api/v1/users", controller.GetUsers)

}
