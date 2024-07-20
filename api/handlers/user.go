package handlers

import (
	"fmt"
	"time"

	"ai.zustack.backend/internal/database"
	"ai.zustack.backend/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	var payload database.User
	if err := c.BodyParser(&payload); err != nil {
		c.SendStatus(fiber.StatusBadRequest)
	}

	if len(payload.Username) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "username must not be empty.",
			"code":  "7013",
		})
	}

	if len(payload.Username) > 55 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "username must be less than 55 characters.",
		})
	}

	if len(payload.Password) > 55 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "password must be less than 55 characters.",
		})
	}

	if payload.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "password must not be empty.",
		})
	}

	user, err := database.GetUserByUsername(payload.Username)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid credentials",
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid credentials",
		})
	}

	tokenByte := jwt.New(jwt.SigningMethodHS256)

	now := time.Now().UTC()
	claims := tokenByte.Claims.(jwt.MapClaims)
	expDuration := time.Hour * 24 * 30
	claims["sub"] = user.ID
	claims["exp"] = now.Add(expDuration).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	tokenString, err := tokenByte.SignedString([]byte(utils.GetEnv("SECRET_KEY")))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot create token",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": tokenString})
}

func Register(c *fiber.Ctx) error {
	var payload database.User
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No se pudo procesar la solicitud.",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot hash password",
		})
	}

	_, err = database.CreateUser(payload.Username, string(hashedPassword))
	if err != nil {
		return c.SendStatus(500)
	}

	return c.SendStatus(201)
}
