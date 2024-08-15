package api

import (
	"ai.zustack.backend/api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Setup() *fiber.App {
	app := fiber.New()
  app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://ai.zustack.com",
		AllowCredentials: true,
	}))
	app.Static("/web/uploads", "./web/uploads")
	routes.UserRoutes(app)
	routes.ImageRoutes(app)
	return app
}
