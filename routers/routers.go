package routers

import (
	"github.com/flourish-ship/work-attendance/consts"
	"github.com/flourish-ship/work-attendance/web/api"
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

func Routers() (frame *iris.Framework) {
	frame = iris.New()
	frame.Adapt(httprouter.New())

	//StaticHandler(frame)
	RegisterApi(frame)

	return
}

func StaticHandler(frame *iris.Framework) {
	frame.StaticWeb("/static", "./static")
}

func RegisterApi(webApi *iris.Framework) {
	party := webApi.Party(consts.URL_PRFIX)
	{
		//user part
		party.Get("/user", api.GetUser)
		party.Post("/user", api.Create)
	}
}
