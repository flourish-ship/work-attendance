package main

import (
	"runtime"

	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/flourish-ship/mongo-cli/services/mongo"
	"github.com/flourish-ship/skeleton/debuger"
	"github.com/flourish-ship/work-attendance/consts"
	"github.com/flourish-ship/work-attendance/routers"
)

func ini() {
	LoadConfig()
}

func main() {
	server := routers.Routers()

	runtime.GOMAXPROCS(runtime.NumCPU())
	err := mongo.Startup(
		consts.SessionName,
		consts.HostName,
		consts.DbName,
		consts.UserName,
		consts.PassWord,
	)

	if err != nil {
		fmt.Print("连接 MongoDB 失败...")
		return
	} else {
		fmt.Print("连接 MongoDB 成功...")
	}
	defer mongo.Shutdown()
	server.Listen(":3000")
}

func LoadConfig() {
	conf, err := config.NewConfig("ini", "app.conf")

	if err != nil {
		debuger.Display("配置文件获取失败:", err)
		return
	}

	runModel := conf.String("RunMode")
	consts.SessionName = conf.String(runModel + "::SessionName")
	consts.HostName = conf.String(runModel + "::HostName")
	consts.DbName = conf.String(runModel + "::DbName")
}
