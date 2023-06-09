package main

import (
	"changeme/dbTools"
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

}

func TestCamelString(t *testing.T) {
	str := "basic_user_info"
	t.Log(camelString(str))
}
