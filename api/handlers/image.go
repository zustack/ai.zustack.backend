package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"ai.zustack.backend/internal/database"
	"ai.zustack.backend/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func GetImages(c *fiber.Ctx) error {
	images, err := database.GetImages()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get images",
		})
	}
	return c.JSON(images)
}

func GenerateImage(c *fiber.Ctx) error {
	var payload database.Image
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	if len(payload.Prompt) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Prompt is required",
		})
	}

	if len(payload.Prompt) >= 155 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Prompt is too long, only 155 characters are allowed",
		})
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to prepare request",
		})
	}

	resp, err := http.Post(utils.GetEnv("SECRET_SOUCE_URL"), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to send request to image generation server",
			"details": err.Error(),
		})
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to read response from image generation server",
		})
	}

	if resp.StatusCode != http.StatusOK {
		return c.Status(resp.StatusCode).Send(body)
	}

	_, err = database.GenerateImage(payload.Prompt, string(body), "1", true)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.SendStatus(200)

}
