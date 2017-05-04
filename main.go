package main

import (
	"fmt"

	"github.com/flourish-ship/work-attendance/database/migrate"
	_ "github.com/flourish-ship/work-attendance/database/orm"

	"github.com/flourish-ship/work-attendance/routers"
)

func main() {
	migrate.AutoMigrate()
	server := routers.Routers()

	fmt.Println("listen on 3000")
	server.Listen(":3000")
}
