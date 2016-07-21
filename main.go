package main

import (
	"runtime"

	"fmt"
	"github.com/flourish-ship/mongo-cli/services/mongo"
	"github.com/flourish-ship/work-attendance/consts"
	"github.com/flourish-ship/work-attendance/routers"
)

func main() {
	server := routers.Routers()

	runtime.GOMAXPROCS(runtime.NumCPU())
	err := mongo.Startup(consts.SESSIONNAME, consts.HOSTS, consts.DBNAME, "", "")

	if err != nil {
		fmt.Print("连接 MongoDB 失败...")
		return
	} else {
		fmt.Print("连接 MongoDB 成功...")
	}
	defer mongo.Shutdown()
	server.Listen(":3000")
}
