package routes

import (
	"ai.zustack.backend/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	app.Post("/login", handlers.Login)
	app.Post("/register", handlers.Register)
}
