package app

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Xilesun/sheethub/client"
	"github.com/Xilesun/sheethub/server/app/api"
	"github.com/Xilesun/sheethub/server/infra/config"
	"github.com/Xilesun/sheethub/server/infra/db"
	"github.com/Xilesun/sheethub/server/infra/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

// IApp is the interface that defines the application.
type IApp interface {
	Install(ctx context.Context) error
	IsInstalled() bool
	Start() error
}

// App is the implementation of the application.
type App struct {
	DB *db.DB
	*fiber.App
}

// New creates a new application.
func New(db *db.DB) IApp {
	app := &App{
		DB:  db,
		App: fiber.New(),
	}
	app.Use(logger.Middleware)
	return app
}

// IsInstalled returns if the application is installed.
func (app *App) IsInstalled() bool {
	return true
}

// Install initializes the application.
func (app *App) Install(ctx context.Context) error {
	if app.IsInstalled() {
		return nil
	}
	return app.DB.Migrator.Up(ctx)
}

// Start starts the application.
func (app *App) Start() error {
	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(client.EmbedDirStatic),
		PathPrefix: "dist",
	}))
	app.Mount("/api", api.Routes())

	port := strconv.Itoa(config.Get().App.Port)
	return app.Listen(":" + port)
}
