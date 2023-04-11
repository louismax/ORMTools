package main

import (
	"fmt"
	"os"
	"testing"
)

func TestName(t *testing.T) {
	//err := os.Chdir(`C:\Users\alone\ORMTools`)
	//if err != nil {
	//	t.Log(err)
	//	return
	//}
	//path, _ := os.Getwd()
	//t.Log(path)
	//// 获取文件的绝对路径
	//path2, err := filepath.Abs(`./a.dat`)
	//if err != nil {
	//	fmt.Println("file path error: " + err.Error())
	//	return
	//}
	path2 := fmt.Sprintf(`%s\AppData\server.dat`, getUserAppDataPath())
	t.Log(path2)
	file, err := os.OpenFile(path2, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("创建文件失败:%+v\n", err)
	}
	defer file.Close()
	_, err = file.WriteString("123")
	if err != nil {
		fmt.Println(err)
	}

	//path, err := filepath.Abs(`./`)
	//if err != nil {
	//	fmt.Println("file path error: " + err.Error())
	//	return
	//}
	//t.Log(path)
	//path2, err := filepath.Rel(`C:\Users\alone\ORMTools\con.bat`, path)
	//if err != nil {
	//	fmt.Println("file path error: " + err.Error())
	//	return
	//}
	//t.Log(path2)
}

//func createFileWithDir(path string, name string, content string) {
//	err := os.MkdirAll(path, os.ModePerm)
//	if err != nil {
//		fmt.Printf("创建文件夹失败:%+v", err)
//	}
//	file, err := os.OpenFile(path+"\\"+name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
//	if err != nil {
//		fmt.Printf("创建文件失败:%+v\n", err)
//	}
//	defer file.Close()
//	_, err = file.WriteString(content)
//	if err != nil {
//		fmt.Println(err)
//	}
//}
