package main

import (
	"context"
	"encoding/gob"
	"fmt"
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

type serverConfig struct {
	LocalName     string `json:"local_name"`
	DbType        string `json:"dbType"`
	Host          string `json:"host"`
	Port          string `json:"port"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	HasRecordPwd  bool   `json:"has_record_pwd"`
	HasUseSSH     bool   `json:"hasUseSSH"`
	SshHost       string `json:"ssh_host"`
	SshPort       string `json:"ssh_port"`
	SshUser       string `json:"ssh_user"`
	HasSshKeyfile bool   `json:"has_ssh_keyfile"`
	SshKeyfile    string `json:"ssh_keyfile"`
	HasSshPass    bool   `json:"has_ssh_pass"`
	SshPassword   string `json:"ssh_password"`
}

var serverDataPath string
var serverList = make([]serverConfig, 0)

// startup is called when the app starts. The context is saved
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	serverDataPath = AppDataPath + `\server.dat`
	if ok, _ := PathExists(serverDataPath); !ok {
		file, err := os.OpenFile(serverDataPath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			runtime.LogErrorf(ctx, "文件创建失败", err.Error())
			return
		}
		defer func() {
			_ = file.Close()
		}()
		encoder := gob.NewEncoder(file)
		err = encoder.Encode(serverList)
		if err != nil {
			runtime.LogErrorf(ctx, "编码错误", err.Error())
			return
		}
	} else {
		file, err := os.Open(serverDataPath)
		if err != nil {
			runtime.LogErrorf(ctx, "文件打开失败", err.Error())
			return
		}
		defer func() {
			_ = file.Close()
		}()
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(&serverList)
		if err != nil {
			runtime.LogErrorf(ctx, "解码失败", err.Error())
		} else {
			fmt.Println("解码成功")
			fmt.Printf("ConList:%+v", serverList)
		}
	}

}

func (a *App) GetUserAppDataPath(name string) string {
	return getUserAppDataPath()
}
