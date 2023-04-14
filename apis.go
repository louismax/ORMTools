package main

import (
	"changeme/constant"
	"changeme/dbTools"
	"changeme/helpers"
	"encoding/gob"
	"fmt"
	"github.com/google/uuid"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/crypto/ssh"
	"gorm.io/gorm"
	"os"
)

func (a *App) GetUserAppDataPath(name string) string {
	return getUserAppDataPath()
}

func (a *App) GetServerConfigList() interface{} {
	tree := make([]TreeData, 0)

	for _, v := range ServerConfigMap {
		tmp := TreeData{
			Key:          v.Key,
			Label:        v.LocalName,
			Children:     nil,
			ConState:     false,
			ObjType:      constant.Connect,
			HasRecordPwd: v.HasRecordPwd,
		}
		if _, ok := ServerConnMap[v.Key]; ok {
			tmp.ConState = true
			tmp.Children = ServerConnMap[v.Key].Children
		}
		tree = append(tree, tmp)
	}

	return a.ReturnSuccess(tree)
}

func (a *App) GetServerConfig(key string) interface{} {
	if _, ok := ServerConfigMap[key]; !ok {
		return a.ReturnError("配置不存在或已删除，请重启应用后重试!")
	}
	res := ServerConfigMap[key]
	planPwd, err := helpers.DecryptByAes(res.Password)
	if err != nil {
		runtime.LogErrorf(a.ctx, "敏感数据解密失败:%+v", err)
		return a.ReturnError("敏感数据解密失败")
	}
	res.Password = string(planPwd)
	if res.HasUseSSH && res.HasSshPass {
		planSSHPwd, err := helpers.DecryptByAes(res.SshPassword)
		if err != nil {
			runtime.LogErrorf(a.ctx, "敏感数据解密失败,err:%+v", err)
			return a.ReturnError("敏感数据解密失败[2]")
		}
		res.SshPassword = string(planSSHPwd)
	}
	return a.ReturnSuccess(res)
}

func (a *App) AddServerConfig(req ServerConfig) interface{} {
	if req.LocalName == "" {
		req.LocalName = fmt.Sprintf("%s_%s:%s", req.DbType, req.Host, req.Port)
	}
	//检查是否存在同名配置
	for _, v := range ServerConfigMap {
		if v.LocalName == req.LocalName {
			return a.ReturnError("已存在同名配置,请修改或完善名称!")
		}
	}
	req.Key = uuid.New().String()
	aesPwd, err := helpers.EncryptByAes([]byte(req.Password))
	if err != nil {
		runtime.LogErrorf(a.ctx, "AES加密失败,err:%+v", err)
		return a.ReturnError("AES加密失败[1]")
	}
	req.Password = aesPwd
	if req.HasUseSSH && req.HasSshPass {
		aesSSHPwd, err := helpers.EncryptByAes([]byte(req.SshPassword))
		if err != nil {
			runtime.LogErrorf(a.ctx, "AES加密失败,err:%+v", err)
			return a.ReturnError("AES加密失败[2]")
		}
		req.SshPassword = aesSSHPwd
	}
	ServerConfigMap[req.Key] = req
	file, err := os.OpenFile(serverDataPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		runtime.LogErrorf(a.ctx, "文件创建失败,err:%+v", err)
		return a.ReturnError("文件创建失败")
	}
	defer func() {
		_ = file.Close()
	}()
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(ServerConfigMap)
	if err != nil {
		runtime.LogErrorf(a.ctx, "gob编码错误,err:%+v", err)
		return a.ReturnError("gob编码错误")
	}
	return a.ReturnSuccess("成功")
}

func (a *App) EditServerConfig(req ServerConfig) interface{} {
	if _, ok := ServerConfigMap[req.Key]; !ok {
		return a.ReturnError("配置不存在或已删除，请重启应用后重试!")
	}
	aesPwd, err := helpers.EncryptByAes([]byte(req.Password))
	if err != nil {
		runtime.LogErrorf(a.ctx, "AES加密失败,err:%+v", err)
		return a.ReturnError("AES加密失败[1]")
	}
	req.Password = aesPwd
	if req.HasUseSSH && req.HasSshPass {
		aesSSHPwd, err := helpers.EncryptByAes([]byte(req.SshPassword))
		if err != nil {
			runtime.LogErrorf(a.ctx, "AES加密失败,err:%+v", err)
			return a.ReturnError("AES加密失败[2]")
		}
		req.SshPassword = aesSSHPwd
	}
	ServerConfigMap[req.Key] = req
	file, err := os.OpenFile(serverDataPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		runtime.LogErrorf(a.ctx, "文件创建失败,err:%+v", err)
		return a.ReturnError("文件创建失败")
	}
	defer func() {
		_ = file.Close()
	}()
	encoder := gob.NewEncoder(file)
	fmt.Printf("%+v", ServerConfigMap)
	err = encoder.Encode(ServerConfigMap)
	if err != nil {
		runtime.LogErrorf(a.ctx, "gob编码错误,err:%+v", err)
		return a.ReturnError("gob编码错误")
	}
	return a.ReturnSuccess("成功")
}

func (a *App) DeleteServerConfig(key string) interface{} {
	if _, ok := ServerConfigMap[key]; !ok {
		return a.ReturnError("配置不存在或已删除，请重启应用后重试!")
	}
	delete(ServerConfigMap, key)

	file, err := os.OpenFile(serverDataPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		runtime.LogErrorf(a.ctx, "文件创建失败,err:%+v", err)
		return a.ReturnError("文件创建失败")
	}
	defer func() {
		_ = file.Close()
	}()
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(ServerConfigMap)
	if err != nil {
		runtime.LogErrorf(a.ctx, "gob编码错误,err:%+v", err)
		return a.ReturnError("gob编码错误")
	}
	return a.ReturnSuccess("成功")
}

func (a *App) TestDBConnect(req ServerConfig) interface{} {
	var (
		db   *gorm.DB
		dial *ssh.Client
		err  error
	)
	if req.HasUseSSH {
		sc := dbTools.SSH{
			Host: req.SshHost,
			User: req.SshUser,
			Port: int(InterfaceToInt64(req.SshPort)),
		}
		if req.HasSshPass {
			sc.Type = constant.Password
			sc.Password = req.SshPassword
		}
		if req.HasSshKeyfile {
			sc.Type = constant.Key
			sc.Password = req.SshKeyfile
		}
		db, dial, err = dbTools.SSHOpenDB(sc, dbTools.MySQL{
			Host:     req.Host,
			User:     req.Username,
			Password: req.Password,
			Port:     int(InterfaceToInt64(req.Port)),
			Database: "mysql",
		})
		if err != nil {
			runtime.LogErrorf(a.ctx, "数据库连接失败,err:%+v", err)
			return a.ReturnError("数据库连接失败,ERR:" + err.Error())
		}
		defer func() {
			_ = dial.Close()
		}()
	} else {
		db, err = dbTools.OpenDB(dbTools.MySQL{
			Host:     req.Host,
			User:     req.Username,
			Password: req.Password,
			Port:     int(InterfaceToInt64(req.Port)),
			Database: "mysql",
		})
		if err != nil {
			runtime.LogErrorf(a.ctx, "数据库连接失败,err:%+v", err)
			return a.ReturnError("数据库连接失败,ERR:" + err.Error())
		}

	}
	sqlDB, err := db.DB()
	if err != nil {
		runtime.LogErrorf(a.ctx, "sql.DB实例获取失败,err:%+v", err)
		return a.ReturnError("sql.DB实例获取失败,ERR:" + err.Error())
	}
	defer func() {
		_ = sqlDB.Close()
	}()
	//Ping
	if err = sqlDB.Ping(); err != nil {
		runtime.LogErrorf(a.ctx, "Ping失败,err:%+v", err)
		return a.ReturnError("Ping失败,ERR:" + err.Error())
	}
	return a.ReturnSuccess(sqlDB.Stats())
}

func (a *App) OpenDBConnect(req ServerConfig) interface{} {
	if _, ok := ServerConfigMap[req.Key]; !ok {
		return a.ReturnError("配置不存在或已删除，请重启应用后重试!")
	}
	dc := dbTools.MySQL{
		Host:     ServerConfigMap[req.Key].Host,
		User:     ServerConfigMap[req.Key].Username,
		Port:     int(InterfaceToInt64(ServerConfigMap[req.Key].Port)),
		Database: "mysql",
	}
	if !ServerConfigMap[req.Key].HasRecordPwd {
		dc.Password = req.Password
	} else {
		decPwd, err := helpers.DecryptByAes(ServerConfigMap[req.Key].Password)
		if err != nil {
			runtime.LogErrorf(a.ctx, "AES解密失败,err:%+v", err)
			return a.ReturnError("AES解密失败[1]")
		}
		dc.Password = string(decPwd)
	}
	if req.HasUseSSH {
		sc := dbTools.SSH{
			Host: ServerConfigMap[req.Key].SshHost,
			User: ServerConfigMap[req.Key].SshUser,
			Port: int(InterfaceToInt64(ServerConfigMap[req.Key].SshPort)),
		}
		if ServerConfigMap[req.Key].HasSshPass {
			sc.Type = constant.Password
			decPwd, err := helpers.DecryptByAes(ServerConfigMap[req.Key].SshPassword)
			if err != nil {
				runtime.LogErrorf(a.ctx, "AES解密失败,err:%+v", err)
				return a.ReturnError("AES解密失败[1]")
			}
			sc.Password = string(decPwd)
		}
		if ServerConfigMap[req.Key].HasSshKeyfile {
			sc.Type = constant.Key
			sc.Password = ServerConfigMap[req.Key].SshKeyfile
		}

		db, dial, err := dbTools.SSHOpenDB(sc, dc)
		if err != nil {
			runtime.LogErrorf(a.ctx, "数据库连接失败,err:%+v", err)
			return a.ReturnError("数据库连接失败,ERR:" + err.Error())
		}
		ServerConnMap[req.Key] = ServerConn{
			DB:        db,
			SshClient: dial,
		}
	} else {
		db, err := dbTools.OpenDB(dc)
		if err != nil {
			runtime.LogErrorf(a.ctx, "数据库连接失败,err:%+v", err)
			return a.ReturnError("数据库连接失败,ERR:" + err.Error())
		}
		ServerConnMap[req.Key] = ServerConn{
			DB:        db,
			SshClient: nil,
		}
	}
	dataBases := make([]string, 0)
	if err := ServerConnMap[req.Key].DB.Debug().Raw("show databases;").Scan(&dataBases).Error; err != nil {
		runtime.LogErrorf(a.ctx, "数据库查询失败,err:%+v", err)
		return a.ReturnError("数据库查询失败,ERR:" + err.Error())
	}
	dbTree := make([]TreeData, 0)
	for _, v := range dataBases {
		dbTree = append(dbTree, TreeData{
			Key:      v,
			Label:    v,
			Children: nil,
			ObjType:  "db",
		})
	}
	svrCon := ServerConnMap[req.Key]
	svrCon.Children = dbTree
	ServerConnMap[req.Key] = svrCon
	return a.ReturnSuccess(dbTree)
}

func (a *App) CloseDBConnect(key string) interface{} {
	if _, ok := ServerConfigMap[key]; !ok {
		return a.ReturnError("配置不存在或已删除，请重启应用后重试!")
	}

	sqlDB, err := ServerConnMap[key].DB.DB()
	if err != nil {
		runtime.LogErrorf(a.ctx, "获取DB实例失败,err:%+v", err)
		return a.ReturnError("获取DB实例失败")
	}
	_ = sqlDB.Close()
	if ServerConnMap[key].SshClient != nil {
		_ = ServerConnMap[key].SshClient.Close()
	}

	delete(ServerConnMap, key)

	return a.ReturnSuccess("操作成功")
}