package service

import (
	"fmt"

	"github.com/flourish-ship/work-attendance/database/orm"
	"github.com/flourish-ship/work-attendance/web/model"
)

func CreateUser(user *model.User) (err error) {
	id, err := orm.NewOrmEngin().Insert(user)
	fmt.Println(id)
	return
}
