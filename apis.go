package main

import (
	"encoding/gob"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
)

func (a *App) SaveServerConfig(req serverConfig) map[string]interface{} {
	fmt.Printf("%+v", req)
	serverList = append(serverList, req)

	file, err := os.OpenFile(serverDataPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		runtime.LogErrorf(a.ctx, "文件创建失败", err.Error())
		return nil
	}
	defer func() {
		_ = file.Close()
	}()
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(serverList)
	if err != nil {
		runtime.LogErrorf(a.ctx, "编码错误", err.Error())
		return nil
	}

	return map[string]interface{}{
		"status": true,
		"code":   1,
		"data":   req,
	}
}

func (a *App) GetServerConfig() interface{} {
	return serverList
}
