package api

import (
	"net/http"

	"github.com/kataras/iris"
)

//args[0] http code,  args[1] error message
func ErrorHandler(u *iris.Context, err error, args ...interface{}) bool {

	var code int
	var message string

	if err == nil {
		return false
	}
	switch len(args) {
	case 1:
		code = args[0].(int)
	case 2:
		code = args[0].(int)
		message = args[1].(string)
	default:
		code = http.StatusBadRequest
		message = err.Error()
	}

	u.Error(message, code)
	return true
}
