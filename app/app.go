package app

import (
	"context"

	"github.com/Xilesun/sheethub/infra/config"
	"github.com/Xilesun/sheethub/infra/db"
	"github.com/Xilesun/sheethub/infra/logger"
)

// App is the implementation of the application.
type App struct {
	ctx context.Context
	DB  *db.DB
}

// New creates a new application.
func New() *App {
	return &App{}
}

// Init initializes the application.
func (app *App) Init(ctx context.Context) {
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
	err = app.DB.Migrator.Up()
	if err != nil {
		logger.Errorf("Migrate database failed: %s", err.Error())
	}
}
