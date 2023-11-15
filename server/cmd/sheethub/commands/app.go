package commands

import (
	"context"

	"github.com/Xilesun/sheethub/server/app"
	"github.com/Xilesun/sheethub/server/infra/config"
	"github.com/Xilesun/sheethub/server/infra/db"
	"github.com/Xilesun/sheethub/server/infra/logger"
	"github.com/spf13/cobra"
)

var sheethub app.IApp
var configFile string

var appCmd = &cobra.Command{
	Use:   "sheethub",
	Short: "Organize, import, export, concatenate sheet files on web application.",
	Long:  "Organize, import, export, concatenate sheet files on web application.",
}

func init() {
	cobra.OnInitialize(func() {
		config.Init(configFile)
	}, setupApp)
	appCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "Config file, default is ./config.yaml")
}

func setupApp() {
	config, _ := config.Read()
	database, err := db.SetupDB(context.Background(), config.DB)
	if err != nil {
		logger.Errorf("Failed to setup database: %v", err)
	}
	sheethub = app.New(database)
}

// Execute executes the root command.
func Execute() {
	if err := appCmd.Execute(); err != nil {
		panic(err)
	}
}
