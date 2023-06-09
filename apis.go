package main

import (
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/louismax/ORMTools/constant"
	"github.com/louismax/ORMTools/dbTools"
	"github.com/louismax/ORMTools/helpers"
	"github.com/mitchellh/go-homedir"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/crypto/ssh"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func (a *App) GetServerConfigList() interface{} {
	tree := make([]TreeData, 0)

	for _, v := range ServerConfigMap {
		tmp := TreeData{
			Key:            v.Key,
			Label:          v.LocalName,
			ObjType:        constant.Connect,
			HasRecordPwd:   v.HasRecordPwd,
			CreateDateUnix: v.CreateDateUnix,
		}
		if _, ok := ServerConnMap[v.Key]; ok {
			tmp.ConState = true
			for _, db := range ServerConnMap[v.Key].Children {
				if len(db.Children) > 0 {
					sort.Slice(db.Children, func(i, j int) bool {
						return strings.Compare(db.Children[i].Label, db.Children[j].Label) < 0
					})
				}
				tmp.Children = append(tmp.Children, db)
			}
			sort.Slice(tmp.Children, func(i, j int) bool {
				return strings.Compare(tmp.Children[i].Label, tmp.Children[j].Label) < 0
			})
		}
		tree = append(tree, tmp)
	}
	sort.Slice(tree, func(i, j int) bool {
		return tree[i].CreateDateUnix < tree[j].CreateDateUnix
	})
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
	req.CreateDateUnix = time.Now().Unix()
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

	_ = svrDatFile.Close() //先关闭文件占用
	svrDatFile, err = os.OpenFile(serverDataPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		runtime.LogErrorf(a.ctx, "文件创建失败,err:%+v", err)
		return a.ReturnError("文件创建失败")
	}
	encoder := gob.NewEncoder(svrDatFile)
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
	req.CreateDateUnix = ServerConfigMap[req.Key].CreateDateUnix
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

	_ = svrDatFile.Close() //先关闭文件占用
	svrDatFile, err = os.OpenFile(serverDataPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		runtime.LogErrorf(a.ctx, "文件创建失败,err:%+v", err)
		return a.ReturnError("文件创建失败")
	}
	encoder := gob.NewEncoder(svrDatFile)
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
	var err error
	_ = svrDatFile.Close() //先关闭文件占用
	svrDatFile, err = os.OpenFile(serverDataPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		runtime.LogErrorf(a.ctx, "文件创建失败,err:%+v", err)
		return a.ReturnError("文件创建失败")
	}
	encoder := gob.NewEncoder(svrDatFile)
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
	//fmt.Printf("%+v", dc)

	if ServerConfigMap[req.Key].HasUseSSH {
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
		fmt.Printf("%+v", sc)

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
	hideDB := make(map[string]interface{})
	for _, v := range UserConfig[constant.ConfigKeyHideDBList].([]interface{}) {
		hideDB[v.(string)] = 1
	}

	for _, v := range dataBases {
		if _, ok := hideDB[v]; ok {
			continue
		}
		dbTree = append(dbTree, TreeData{
			Key:          uuid.New().String(),
			Label:        v,
			Children:     nil,
			ObjType:      "db",
			ParentSvrKey: req.Key,
		})
	}
	sort.Slice(dbTree, func(i, j int) bool {
		return strings.Compare(dbTree[i].Label, dbTree[j].Label) < 0
	})

	svrCon := ServerConnMap[req.Key]
	dbTmp := make(map[string]TreeData)
	for _, v := range dbTree {
		dbTmp[v.Label] = v
	}
	svrCon.Children = dbTmp

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

func (a *App) QueryTableList(key, dbName string) interface{} {
	if _, ok := ServerConnMap[key]; !ok {
		return a.ReturnError("连接不存在或已关闭，请重新连接后重试！")
	}
	tables := make([]TableAction, 0)
	if err := ServerConnMap[key].DB.Debug().Select("table_name,table_comment").Table("information_schema.tables").Where("table_schema = ?", dbName).Find(&tables).Error; err != nil {
		runtime.LogErrorf(a.ctx, "数据库查询失败,err:%+v", err)
		return a.ReturnError("数据库查询失败,ERR:" + err.Error())
	}
	dbTree := make([]TreeData, 0)
	for _, v := range tables {
		dbTree = append(dbTree, TreeData{
			Key:          uuid.New().String(),
			Label:        v.TableName,
			Comment:      v.TableComment,
			Children:     nil,
			ObjType:      "table",
			ParentSvrKey: key,
			ParentDBKey:  ServerConnMap[key].Children[dbName].Label,
		})
	}
	sort.Slice(dbTree, func(i, j int) bool {
		return strings.Compare(dbTree[i].Label, dbTree[j].Label) < 0
	})
	svrCon := ServerConnMap[key]
	dbInfo := svrCon.Children[dbName]
	dbInfo.Children = dbTree
	svrCon.Children[dbName] = dbInfo
	ServerConnMap[key] = svrCon

	return a.ReturnSuccess(dbTree)
}

func (a *App) RefreshDBConnect(key string) interface{} {
	dataBases := make([]string, 0)
	if err := ServerConnMap[key].DB.Debug().Raw("show databases;").Scan(&dataBases).Error; err != nil {
		runtime.LogErrorf(a.ctx, "数据库查询失败,err:%+v", err)
		return a.ReturnError("数据库查询失败,ERR:" + err.Error())
	}
	dbTree := make([]TreeData, 0)
	for _, v := range dataBases {
		dbTree = append(dbTree, TreeData{
			Key:          uuid.New().String(),
			Label:        v,
			Children:     nil,
			ObjType:      "db",
			ParentSvrKey: key,
		})
	}
	sort.Slice(dbTree, func(i, j int) bool {
		return strings.Compare(dbTree[i].Label, dbTree[j].Label) < 0
	})

	svrCon := ServerConnMap[key]
	dbTmp := make(map[string]TreeData)
	for _, v := range dbTree {
		dbTmp[v.Label] = v
	}
	svrCon.Children = dbTmp
	ServerConnMap[key] = svrCon
	return a.ReturnSuccess(dbTree)
}

func (a *App) QueryTableFieldList(key, dbKey, tableName, tbComment string) interface{} {
	if _, ok := ServerConnMap[key]; !ok {
		return a.ReturnError("连接不存在或已关闭，请重新连接后重试！")
	}
	if _, ok := ServerConnMap[key].Children[dbKey]; !ok {
		return a.ReturnError("数据库不存在或已关闭，请重新连接后重试！")
	}
	tableColumn := make([]TableColumnAction, 0)
	if err := ServerConnMap[key].DB.Debug().
		Select("column_name,ordinal_position,column_default,is_nullable,data_type,column_comment").
		Table("information_schema.COLUMNS").
		Where("table_schema = ? AND  table_name = ?", ServerConnMap[key].Children[dbKey].Label, tableName).
		Find(&tableColumn).Error; err != nil {
		runtime.LogErrorf(a.ctx, "表结构查询失败,err:%+v", err)
		return a.ReturnError("表结构查询失败,ERR:" + err.Error())
	}

	maxNameLen := 0
	maxTypeLen := 0
	maxTagLen := 0
	sm := make([]StructColumnAction, 0)
	for _, v := range tableColumn {
		tag := ""
		if UserConfig[constant.ConfigKeyHasJsonTag].(bool) {
			tag += fmt.Sprintf("json:%q", v.ColumnName)
		}
		if UserConfig[constant.ConfigKeyHasGormColumnTag].(bool) {
			if UserConfig[constant.ConfigKeyHasJsonTag].(bool) {
				tag += " "
			}
			tag += fmt.Sprintf("gorm:%q", fmt.Sprintf("column:%s", v.ColumnName))
		}
		tmp := StructColumnAction{
			SName:    camelString(v.ColumnName),
			SType:    DBFieldTypeToStructFieldType(v.DataType),
			SComment: "// " + v.ColumnComment,
		}
		if tag != "" {
			tmp.STag = fmt.Sprintf("%#q", tag)
		}
		//ColumnName      string `json:"columnName" gorm:"column:beast_id"`

		if len(tmp.SName) > maxNameLen {
			maxNameLen = len(tmp.SName)
		}
		if len(tmp.SType) > maxTypeLen {
			maxTypeLen = len(tmp.SType)
		}
		if len(tmp.STag) > maxTagLen {
			maxTagLen = len(tmp.STag)
		}
		sm = append(sm, tmp)
	}

	str := fmt.Sprintf("//%s %s \n", camelString(tableName), tbComment)
	str += fmt.Sprintf("type %s struct{\n", camelString(tableName))
	for _, v := range sm {
		if v.STag == "" {
			str += fmt.Sprintf("\t%-"+strconv.Itoa(maxNameLen)+"s %-"+strconv.Itoa(maxTypeLen)+"s %s\n", v.SName, v.SType, v.SComment)
		} else {
			str += fmt.Sprintf("\t%-"+strconv.Itoa(maxNameLen)+"s %-"+strconv.Itoa(maxTypeLen)+"s %-"+strconv.Itoa(maxTagLen)+"s %s\n", v.SName, v.SType, v.STag, v.SComment)
		}
	}
	str += "}\n"

	if UserConfig[constant.ConfigKeyHasRewriteTableName].(bool) {
		str += fmt.Sprintf("\n// TableName %s\n", camelString(tableName))
		str += fmt.Sprintf("func (%s) TableName() string {\n", camelString(tableName))
		str += fmt.Sprintf("\treturn %q\n", tableName)
		str += "}\n"
	}

	return a.ReturnSuccess(str)
}

func (a *App) GetUserConfig() interface{} {
	//fmt.Printf("%+v", UserConfig)
	return a.ReturnSuccess(UserConfig)
}

func (a *App) EditUserConfigItem(val map[string]interface{}) interface{} {
	var err error
	UserConfig = val
	_ = userCfgFile.Close() //先关闭文件占用
	userCfgFile, err = os.OpenFile(userCfgPath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		runtime.LogErrorf(a.ctx, "文件打开失败,err:%+v", err)
		return a.ReturnError("文件打开失败")
	}
	d, err := yaml.Marshal(&UserConfig)
	if err != nil {
		runtime.LogErrorf(a.ctx, "yaml序列化处理失败,err:%+v", err)
		return a.ReturnError("yaml序列化处理失败")
	}
	_, err = userCfgFile.Write(d)
	if err != nil {
		runtime.LogErrorf(a.ctx, "用户配置文件写入失败,err:%+v", err)
		return a.ReturnError("用户配置文件写入失败")
	}
	err = userCfgFile.Sync()
	if err != nil {
		runtime.LogErrorf(a.ctx, "用户配置文件写入同步失败,err:%+v", err)
		return a.ReturnError("用户配置文件写入同步失败")
	}

	return a.ReturnSuccess(UserConfig)
}

func (a *App) ExportServerConfigList() interface{} {
	b, err := json.Marshal(ServerConfigMap)
	if err != nil {
		runtime.LogErrorf(a.ctx, "json序列化失败,err:%+v", err)
		return a.ReturnError("json序列化失败")
	}
	sEnc := base64.StdEncoding.EncodeToString(b)

	es, err := helpers.EncryptByAes([]byte(sEnc))
	if err != nil {
		runtime.LogErrorf(a.ctx, "AES加密失败,err:%+v", err)
		return a.ReturnError("AES加密失败")
	}
	return a.ReturnSuccess(es)
}

func (a *App) ImportServerConfigList(s string) interface{} {
	es, err := helpers.DecryptByAes(s)
	if err != nil {
		runtime.LogErrorf(a.ctx, "AES解密失败,err:%+v", err)
		return a.ReturnError("文本解析失败，请检查导入文本是否有效！")
	}
	sDec, err := base64.StdEncoding.DecodeString(string(es))
	if err != nil {
		runtime.LogErrorf(a.ctx, "Base64解析失败,err:%+v", err)
		return a.ReturnError("文本解析失败，请检查导入文本是否有效！")
	}
	m := make(map[string]ServerConfig)
	err = json.Unmarshal(sDec, &m)
	if err != nil {
		runtime.LogErrorf(a.ctx, "Json解析失败,err:%+v", err)
		return a.ReturnError("文本解析失败，请检查导入文本是否有效！")
	}
	cu := time.Now().Unix()
	for _, v := range m {
		val := eachSvrLocalName(&v)
		val.CreateDateUnix = cu
		ServerConfigMap[val.Key] = *val
		cu++
	}

	_ = svrDatFile.Close() //先关闭文件占用
	svrDatFile, err = os.OpenFile(serverDataPath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		runtime.LogErrorf(a.ctx, "文件创建失败,err:%+v", err)
		return a.ReturnError("文件创建失败")
	}
	encoder := gob.NewEncoder(svrDatFile)
	err = encoder.Encode(ServerConfigMap)
	if err != nil {
		runtime.LogErrorf(a.ctx, "gob编码错误,err:%+v", err)
		return a.ReturnError("gob编码错误")
	}

	return a.ReturnSuccess("成功")
}

func eachSvrLocalName(val *ServerConfig) *ServerConfig {
	if _, ok := ServerConfigMap[val.Key]; ok {
		val.Key = uuid.New().String()
	}
	for _, v := range ServerConfigMap {
		if v.LocalName == val.LocalName {
			val.LocalName += "_副本"
			eachSvrLocalName(val)
		}
	}
	return val
}

func (a *App) SaveServerConfigFile(s string) interface{} {
	up, err := homedir.Dir()
	if err != nil {
		runtime.LogErrorf(a.ctx, "获取用户路径失败,err:%+v", err)
		return a.ReturnError("获取用户路径失败")
	}

	path, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		DefaultDirectory: up + "\\Desktop",
		DefaultFilename:  "ORMTools服务配置",
		Title:            "保存配置文件",
	})
	if err != nil {
		runtime.LogErrorf(a.ctx, "用户选择路径失败,err:%+v", err)
		return a.ReturnError("用户选择路径失败")
	}
	if path != "" {
		f, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
		if err != nil {
			runtime.LogErrorf(a.ctx, "文件创建失败,err:%+v", err)
			return a.ReturnError("文件创建失败")
		}
		defer func() {
			_ = f.Close()
		}()
		_, err = f.Write([]byte(s))
		if err != nil {
			runtime.LogErrorf(a.ctx, "文件写入失败,err:%+v", err)
			return a.ReturnError("文件写入失败")
		}
		return a.ReturnSuccess(fmt.Sprintf("已保存在%s", path))
	}
	return a.ReturnSuccess("cancel")
}

func (a *App) ReadServerConfigFile() interface{} {
	path, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{})
	if err != nil {
		runtime.LogErrorf(a.ctx, "用户选择文件失败,err:%+v", err)
		return a.ReturnError("用户选择文件失败")
	}
	if path != "" {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			runtime.LogErrorf(a.ctx, "文件读取失败,err:%+v", err)
			return a.ReturnError("文件读取失败")
		}
		return a.ReturnSuccess(string(data))
	}
	return a.ReturnSuccess("cancel")
}
