package main

import (
	"changeme/constant"
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

var serverDataPath string
var userCfgPath string
var svrDatFile *os.File
var userCfgFile *os.File

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
			runtime.LogErrorf(a.ctx, "gob编码错误,err:%+v", err)
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

	fmt.Println("开始读取配置文件")
	userCfgPath = AppDataPath + `\UserConfig.yaml`
	if ok, _ := PathExists(userCfgPath); !ok {
		fmt.Println("创建文件")
		userCfgFile, err = os.OpenFile(userCfgPath, os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			runtime.LogErrorf(ctx, "文件创建失败,err:%+v", err)
			return
		}
		UserConfig = DefaultConfig //使用默认配置创建
		d, err := yaml.Marshal(&UserConfig)
		if err != nil {
			runtime.LogErrorf(ctx, "yaml序列化处理失败,err:%+v", err)
			return
		}
		_, err = userCfgFile.Write(d)
		if err != nil {
			runtime.LogErrorf(ctx, "用户配置文件写入失败,err:%+v", err)
			return
		}
		err = userCfgFile.Sync()
		if err != nil {
			runtime.LogErrorf(ctx, "用户配置文件写入同步失败,err:%+v", err)
			return
		}

	} else {
		dataBytes, err := os.ReadFile(userCfgPath)
		if err != nil {
			fmt.Println("读取用户配置文件失败：", err)
			return
		}

		err = yaml.Unmarshal(dataBytes, &UserConfig)
		if err != nil {
			runtime.LogErrorf(ctx, "用户配置文件解析失败,err:%+v", err)
			return
		}
		fmt.Printf("配置文件内容%+v\n", UserConfig)

		userCfgFile, err = os.OpenFile(userCfgPath, os.O_RDWR, 0666)
		if err != nil {
			runtime.LogErrorf(ctx, "文件创建失败,err:%+v", err)
			return
		}
	}

	windowTheme := ""
	if _, ok := UserConfig[constant.ConfigKeyWT]; ok {
		windowTheme = UserConfig[constant.ConfigKeyWT].(string)
	} else {
		windowTheme = DefaultConfig[constant.ConfigKeyWT].(string)
	}
	if windowTheme == constant.ThemeSystemDefault {
		runtime.WindowSetSystemDefaultTheme(ctx)
	} else if windowTheme == constant.ThemeDark {
		runtime.WindowSetDarkTheme(ctx)
	} else {
		runtime.WindowSetLightTheme(ctx)
	}

	return
}

func (_ *App) beforeClose(ctx context.Context) (prevent bool) {
	_ = svrDatFile.Close()
	_ = userCfgFile.Close()
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
