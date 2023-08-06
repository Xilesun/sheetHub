package main

import (
	"context"
	"flag"

	"github.com/Xilesun/sheethub/server/app"
	"github.com/Xilesun/sheethub/server/infra/config"
	"github.com/Xilesun/sheethub/server/infra/db"
)

func main() {
	configFile := flag.String("config", "config.yaml", "The path to the config file")
	flag.Parse()
	config, err := config.Read(*configFile)
	if err != nil {
		panic(err)
	}
	database, err := db.SetupDB(context.Background(), config.DB)
	if err != nil {
		panic(err)
	}
	app := app.New(database)
	app.Start()
}
