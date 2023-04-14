package main

import (
	"changeme/dbTools"
	"fmt"
	"strings"
	"testing"
)

func TestSSHGorm(t *testing.T) {
	db, dial, err := dbTools.SSHOpenDB(dbTools.SSH{
		Host: "192.168.0.202",
		User: "root",
		Port: 22,
		//KeyFile: "~/.ssh/id_rsa",
		Password: "123456",
		Type:     "PASSWORD", // PASSWORD or KEY
	}, dbTools.MySQL{
		Host:     "192.168.0.201",
		User:     "root",
		Password: "123456",
		Port:     3306,
		Database: "mysql",
	})
	defer func() {
		_ = dial.Close()
	}()
	t.Log(dial.User())
	t.Log(dial.RemoteAddr().String())

	if err != nil {
		return
	}
	if err != nil {
		t.Logf("数据库连接失败,err:%+v", err)
		return
	}
	sqlDB, err := db.DB()
	if err != nil {
		t.Logf("database/sql 获取失败,err:%+v", err)
		return
	}
	//Ping
	if err = sqlDB.Ping(); err != nil {
		t.Logf("sql ping 失败,err:%+v", err)
		return
	}
	//返回数据库统计信息
	t.Logf("实时数据库统计信息:%+v", sqlDB.Stats())

}

func TestGorm(t *testing.T) {
	db, err := dbTools.OpenDB(dbTools.MySQL{
		Host:     "192.168.0.201",
		User:     "root",
		Password: "123456",
		Port:     3306,
		Database: "mysql",
	})
	if err != nil {
		t.Logf("数据库连接失败,err:%+v", err)
		return
	}
	sqlDB, err := db.DB()
	if err != nil {
		t.Logf("database/sql 获取失败,err:%+v", err)
		return
	}
	//Ping
	if err = sqlDB.Ping(); err != nil {
		t.Logf("sql ping 失败,err:%+v", err)
		return
	}
	//返回数据库统计信息
	t.Logf("实时数据库统计信息:%+v", sqlDB.Stats())
	//dataBases := make([]string, 0)
	//if err = db.Raw("show databases;").Scan(&dataBases).Error; err != nil {
	//	t.Logf("sql执行失败,err:%+v", err)
	//	return
	//}
	//fmt.Println(dataBases)
	type TableAction struct {
		TableName    string
		TableComment string
	}
	tables := make([]TableAction, 0)
	if err := db.Debug().Select("table_name,table_comment").Table("information_schema.tables").Where("table_schema = ?", "base_auth").Find(&tables).Error; err != nil {
		t.Logf("数据库查询失败,err:%+v", err)
		return
	}
	fmt.Println(tables)
}

func TestString(t *testing.T) {
	srcString := "This a string"
	destString := "this a string"

	if strings.Compare(srcString, destString) == 0 {
		fmt.Println("Equals")
	} else {
		fmt.Println("Not Equals")
	}
}
