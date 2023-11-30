package main

import (
	"embed"

	"github.com/Xilesun/sheethub/app"
	"github.com/Xilesun/sheethub/infra/logger"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := app.New()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "sheethub",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.Start,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		logger.Panicf("Start application failed: %s", err.Error())
	}
}
