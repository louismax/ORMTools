package main

import (
	"golang.org/x/crypto/ssh"
	"gorm.io/gorm"
)

var serverDataPath string
var ServerConfigMap map[string]ServerConfig
var ServerConnMap map[string]ServerConn

type ServerConfig struct {
	Key            string `json:"key"`
	LocalName      string `json:"local_name"`
	DbType         string `json:"dbType"`
	Host           string `json:"host"`
	Port           string `json:"port"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	HasRecordPwd   bool   `json:"has_record_pwd"`
	HasUseSSH      bool   `json:"hasUseSSH"`
	SshHost        string `json:"ssh_host"`
	SshPort        string `json:"ssh_port"`
	SshUser        string `json:"ssh_user"`
	HasSshKeyfile  bool   `json:"has_ssh_keyfile"`
	SshKeyfile     string `json:"ssh_keyfile"`
	HasSshPass     bool   `json:"has_ssh_pass"`
	SshPassword    string `json:"ssh_password"`
	CreateDateUnix int64  `json:"create_date_unix"`
}

type ServerConn struct {
	DB        *gorm.DB
	SshClient *ssh.Client
	Children  map[string]TreeData `json:"children"`
}

func (a *App) ReturnSuccess(Data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"State": true,
		"Data":  Data,
	}
}

func (a *App) ReturnError(msg string) map[string]interface{} {
	return map[string]interface{}{
		"State":   false,
		"Message": msg,
	}
}

type TreeData struct {
	Key            string     `json:"key"`
	Label          string     `json:"label"`
	Comment        string     `json:"comment"`
	Children       []TreeData `json:"children"`
	ConState       bool       `json:"conState"`
	ObjType        string     `json:"obj_type"`
	HasRecordPwd   bool       `json:"has_record_pwd"`
	CreateDateUnix int64      `json:"create_date_unix"`
	ParentSvrKey   string     `json:"parentSvrKey"`
	ParentDBKey    string     `json:"parentDBKey"`
}

type TableAction struct {
	TableName    string
	TableComment string
}

type TableColumnAction struct {
	ColumnName      string
	OrdinalPosition int64
	ColumnDefault   string
	IsNullable      string
	DataType        string
	ColumnComment   string
}

type StructColumnAction struct {
	SName    string
	SType    string
	STag     string
	SComment string
}
