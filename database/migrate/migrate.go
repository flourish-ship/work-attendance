package migrate

import (
	"github.com/flourish-ship/work-attendance/database/orm"
	"github.com/flourish-ship/work-attendance/web/model"
)

func AutoMigrate() {
	engin := orm.NewOrmEngin()
	engin.CreateTables(model.User{})
	engin.Sync2(model.User{})
}
