package main

<<<<<<< HEAD
import (
	"github.com/kataras/iris"
)

func main() {

	iris.Get("/hi", func(ctx *iris.Context) {
		ctx.Write("Hi %s", "iris")
	})
	iris.Listen(":8080")
=======
import "github.com/kataras/iris"

func main() {
	iris.Get("/test", func(ctx *iris.Context) {
		ctx.Write("It's OK!")
	})

	iris.Listen(":6000")
>>>>>>> a5db72ba10cc61fc48411b451e8a281991891212
}
