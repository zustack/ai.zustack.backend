package routes

import (
	"ai.zustack.backend/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func ImageRoutes(app *fiber.App) {
	app.Post("/generate/image", handlers.GenerateImage)
	app.Get("/get/images", handlers.GetImages)
}
