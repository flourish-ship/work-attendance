package routers

import "github.com/kataras/iris"

func Routers() (frame *iris.Framework) {
	frame = iris.New()

	StaticHandler(frame)
	RegisterApi(frame)

	return
}

func StaticHandler(frame *iris.Framework) {
	frame.StaticWeb("/static", "./static", 1)
}

func RegisterApi(frame *iris.Framework) {
}
