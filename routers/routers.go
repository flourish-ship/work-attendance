package routers

import (
	"github.com/flourish-ship/work-attendance/api"
	"github.com/flourish-ship/work-attendance/socket"
	"github.com/kataras/iris"
)

func Routers() (frame *iris.Framework) {

	frame = iris.New()
	userApi := api.UserApi{}
	chatApi := socket.ChartRoomApi{}

	userApi.Register(frame)
	chatApi.Chart(frame)

	return
}
