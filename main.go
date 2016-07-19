package main

import (
	"github.com/flourish-ship/work-attendance/api"
	"github.com/kataras/iris"
)

func main() {
	iris.API("/users", api.UserApi{})

	iris.Listen(":3000")
}
