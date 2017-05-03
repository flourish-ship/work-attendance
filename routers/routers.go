package routers

import (
	"github.com/flourish-ship/work-attendance/consts"
	"github.com/flourish-ship/work-attendance/web/api"
	"github.com/kataras/iris"
)

func Routers() (frame *iris.Framework) {
	frame = iris.New()

	StaticHandler(frame)
	RegisterApi(frame)

	return
}

func StaticHandler(frame *iris.Framework) {
	frame.StaticWeb("/static", "./static", 1)
}

func RegisterApi(webApi *iris.Framework) {
	party := webApi.Party(consts.URL_PRFIX)
	{
		//user part
		party.Get("/user", api.GetUser)
		party.Post("/user", api.Create)
	}
}
