package routers

import (
	"github.com/flourish-ship/work-attendance/api"
	"github.com/flourish-ship/work-attendance/render"
	"github.com/flourish-ship/work-attendance/socket"
	"github.com/kataras/iris"
)

func Routers() (frame *iris.Framework) {

	frame = iris.New()
	userApi := api.UserApi{}
	chatApi := socket.ChartRoomApi{}
	homeRender := render.HomeRender{}

	userApi.Register(frame)
	homeRender.Register(frame)
	chatApi.Chart(frame)

	return
}
