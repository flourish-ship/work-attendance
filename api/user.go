package api

import (
	"github.com/flourish-ship/work-attendance/model"
	"github.com/kataras/iris"
)

type UserApi struct {
	*iris.Context
}

// GET /users
func (u UserApi) Get() {
	//u.Write("Get from /users")
	user := model.User{UserName: "mojo-zd", Gender: "male"}
	u.JSON(iris.StatusOK, user)
}
