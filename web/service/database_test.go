package service

import (
	"testing"

	"github.com/astaxie/beego/utils"
	"github.com/flourish-ship/work-attendance/web/model"
)

func Test_Insert_User(t *testing.T) {
	user := &model.User{UserName: "mojozjd"}
	CreateUser(user)
	utils.Display("error info is ", user)
}
