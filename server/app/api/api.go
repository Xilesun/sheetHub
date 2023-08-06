package api

import "github.com/gofiber/fiber/v2"

// Routes defines the routes of the API.
func Routes() *fiber.App {
	api := fiber.New()
	return api
}
