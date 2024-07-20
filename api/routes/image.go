package routes

import (
	"ai.zustack.backend/api/handlers"
	"ai.zustack.backend/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func ImageRoutes(app *fiber.App) {
	app.Post("/generate/image",middleware.DeserializeUser, handlers.GenerateImage)
	app.Get("/get/images", handlers.GetImages)
}
