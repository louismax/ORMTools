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
//float32 :-3.403e38 ~ 3.403e38
//float64 :-1.798e308 ~ 1.798e308
const (
	MySqlTinyInt   string = "tinyint"   //-128~127
	MySqlSmallInt  string = "smallint"  //-32768~32767
	MySqlMediumInt string = "mediumint" //-8388608~8388607
	MySqlInteger   string = "integer"   //-2147483648~2147483647
	MySqlInt       string = "int"       //-2147483648~2147483647
	MySqlBigInt    string = "bigint"    //-9223372036854775808~9223372036854775807
	MysqlFloat     string = "float"     //-3.402823466E+38～-1.175494351E-38
	MysqlDouble    string = "double"    //-1.7976931348623157E+308～-2.2250738585072014E-308
	MySqlDecimal   string = "decimal"
)

const (
	MySqlChar     string = "char"     //0-255
	MySqlVarChar  string = "varchar"  //0-65535
	MySqlText     string = "text"     //0-65535
	MySqlLongText string = "longtext" //0-4294967295
	MySqlBlob     string = "blob"     //二进制 0-65535
	MySqlLongBlob string = "longblob" //二进制 0-4294967295
)

const (
	MySqlDate      string = "date"      //1000-01-01/9999-12-31 YYYY-MM-DD
	MySqlTime      string = "time"      //’-838:59:59’/838:59:59’ HH:MM:SS
	MySqlYear      string = "year"      //1901/2155 YYYY
	MySqlDateTime  string = "datetime"  //1000-01-01 00:00:00/9999-12-31 23:59:59  YYYY-MM-DD HH:MM:SS
	MySqlTimeStamp string = "timestamp" //1970-01-01 00:00:00  YYYYMMDD HHMMSS
)
