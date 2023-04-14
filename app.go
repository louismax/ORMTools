package main

import (
	"context"
	"encoding/gob"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	serverDataPath = AppDataPath + `\server.dat`
	ServerConfigMap = make(map[string]ServerConfig)
	ServerConnMap = make(map[string]ServerConn)
	if ok, _ := PathExists(serverDataPath); !ok {
		file, err := os.OpenFile(serverDataPath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			runtime.LogErrorf(ctx, "文件创建失败,err:%+v", err)
			return
		}
		defer func() {
			_ = file.Close()
		}()
		encoder := gob.NewEncoder(file)
		err = encoder.Encode(ServerConfigMap)
		if err != nil {
			runtime.LogErrorf(ctx, "编码错误,err:%+v", err)
			return
		}
	} else {
		file, err := os.Open(serverDataPath)
		if err != nil {
			runtime.LogErrorf(ctx, "文件打开失败,err:%+v", err)
			return
		}
		defer func() {
			_ = file.Close()
		}()
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(&ServerConfigMap)
		if err != nil {
			runtime.LogErrorf(ctx, "解码失败,err:%+v", err)
			return
		}
	}
	return
}
