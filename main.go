package main

import "github.com/flourish-ship/work-attendance/routers"

func main() {
	server := routers.Routers()
	server.Listen(":3000")
}
