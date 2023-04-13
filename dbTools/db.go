package dbTools

import (
	"context"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh"
	gormMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io/ioutil"
	"net"
)

type Dialer struct {
	Client *ssh.Client
}

type SSH struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Port     int    `json:"port"`
	Type     string `json:"type"`
	Password string `json:"password"`
	KeyFile  string `json:"key"`
}

type MySQL struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func (v *Dialer) Dial(address string) (net.Conn, error) {
	return v.Client.Dial("tcp", address)
}

func (s *SSH) DialWithPassword() (*ssh.Client, error) {
	address := fmt.Sprintf("%s:%d", s.Host, s.Port)
	config := &ssh.ClientConfig{
		User: s.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(s.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	return ssh.Dial("tcp", address, config)
}

func (s *SSH) DialWithKeyFile() (*ssh.Client, error) {
	address := fmt.Sprintf("%s:%d", s.Host, s.Port)
	config := &ssh.ClientConfig{
		User:            s.User,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	if k, err := ioutil.ReadFile(s.KeyFile); err != nil {
		return nil, err
	} else {
		signer, err := ssh.ParsePrivateKey(k)
		if err != nil {
			return nil, err
		}
		config.Auth = []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		}
	}
	return ssh.Dial("tcp", address, config)
}

func SSHOpenDB(sc SSH, dc MySQL) (*gorm.DB, *ssh.Client, error) {
	var (
		dial *ssh.Client
		err  error
	)
	switch sc.Type {
	case "KEY":
		dial, err = sc.DialWithKeyFile()
	case "PASSWORD":
		dial, err = sc.DialWithPassword()
	default:
		panic("unknown ssh type.")
	}
	if err != nil {
		return nil, nil, err
	}
	// 注册ssh代理
	//mysql.RegisterDial("mysql+ssh", (&Dialer{client: dial}).Dial)
	netKey := fmt.Sprintf("ssh_%s_%d", sc.Host, sc.Port)
	//mysql.RegisterDialContext("mysql+ssh", func(_ context.Context, addr string) (net.Conn, error) {
	mysql.RegisterDialContext(netKey, func(_ context.Context, addr string) (net.Conn, error) {
		return (&Dialer{Client: dial}).Dial(addr)
	})
	// 填写注册的mysql网络
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dc.User, dc.Password, netKey, dc.Host, dc.Port, dc.Database)
	db, err := gorm.Open(gormMysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		return nil, nil, err
	}
	return db, dial, nil
}

func OpenDB(dc MySQL) (*gorm.DB, error) {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dc.User, dc.Password, dc.Host, dc.Port, dc.Database)
	db, err := gorm.Open(gormMysql.Open(conn), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		return nil, err
	}
	return db, nil
}
