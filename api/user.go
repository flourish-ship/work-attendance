package api

import (
	"encoding/json"

	"github.com/flourish-ship/skeleton/errors"
	"github.com/flourish-ship/work-attendance/consts"
	"github.com/flourish-ship/work-attendance/model"

	"github.com/flourish-ship/skeleton/response"
	//"github.com/flourish-ship/work-attendance/auth"
	"github.com/kataras/iris"
	"gopkg.in/mgo.v2/bson"
)

type UserApi struct {
}

func (userApi *UserApi) Register(api *iris.Framework) {

	user := api.Party(consts.URL_PRFIX + "/user")
	{
		user.Get("", GetUser)
		user.Post("/singin", SingIn)
		user.Post("", Create)
	}
}

func GetUser(u *iris.Context) {
	user := model.User{UserName: "mojo-zd", Gender: "male", NickName: "mojo"}
	model.GetUser(bson.ObjectIdHex("5747c0d3e10624155678e880"))
	u.JSON(iris.StatusOK, user)
}

// 登陆
func SingIn(u *iris.Context) {

	user := model.User{}

	if err := json.Unmarshal(u.Request.Body(), &user); err != nil {
		u.JSON(errors.BADREQUEST_CODE, response.DefaultResponse(consts.BAD_PARAMS_ERROR))
		return
	}

	//auth.Authentication(user.UserName, user.PassWord)
	//authentication := auth.Authentication(user.UserName, user.PassWord)
	//u.Session().Set("authUser", authentication)
}

func Create(u *iris.Context) {
	user := model.User{}
	if err := u.ReadForm(&user); err != nil {
		u.JSON(errors.BADREQUEST_CODE, err)
		return
	}
	u.JSON(iris.StatusOK, user)
}
