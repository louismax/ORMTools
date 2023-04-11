package main

import (
	"github.com/mitchellh/go-homedir"
	"log"
	"os"
)

var AppDataPath string

//CheckAppDataPath 检查应用数据路径
func CheckAppDataPath(path string) {
	if ok, _ := PathExists(path + `\AppData`); !ok {
		//创建目录
		createDir(path + `\AppData`)
	}
	AppDataPath = path + `\AppData`
	if ok, _ := PathExists(path + `\log`); !ok {
		//创建目录
		createDir(path + `\log`)
	}
}

//getUserAppDataPath 获取用户应用路径
func getUserAppDataPath() string {
	path, _ := homedir.Dir()
	if ok, _ := PathExists(path + `\ORMTools`); !ok {
		//创建目录
		createDir(path + `\ORMTools`)
	}
	path += `\ORMTools`
	return path
}

//PathExists 检查路径是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//createDir 创建文件夹
func createDir(dirName string) bool {
	err := os.MkdirAll(dirName, 0777)
	if err != nil {
		log.Println(err)
		return false
	}
	err = os.Chmod(dirName, 0777)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
