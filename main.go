package main

import (
	"embed"
	"fmt"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"time"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	//检查应用目录
	usePath := getUserAppDataPath()
	CheckAppDataPath(usePath)

	err := wails.Run(&options.App{
		Title:  "ORMTools",
		Width:  1024,
		Height: 768,
		//Assets: assets,
		AssetServer: &assetserver.Options{
			Assets: assets,
			//Handler: NewFileLoader(),
		},
		MinWidth:      800,
		MinHeight:     600,
		OnStartup:     app.startup,
		OnBeforeClose: app.beforeClose,
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
