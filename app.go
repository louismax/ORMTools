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

var svrDatFile *os.File

// startup is called when the app starts. The context is saved
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	var err error
	serverDataPath = AppDataPath + `\server.dat`
	ServerConfigMap = make(map[string]ServerConfig)
	ServerConnMap = make(map[string]ServerConn)
	if ok, _ := PathExists(serverDataPath); !ok {
		svrDatFile, err = os.OpenFile(serverDataPath, os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			runtime.LogErrorf(ctx, "文件创建失败,err:%+v", err)
			return
		}

		encoder := gob.NewEncoder(svrDatFile)
		err = encoder.Encode(ServerConfigMap)
		if err != nil {
			runtime.LogErrorf(ctx, "编码错误,err:%+v", err)
			return
		}
	} else {
		svrDatFile, err = os.OpenFile(serverDataPath, os.O_RDWR, 0666)
		if err != nil {
			runtime.LogErrorf(ctx, "文件创建失败,err:%+v", err)
			return
		}
		decoder := gob.NewDecoder(svrDatFile)
		err = decoder.Decode(&ServerConfigMap)
		if err != nil {
			runtime.LogErrorf(ctx, "解码失败,err:%+v", err)
			return
		}
	}
	return
}

func (_ *App) beforeClose(ctx context.Context) (prevent bool) {
	//dialog, err := runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
	//	Type:    runtime.InfoDialog,
	//	Title:   "Quit?",
	//	Message: "Are you sure you want to quit?",
	//})
	//
	//if err != nil {
	//	return false
	//}
	//return dialog != "Yes"
	_ = svrDatFile.Close()
	for _, v := range ServerConnMap {
		sqlDB, _ := v.DB.DB()
		_ = sqlDB.Close()
		if v.SshClient != nil {
			_ = v.SshClient.Close()
		}
	}
	return false
}
