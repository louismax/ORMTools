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
	constant.ConfigKeyHideTableList:       []string{},
	constant.ConfigKeyHideTableColumnList: []string{},
	constant.ConfigKeyMySqlToStructFieldType: map[string]string{
		constant.MySqlBigInt:   "int64",
		constant.MySqlTinyInt:  "uint",
		constant.MySqlInt:      "int",
		constant.MySqlChar:     "string",
		constant.MySqlVarChar:  "string",
		constant.MySqlText:     "string",
		constant.MySqlDate:     "time.Time",
		constant.MySqlDateTime: "time.Time",
		constant.MySqlDouble:   "float64",
	},
}

var UserConfig = make(map[interface{}]interface{})
