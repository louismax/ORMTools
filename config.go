package main

import "changeme/constant"

var DefaultConfig = map[interface{}]interface{}{
	"WindowTheme": constant.ThemeLight,
}
var UserConfig = make(map[interface{}]interface{})
