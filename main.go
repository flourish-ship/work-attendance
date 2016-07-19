package main

import (
	_ "github.com/flourish-ship/work-attendance/routers"
	"github.com/kataras/iris"
)

func main() {

	iris.Listen(":3000")
}
