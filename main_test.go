package main

import (
	"changeme/dbTools"
	"context"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh"
	gormMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"net"
	"testing"
)

func TestSSHGorm(t *testing.T) {
	client := dbTools.SSH{
		Host: "192.168.0.201",
		User: "root",
		Port: 50122,
		//KeyFile: "~/.ssh/id_rsa",
		Password: "123456",
		Type:     "PASSWORD", // PASSWORD or KEY
	}

	my := dbTools.MySQL{
		Host:     "127.0.0.1",
		User:     "root",
		Password: "123456",
		Port:     3306,
		Database: "mysql",
	}

	var (
		dial *ssh.Client
		err  error
	)
	switch client.Type {
	case "KEY":
		dial, err = client.DialWithKeyFile()
	case "PASSWORD":
		dial, err = client.DialWithPassword()
	default:
		panic("unknown ssh type.")
	}
	if err != nil {
		t.Logf("ssh connect error: %s", err.Error())
		return
	}
	defer dial.Close()

	// 注册ssh代理
	//mysql.RegisterDial("mysql+ssh", (&Dialer{client: dial}).Dial)
	mysql.RegisterDialContext("mysql+ssh", func(_ context.Context, addr string) (net.Conn, error) {
		return (&dbTools.Dialer{Client: dial}).Dial(addr)
	})

	// 填写注册的mysql网络
	dsn := fmt.Sprintf("%s:%s@mysql+ssh(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		my.User, my.Password, my.Host, my.Port, my.Database)
	db, err := gorm.Open(gormMysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
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
	User := "root"
	Password := "123456"
	Host := "192.168.0.201:3306"
	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", User, Password, Host, "mysql")
	db, err := gorm.Open(gormMysql.Open(conn), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
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
