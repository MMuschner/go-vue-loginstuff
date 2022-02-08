package routes

import (
	"github.com/gofiber/fiber/v2"
	"loginstuff/controllers"
)

func Setup(app *fiber.App) {
	// Setup api routes
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Get("/api/logout", controllers.Logout)
}
