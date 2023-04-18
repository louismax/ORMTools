package main

import (
	"changeme/constant"
)

var DefaultConfig = map[interface{}]interface{}{
	constant.ConfigKeyWT:              constant.ThemeSystemDefault,
	constant.ConfigKeyHasTableComment: true,
	constant.ConfigKeyHasTableRI:      true,
	constant.ConfigKeyHideDBList: []string{
		"information_schema", "mysql", "performance_schema", "sys",
	},
	constant.ConfigKeyHideTableList:          []string{},
	constant.ConfigKeyHideTableColumnList:    []string{},
	constant.ConfigKeyMySqlToStructFieldType: map[string]string{},
}

var UserConfig = make(map[interface{}]interface{})
