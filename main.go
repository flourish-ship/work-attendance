package main

import (
	"github.com/kataras/iris"
)

func main() {

	iris.Get("/hi", func(ctx *iris.Context) {
		ctx.Write("Hi %s", " test")
	})
	iris.Listen(":5000")
}
