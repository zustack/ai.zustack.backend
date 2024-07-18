package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"ai.zustack.backend/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type InputGenerateImage struct {
	Prompt         string `json:"prompt"`
	CfgScale       int    `json:"cfg_scale"`
	AspectRatio    string `json:"aspect_ratio"`
	Seed           int    `json:"seed"`
	Steps          int    `json:"steps"`
	NegativePrompt string `json:"negative_prompt"`
}

func GenerateImageSD3M(c *fiber.Ctx) error {
	var payload InputGenerateImage
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body.",
		})
	}

	invokeURL := "https://ai.api.nvidia.com/v1/genai/stabilityai/stable-diffusion-3-medium"
	nvidiaAPIKey := utils.GetEnv("NVIDIA_API_KEY")
	authHeader := fmt.Sprintf("Bearer %s", nvidiaAPIKey)
	acceptHeader := "application/json"

	data := InputGenerateImage{
		Prompt:         payload.Prompt,
		CfgScale:       payload.CfgScale,
		AspectRatio:    payload.AspectRatio,
		Seed:           payload.Seed,
		Steps:          payload.Steps,
		NegativePrompt: payload.NegativePrompt,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	req, err := http.NewRequest("POST", invokeURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	req.Header.Set("Authorization", authHeader)
	req.Header.Set("Accept", acceptHeader)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(resp.StatusCode).SendString(string(body))
}
