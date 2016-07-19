package routers

import (
	"github.com/flourish-ship/work-attendance/api"

	"github.com/kataras/iris"
)

func init() {
	iris.API("/users", api.UserApi{})
}
