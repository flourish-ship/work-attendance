package main

import (
	"github.com/flourish-ship/work-attendance/database/migrate"
	_ "github.com/flourish-ship/work-attendance/database/orm"
	"github.com/flourish-ship/work-attendance/routers"
)

func main() {
	migrate.AutoMigrate()
	server := routers.Routers()
	server.Listen(":3000")
}
