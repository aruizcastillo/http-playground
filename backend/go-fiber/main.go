package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

type MessageRequest struct {
	Message string `json:"message"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	port := os.Getenv("PORT")
	frontendOrigin := os.Getenv("FRONTEND_ORIGIN")

	app := fiber.New()

	app.Use(func(c fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", frontendOrigin)
		c.Set("Access-Control-Allow-Headers", "Content-Type")
		c.Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusNoContent)
		}

		return c.Next()
	})

	app.Get("/api/message", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello from GO-FIBER backend",
		})
	})

	app.Post("/api/message", func(c fiber.Ctx) error {
		var body MessageRequest

		if err := c.Bind().Body(&body); err != nil {
			return fiber.ErrBadRequest
		}

		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("Message received by GO FIBER backend: %q", body.Message),
		})
	})

	app.Listen(":" + port)
}
