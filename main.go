package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/api/auth/:id", func(c *fiber.Ctx) error {
		param := struct {
			ID uint `params:"id"`
		}{}

		c.ParamsParser(&param) // "{"id": 111}"
		if param.ID == 111 {
			return c.SendString(`{"status": "ok"}`)
		}
		return c.Status(fiber.StatusUnauthorized).SendString(`{"status": "error"}`)
	})

	app.Listen(":3000")
}
