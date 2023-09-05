package api

import "github.com/gofiber/fiber/v2"

// Routes defines the routes of the API.
func Routes() *fiber.App {
	api := fiber.New()

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	return api
}
