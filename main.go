package main

import (
	"embed"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"time"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	//检查应用目录
	usePath := getUserAppDataPath()
	CheckAppDataPath(usePath)

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "ORMTools",
		Width:  1024,
		Height: 768,
		Assets: assets,
		//BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		MinWidth:  800,
		MinHeight: 600,
		OnStartup: app.startup,
		Bind: []interface{}{
			app,
		},
		LogLevel: logger.INFO,
		Logger:   logger.NewFileLogger(fmt.Sprintf(`%s\log\run_log_%s.log`, usePath, time.Now().Format("20060102"))),
		Debug: options.Debug{
			OpenInspectorOnStartup: true, //开启debug
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
