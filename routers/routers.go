package routers

import (
	"github.com/flourish-ship/work-attendance/api"
	"github.com/flourish-ship/work-attendance/render"
	"github.com/flourish-ship/work-attendance/socket"
	"github.com/kataras/iris"
)

func Routers() (frame *iris.Framework) {
	frame = iris.New()

	StaticHandler(frame)
	RegisterApi(frame)

	return
}

func StaticHandler(frame *iris.Framework) {
	//frame.StaticHandler("static", 0, false, false, []string{})
	//frame.StaticServe("static")
	frame.StaticWeb("/static", "./static", 1)
}

func RegisterApi(frame *iris.Framework) {
	userApi := api.UserApi{}
	chatApi := socket.ChartRoomApi{}
	homeRender := render.HomeRender{}

	userApi.Register(frame)
	homeRender.Register(frame)
	chatApi.Chart(frame)
}
