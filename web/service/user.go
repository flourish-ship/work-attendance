package service

import (
	"github.com/flourish-ship/work-attendance/database/orm"
	"github.com/flourish-ship/work-attendance/web/model"
)

func CreateUser(user *model.User) (err error) {
	engin := orm.NewOrmEngin()
	_, err = engin.Insert(user)
	return
}
