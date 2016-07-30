package render

import (
	"github.com/flourish-ship/work-attendance/consts"
	"github.com/kataras/iris"
)

type HomeRender struct {
}

type HomeData struct {
	Title   string
	Message string
}

func Config(frame *iris.Framework) {
	frame.Config.IsDevelopment = true
	frame.Config.Gzip = true
	frame.Config.Charset = "utf-8"
}

func (render *HomeRender) Register(api *iris.Framework) {
	Config(api)
	homeRender := api.Party(consts.URL_PRFIX + "/home")
	{
		homeRender.Get("", HomeIndex)
	}
}

func HomeIndex(c *iris.Context) {
	homeData := HomeData{Title: "Iris", Message: "hello world"}
	c.Render("home.html", homeData)
}
