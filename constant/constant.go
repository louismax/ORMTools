package constant

const (
	Connect  string = "connect"
	DB       string = "db"
	Table    string = "table"
	Password string = "PASSWORD"
	Key      string = "KEY"
)

const (
	ConfigKeyWT                     string = "WindowTheme"
	ConfigKeyHideDBList             string = "HideDBList"
	ConfigKeyHideTableList          string = "HideTableList"
	ConfigKeyHideTableColumnList    string = "HideTableColumnList"
	ConfigKeyHasTableComment        string = "HasTableComment"
	ConfigKeyHasTableRI             string = "ConfigKeyHasTableRI"
	ConfigKeyMySqlToStructFieldType string = "MySqlToStructFieldType"
)

const (
	ThemeSystemDefault string = "SystemDefault"
	ThemeLight         string = "Light"
	ThemeDark          string = "Dark"
)

//int8: -128 ~ 127
//int16: -32768 ~ 32767
//int32: -2147483648 ~ 2147483647
//int64: -9223372036854775808 ~ 9223372036854775807
//uint8: 0 ~ 255
//uint16: 0 ~ 65535
//uint32: 0 ~ 4294967295
//uint64: 0 ~ 18446744073709551615
const (
	MySqlTinyInt   string = "tinyint"   //-128~127
	MySqlSmallInt  string = "smallint"  //-32768~32767
	MySqlMediumInt string = "mediumint" //-8388608~8388607
	MySqlInteger   string = "integer"   //-2147483648~2147483647
	MySqlInt       string = "int"       //-2147483648~2147483647
	MySqlBigInt    string = "bigint"    //-9223372036854775808~9223372036854775807

)

const (
	MySqlChar     string = "char"
	MySqlVarChar  string = "varchar"
	MySqlText     string = "text"
	MySqlDate     string = "date"
	MySqlDateTime string = "datetime"
	MySqlDouble   string = "float64"
)

