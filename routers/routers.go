package routers

import (
	"github.com/flourish-ship/work-attendance/api"
	"github.com/kataras/iris"
)

func Routers() (frame *iris.Framework) {

	frame = iris.New()
	userApi := api.UserApi{}
	userApi.Register(frame)

	return
}
