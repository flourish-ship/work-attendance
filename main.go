package main

import "github.com/kataras/iris"

func main() {
	iris.Get("/test", func(ctx *iris.Context) {
		ctx.Write("It's OK!")
	})

	iris.Listen(":6000")
}
