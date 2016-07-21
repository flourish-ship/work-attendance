package api

import (
	"github.com/flourish-ship/skeleton/errors"
	"github.com/flourish-ship/work-attendance/consts"
	"github.com/flourish-ship/work-attendance/model"

	"github.com/kataras/iris"
	"gopkg.in/mgo.v2/bson"
)

type UserApi struct {
}

func (userApi *UserApi) Register(api *iris.Framework) {
	user := api.Party(consts.URL_PRFIX + "/user")
	{
		user.Get("", GetUser)
		user.Post("", Create)
	}
}

func GetUser(u *iris.Context) {
	user := model.User{UserName: "mojo-zd", Gender: "male", NickName: "mojo"}
	model.GetUser(bson.ObjectIdHex("5747c0d3e10624155678e880"))
	u.JSON(iris.StatusOK, user)
}

func Create(u *iris.Context) {
	user := model.User{}
	if err := u.ReadForm(&user); err != nil {
		u.JSON(errors.BADREQUEST_CODE, err)
		return
	}
	u.JSON(iris.StatusOK, user)
}
