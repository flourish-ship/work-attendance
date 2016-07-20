package main

import (
	"github.com/flourish-ship/work-attendance/routers"
)

func main() {
	routers.Routers().Listen(":3000")
}
