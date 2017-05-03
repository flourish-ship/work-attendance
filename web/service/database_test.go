package service

import (
	"testing"

	"github.com/astaxie/beego/utils"
	"github.com/flourish-ship/work-attendance/web/model"
)

func Test_Insert_User(t *testing.T) {
	err := CreateUser(&model.User{UserName: "mojo"})
	utils.Display("error info is ", err)
}
