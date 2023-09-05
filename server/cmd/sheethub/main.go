package main

import (
	"context"
	"flag"

	"github.com/Xilesun/sheethub/server/app"
	"github.com/Xilesun/sheethub/server/infra/config"
	"github.com/Xilesun/sheethub/server/infra/db"
	"github.com/Xilesun/sheethub/server/infra/logger"
)

func main() {
	configFile := flag.String("config", "config.yaml", "The path to the config file")
	flag.Parse()
	config, err := config.Read(*configFile)
	if err != nil {
		logger.Errorf("Failed to read config: %v", err)
	}
	database, err := db.SetupDB(context.Background(), config.DB)
	if err != nil {
		logger.Errorf("Failed to setup database: %v", err)
	}
	app := app.New(database)
	if err := app.Start(); err != nil {
		logger.Panicf("Failed to start application: %v", err)
	}
}
