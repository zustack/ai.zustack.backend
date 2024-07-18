package routes

import (
	"ai.zustack.backend/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func ImageRoutes(app *fiber.App) {
	app.Post("/image/stable-diffusion-3-medium", handlers.GenerateImageSD3M)
}
