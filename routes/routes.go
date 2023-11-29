package routes

import (
	"github.com/ahmedibra28/go-fiber-boilerplate/controller"
	"github.com/ahmedibra28/go-fiber-boilerplate/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	// auth
	app.Post("/api/v1/login", controller.Login)

	app.Use(middleware.AuthMiddleware)

	// users
	app.Get("/api/v1/users", controller.GetUsers)
	app.Post("/api/v1/users", controller.CreateUser)
	app.Put("/api/v1/users/:id", controller.UpdateUser)
	app.Delete("/api/v1/users/:id", controller.DeleteUser)

	// permissions
	app.Get("/api/v1/permissions", controller.GetPermissions)
	app.Post("/api/v1/permissions", controller.CreatePermission)
	app.Put("/api/v1/permissions/:id", controller.UpdatePermission)
	app.Delete("/api/v1/permissions/:id", controller.DeletePermission)

	// client-permissions
	app.Get("/api/v1/client-permissions", controller.GetClientPermissions)
	app.Post("/api/v1/client-permissions", controller.CreateClientPermission)
	app.Put("/api/v1/client-permissions/:id", controller.UpdateClientPermission)
	app.Delete("/api/v1/client-permissions/:id", controller.DeleteClientPermission)

	// roles
	app.Get("/api/v1/roles", controller.GetRoles)
	app.Post("/api/v1/roles", controller.CreateRole)
	app.Put("/api/v1/roles/:id", controller.UpdateRole)
	app.Delete("/api/v1/roles/:id", controller.DeleteRole)

}
