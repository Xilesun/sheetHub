package main

import (
	"log"
	"net/http"

	"github.com/Xilesun/sheethub/client"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func main() {
	app := fiber.New()

	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(client.EmbedDirStatic),
		PathPrefix: "dist",
	}))

	log.Fatal(app.Listen(":3000"))
}
