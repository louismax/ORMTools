package main

import (
	"encoding/json"
	"github.com/mitchellh/go-homedir"
	"log"
	"os"
	"strconv"
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

// InterfaceToInt64 interface{} 转 int64
func InterfaceToInt64(value interface{}) int64 {
	// interface 转 int64
	var key int64
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		keys := strconv.FormatFloat(ft, 'f', -1, 64)
		key, _ = strconv.ParseInt(keys, 10, 64)
	case float32:
		ft := value.(float32)
		keys := strconv.FormatFloat(float64(ft), 'f', -1, 64)
		key, _ = strconv.ParseInt(keys, 10, 64)
	case int:
		it := value.(int)
		keys := strconv.Itoa(it)
		key, _ = strconv.ParseInt(keys, 10, 64)
	case uint:
		it := value.(uint)
		keys := strconv.Itoa(int(it))
		key, _ = strconv.ParseInt(keys, 10, 64)
	case int8:
		it := value.(int8)
		keys := strconv.Itoa(int(it))
		key, _ = strconv.ParseInt(keys, 10, 64)
	case uint8:
		it := value.(uint8)
		keys := strconv.Itoa(int(it))
		key, _ = strconv.ParseInt(keys, 10, 64)
	case int16:
		it := value.(int16)
		keys := strconv.Itoa(int(it))
		key, _ = strconv.ParseInt(keys, 10, 64)
	case uint16:
		it := value.(uint16)
		keys := strconv.Itoa(int(it))
		key, _ = strconv.ParseInt(keys, 10, 64)
	case int32:
		it := value.(int32)
		keys := strconv.Itoa(int(it))
		key, _ = strconv.ParseInt(keys, 10, 64)
	case uint32:
		it := value.(uint32)
		keys := strconv.Itoa(int(it))
		key, _ = strconv.ParseInt(keys, 10, 64)
	case int64:
		it := value.(int64)
		keys := strconv.FormatInt(it, 10)
		key, _ = strconv.ParseInt(keys, 10, 64)
	case uint64:
		it := value.(uint64)
		keys := strconv.FormatUint(it, 10)
		key, _ = strconv.ParseInt(keys, 10, 64)
	case string:
		keys := value.(string)
		key, _ = strconv.ParseInt(keys, 10, 64)
	case []byte:
		keys := string(value.([]byte))
		key, _ = strconv.ParseInt(keys, 10, 64)
	default:
		newValue, _ := json.Marshal(value)
		keys := string(newValue)
		key, _ = strconv.ParseInt(keys, 10, 64)
	}

	return key
}
