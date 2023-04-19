package main

import (
	"changeme/constant"
)

var DefaultConfig = map[string]interface{}{
	constant.ConfigKeyWT:                  constant.ThemeSystemDefault, //主题
	constant.ConfigKeyHasTableComment:     true,                        //是否显示表备注
	constant.ConfigKeyHasRewriteTableName: true,                        //是否重写表名
	constant.ConfigKeyHasJsonTag:          true,
	constant.ConfigKeyHasGormColumnTag:    false,
	constant.ConfigKeyHideDBList: []string{
		"information_schema", "mysql", "performance_schema", "sys",
	},
	constant.ConfigKeyHideTableList:          []string{},
	constant.ConfigKeyHideTableColumnList:    []string{},
	constant.ConfigKeyMySqlToStructFieldType: map[string]string{},
}

var UserConfig = make(map[string]interface{})
