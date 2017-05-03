package consts

const (
	URL_PRFIX = "/api/v1"
)

// error warning info
var (
	BAD_PARAMS_ERROR = "参数传递错误!"
)

// db info
var (
	DATA_SOURCE               = "%s:%s@tcp(%s:%s)/%s?charset=utf8"
	DATABASE_DEFAULT_USER     = "root"
	DATABASE_DEFAULT_PASSWORD = "root123"
	DATABASE_DEFAULT_PORT     = "3306"
	DATABASE_DEFAULT_HOST     = "127.0.0.1"
	DATABASE_DEFAULT          = "attendence"
)
