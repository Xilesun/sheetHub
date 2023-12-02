package app

import (
	"context"

	"github.com/Xilesun/sheethub/infra/config"
	"github.com/Xilesun/sheethub/infra/db"
	"github.com/Xilesun/sheethub/infra/logger"
)

// IApp is the interface that defines the application.
type IApp interface {
	Start(ctx context.Context)
	Install(ctx context.Context) error
	IsInstalled() bool
}

// App is the implementation of the application.
type App struct {
	ctx context.Context
	DB  *db.DB
}

// New creates a new application.
func New() IApp {
	return &App{}
}

// Start starts the application.
func (app *App) Start(ctx context.Context) {
	conf, err := config.Init()
	if err != nil {
		logger.Errorf("Initialize configuration failed: %s", err.Error())
	}
	database, err := db.SetupDB(ctx, conf.DB)
	if err != nil {
		logger.Errorf("Setup database failed: %s", err.Error())
	}
	app.DB = database
	app.ctx = ctx
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
