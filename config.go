package main

import "changeme/constant"

var DefaultConfig = map[interface{}]interface{}{
	constant.ConfigKeyWT: constant.ThemeSystemDefault,
	constant.ConfigKeyHDL: []string{
		"information_schema", "mysql", "performance_schema", "sys",
	},
}

var UserConfig = make(map[interface{}]interface{})
