package main

import (
	"context"
	"encoding/gob"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gopkg.in/yaml.v3"
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
var userCfgFile *os.File

// startup is called when the app starts. The context is saved
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	var err error
	serverDataPath := AppDataPath + `\server.dat`
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

	userCfgPath := AppDataPath + `\UserConfig.yaml`
	if ok, _ := PathExists(userCfgPath); !ok {
		userCfgFile, err = os.OpenFile(userCfgPath, os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			runtime.LogErrorf(ctx, "文件创建失败,err:%+v", err)
			return
		}
		d, errX := yaml.Marshal(&UserConfig)
		if errX != nil {
			runtime.LogErrorf(ctx, "yaml序列化处理失败,err:%+v", errX)
			return
		}
		_, err = userCfgFile.Write(d)
		if err != nil {
			runtime.LogErrorf(ctx, "用户配置文件写入失败,err:%+v", errX)
			return
		}
		err = userCfgFile.Sync()
		if err != nil {
			runtime.LogErrorf(ctx, "用户配置文件写入同步失败,err:%+v", errX)
			return
		}

	} else {
		dataBytes, err := os.ReadFile(userCfgPath)
		if err != nil {
			fmt.Println("读取用户配置文件失败：", err)
			return
		}
		fmt.Println("yaml 文件的内容: \n", string(dataBytes))
		err = yaml.Unmarshal(dataBytes, &UserConfig)
		if err != nil {
			runtime.LogErrorf(ctx, "用户配置文件解析失败,err:%+v", err)
			return
		}

		userCfgFile, err = os.OpenFile(userCfgPath, os.O_RDWR, 0666)
		if err != nil {
			runtime.LogErrorf(ctx, "文件创建失败,err:%+v", err)
			return
		}
	}

	return
}

func (_ *App) beforeClose(ctx context.Context) (prevent bool) {
	_ = svrDatFile.Close()
	for _, v := range ServerConnMap {
		sqlDB, _ := v.DB.DB()
		_ = sqlDB.Close()
		if v.SshClient != nil {
			_ = v.SshClient.Close()
		}
	}
	runtime.LogInfof(ctx, "应用主动退出!")
	return false
}
