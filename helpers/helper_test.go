package helpers

import (
	"fmt"
	"testing"
)

func TestAES(t *testing.T) {
	//加密
	str, _ := EncryptByAes([]byte("ORMTools"))
	//解密
	str1, _ := DecryptByAes(str)
	//打印
	fmt.Printf(" 加密：%v\n 解密：%s\n ",
		str, str1,
	)
}
