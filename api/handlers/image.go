package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"ai.zustack.backend/internal/database"
	"ai.zustack.backend/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func GetImageByID(c *fiber.Ctx) error {
	id := c.Params("id")
	image, err := database.GetImageByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Failed to get image",
		})
	}
	return c.JSON(image)
}

type ImageResponse struct {
	Data       []database.Image `json:"data"`
	PreviousID *int             `json:"previousId"`
	NextID     *int             `json:"nextId"`
}

func GetUserImages(c *fiber.Ctx) error {
	user := c.Locals("user").(*database.User)
	cursor, err := strconv.Atoi(c.Query("cursor", "0"))
	if err != nil {
		return c.Status(400).SendString("Invalid cursor")
	}

	limit := 30
	images, err := database.GetUserImages(user.ID, limit, cursor)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get images",
		})
	}

	totalCount, err := database.GetUserImagesCount(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get images count",
		})
	}

	var previousID, nextID *int
	if cursor > 0 {
		prev := cursor - limit
		if prev < 0 {
			prev = 0
		}
		previousID = &prev
	}
	if cursor+limit < totalCount {
		next := cursor + limit
		nextID = &next
	}

	response := ImageResponse{
		Data:       images,
		PreviousID: previousID,
		NextID:     nextID,
	}

	return c.JSON(response)
}

func GetImages(c *fiber.Ctx) error {
	cursor, err := strconv.Atoi(c.Query("cursor", "0"))
	if err != nil {
		return c.Status(400).SendString("Invalid cursor")
	}

	limit := 30
	searchParam := c.Query("q", "")
	searchParam = "%" + searchParam + "%"
	images, err := database.GetImages(searchParam, limit, cursor)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get images",
		})
	}

	totalCount, err := database.GetImagesCount()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get images count",
		})
	}

	var previousID, nextID *int
	if cursor > 0 {
		prev := cursor - limit
		if prev < 0 {
			prev = 0
		}
		previousID = &prev
	}
	if cursor+limit < totalCount {
		next := cursor + limit
		nextID = &next
	}

	response := ImageResponse{
		Data:       images,
		PreviousID: previousID,
		NextID:     nextID,
	}

	return c.JSON(response)
}

func GenerateImage(c *fiber.Ctx) error {
	user := c.Locals("user").(*database.User)
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

	_, err = database.GenerateImage(payload.Prompt, string(body), user.ID, true)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.SendStatus(200)
}
